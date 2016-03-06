package logrus_http

import (
	"github.com/Sirupsen/logrus"
	"golang.org/x/net/websocket"
)

type WebsocketHook struct {
	Websocket    	*websocket.Conn
	EventName     	string
	LogExtraFields  map[string]interface{}
}

// Creates a hook to be added to an instance of logger. This is called with
// `hook, err := NewWebsocketHook("http://log-server/post_new_log", "logBody")`
// `if err == nil { log.Hooks.Add(hook) }`
func NewWebsocketHook(endpoint string, origin string, event string,
	extraLogFields map[string]interface{}) (*WebsocketHook, error) {
	ws, err := websocket.Dial(endpoint, "", origin)
	if err != nil {
		return nil, err
	}

	return &WebsocketHook{ws, event, extraLogFields}, nil
}

func (hook *WebsocketHook) Fire(entry *logrus.Entry) error {
	line, err := entry.WithFields(hook.LogExtraFields).String()
	if err != nil {
		return err
	}

	_, err = hook.Websocket.Write([]byte(line))
	if err != nil {
		return err
	}

	return nil
}

func (hook *WebsocketHook) Levels() []logrus.Level {
	return logrus.AllLevels
}
