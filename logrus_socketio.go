package logrus_socketio

import (
	"github.com/Sirupsen/logrus"
	"github.com/zhouhui8915/go-socket.io-client"
)

type SocketIOHook struct {
	Client	    	*socketio_client.Client
	EventName     	string
	LogExtraFields	map[string]interface{}
}

func NewSocketIOHook(uri string, event string, extraLogFields map[string]interface{}) (*SocketIOHook, error) {
	opts := &socketio_client.Options{
		Transport: "websocket",
		Query:     make(map[string]string),
	}

	client, err := socketio_client.NewClient(uri, opts)
	if err != nil {
		return &SocketIOHook{}, err
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
