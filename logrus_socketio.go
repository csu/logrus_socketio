package logrus_http

import (
	"github.com/Sirupsen/logrus"
	"github.com/zhouhui8915/go-socket.io-client"
)

type SocketIOHook struct {
	Client	    	*socketio_client.Client
	EventName     	string
	LogExtraFields  map[string]interface{}
}

// Creates a hook to be added to an instance of logger. This is called with
// `hook, err := NewSocketIOHook("http://log-server/post_new_log", "logBody")`
// `if err == nil { log.Hooks.Add(hook) }`
func NewSocketIOHook(uri string, event string, extraLogFields map[string]interface{}) (*SocketIOHook, error) {
	opts := &socketio_client.Options{
		Transport: "websocket",
	}

	client, err := socketio_client.NewClient(uri, opts)
	if err != nil {
		return nil, err
	}

	return &SocketIOHook{client, event, extraLogFields}, nil
}

func (hook *SocketIOHook) Fire(entry *logrus.Entry) error {
	line, err := entry.WithFields(hook.LogExtraFields).String()
	if err != nil {
		return err
	}

	hook.Client.Emit(hook.EventName, line)

	return nil
}

func (hook *SocketIOHook) Levels() []logrus.Level {
	return logrus.AllLevels
}
