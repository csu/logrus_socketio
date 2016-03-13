package logrus_socketio

import (
	"github.com/Sirupsen/logrus"
	"testing"
	"github.com/googollee/go-socket.io"
	"net/http"
)

func TestPrint(t *testing.T) {
	// start up server
	server, err := socketio.NewServer(nil)
	if err != nil {
		t.Error(err)
	}
	server.On("connection", func(so socketio.Socket) {})
	server.On("error", func(so socketio.Socket, err error) {
		t.Error(err)
	})
	http.Handle("/socket.io/", server)
	go http.ListenAndServe(":3000", nil)

	log := logrus.New()
	log.Formatter = new(logrus.JSONFormatter)

	m := make(map[string]interface{})

	hook, err := NewSocketIOHook("http://localhost:3000", "log", m)
	if err != nil {
		t.Error(err)
		t.Errorf("Unable to create hook.")
	}

	log.Hooks.Add(hook)

	log.Info("It worked!")
}