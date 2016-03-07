package logrus_socketio

import (
  "github.com/Sirupsen/logrus"
  "testing"
)

func TestPrint(t *testing.T) {
  log := logrus.New()
  log.Formatter = new(logrus.JSONFormatter)

  m := make(map[string]interface{})
  
  hook, err := NewSocketIOHook("http://logserver.christopher.su/", "log", m)
  if err != nil {
    t.Errorf("Unable to create hook.")
  }

  log.Hooks.Add(hook)

  log.Info("It worked!")
}