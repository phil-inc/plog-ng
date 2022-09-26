package core

import (
	"runtime"

	log "github.com/sirupsen/logrus"
)

func WithFields(fields Fields) *entry {
	e.Entry = e.WithFields((log.Fields)(fields))
	return e
}

func Info(s string) {
	pc, file, line, ok := runtime.Caller(1)
	if ok {
		e.WithFields(getAdditionalFields(pc, file, line)).Info(s)
		return
	}
	e.Info(s)
}

func Infof(s string, args ...interface{}) {
	pc, file, line, ok := runtime.Caller(1)
	if ok {
		e.WithFields(getAdditionalFields(pc, file, line)).Infof(s, args...)
		return
	}
	e.Infof(s, args...)
}

func Debug(s string) {
	pc, file, line, ok := runtime.Caller(1)
	if ok {
		e.WithFields(getAdditionalFields(pc, file, line)).Debug(s)
		return
	}
	e.Debug(s)

}

func Debugf(s string, args ...interface{}) {
	pc, file, line, ok := runtime.Caller(1)
	if ok {
		e.WithFields(getAdditionalFields(pc, file, line)).Debugf(s, args...)
		return
	}
	e.Debugf(s, args...)
}

func Warn(s string) {
	pc, file, line, ok := runtime.Caller(1)
	if ok {
		e.WithFields(getAdditionalFields(pc, file, line)).Warn(s)
		return
	}
	e.Warn(s)
}

func Warnf(s string, args ...interface{}) {
	pc, file, line, ok := runtime.Caller(1)
	if ok {
		e.WithFields(getAdditionalFields(pc, file, line)).Warnf(s, args...)
		return
	}
	e.Warnf(s, args...)
}

func Error(s string) {
	pc, file, line, ok := runtime.Caller(1)
	if ok {
		e.WithFields(getAdditionalFields(pc, file, line)).Error(s)
		return
	}
	e.Error(s)
}

func Errorf(s string, args ...interface{}) {
	pc, file, line, ok := runtime.Caller(1)
	if ok {
		e.WithFields(getAdditionalFields(pc, file, line)).Errorf(s, args...)
		return
	}
	e.Errorf(s, args...)
}
