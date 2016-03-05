package logrus_http

import (
	"net/http"
	"net/url"

	"github.com/Sirupsen/logrus"
)

type HttpHook struct {
	RequestEndpoint    string
	RequestFormKey     string
	RequestExtraFields map[string]string
	LogExtraFields     map[string]interface{}
}

// Creates a hook to be added to an instance of logger. This is called with
// `hook, err := NewHttpHook("http://log-server/post_new_log", "logBody")`
// `if err == nil { log.Hooks.Add(hook) }`
func NewHttpHook(endpoint string, formKey string, requestExtraFields map[string]string,
	extraLogFields map[string]interface{}) (*HttpHook, error) {
	return &HttpHook{endpoint, formKey, requestExtraFields, extraLogFields}, nil
}

func (hook *HttpHook) Fire(entry *logrus.Entry) error {
	line, err := entry.WithFields(hook.LogExtraFields).String()
	if err != nil {
		return err
	}

	reqForm := url.Values{}

	// add in extra fields, if any
	for k, v := range hook.RequestExtraFields {
		reqForm.Set(k, v)
	}

	// add log line
	reqForm.Set(hook.RequestFormKey, line)

	resp, err := http.PostForm(hook.RequestEndpoint, reqForm)
	if err != nil {
		return err
	}
	resp.Body.Close()
	return nil
}

func (hook *HttpHook) Levels() []logrus.Level {
	return logrus.AllLevels
}
