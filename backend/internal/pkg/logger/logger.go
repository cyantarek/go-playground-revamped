package logger

import (
	"github.com/keepeye/logrus-filename"
	"github.com/sebest/logrusly"
	"github.com/sirupsen/logrus"
)

var Log *logrus.Logger

var logglyToken string = "fb9d10b8-26fc-4e46-ab35-0db8d28024ac"

func New() {
	Log = logrus.New()
	filenameHook := filename.NewHook()
	filenameHook.Field = "line"
	Log.SetFormatter(&logrus.JSONFormatter{})
	hook := logrusly.NewLogglyHook(logglyToken, "https://logs-01.loggly.com/bulk/", logrus.WarnLevel, "go", "logrus")

	Log.Hooks.Add(hook)
	Log.Hooks.Add(filenameHook)

	// hook.Flush()
}
