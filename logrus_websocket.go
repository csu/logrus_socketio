package logrus_http

import (
	"github.com/Sirupsen/logrus"
	"github.com/googollee/go-socket.io"
)

type SocketIOHook struct {
	Socket	    	socketio.Socket
	EventName     	string
	LogExtraFields  map[string]interface{}
}

// Creates a hook to be added to an instance of logger. This is called with
// `hook, err := NewSocketIOHook("http://log-server/post_new_log", "logBody")`
// `if err == nil { log.Hooks.Add(hook) }`
func NewSocketIOHook(endpoint string, event string, extraLogFields map[string]interface{}) (*SocketIOHook, error) {
	server, err := socketio.NewServer(endpoint)
	if err != nil {
		return nil, err
	}

	server.On("connection", func(so socketio.Socket) {
		return &SocketIOHook{so, event, extraLogFields}, nil
	})

	return nil, nil
}

func (hook *SocketIOHook) Fire(entry *logrus.Entry) error {
	line, err := entry.WithFields(hook.LogExtraFields).String()
	if err != nil {
		return err
	}

	_, err = hook.Socket.Emit(hook.EventName, line)
	if err != nil {
		return err
	}

	return nil
}

func (hook *SocketIOHook) Levels() []logrus.Level {
	return logrus.AllLevels
}
