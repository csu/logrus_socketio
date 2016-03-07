package logrus_socketio

import (
	"github.com/Sirupsen/logrus"
	"github.com/zhouhui8915/go-socket.io-client"
)

type SocketIOHook struct {
	Client	    	*socketio_client.Client
	EventName     	string
}

func NewSocketIOHook(uri string, event string) (*SocketIOHook, error) {
	opts := &socketio_client.Options{
		Transport: "websocket",
	}

	client, err := socketio_client.NewClient(uri, opts)
	if err != nil {
		return nil, err
	}

	return &SocketIOHook{client, event}, nil
}

func (hook *SocketIOHook) Fire(entry *logrus.Entry) error {
	//line, err := entry.WithFields(hook.LogExtraFields).String()
	line, err := entry.String()
	if err != nil {
		return err
	}

	hook.Client.Emit(hook.EventName, line)

	return nil
}

func (hook *SocketIOHook) Levels() []logrus.Level {
	return logrus.AllLevels
}
