package core

import (
	"fmt"
	"runtime"
	"strings"

	log "github.com/sirupsen/logrus"
)

type Fields map[string]interface{}

func Init() {
	log.SetFormatter(&log.JSONFormatter{})
	// log.SetLevel(log.DebugLevel) //Note this must be enabled from the application
}

func Info(msg string) {
	pc, file, line, ok := runtime.Caller(1)
	if ok {
		log.WithFields(getAdditionalFields(pc, file, line)).Info(msg)
		return
	}

	log.Info(msg)
}

func Infof(s string, args ...interface{}) {
	pc, file, line, ok := runtime.Caller(1)
	if ok {
		log.WithFields(getAdditionalFields(pc, file, line)).Infof(s, args...)
		return
	}

	log.Infof(s, args...)
}

func InfoWithAdditionalFields(s string, args interface{}) {
	pc, file, line, ok := runtime.Caller(1)
	if ok {
		l := log.WithFields(getAdditionalFields(pc, file, line))
		l = mergeAddionalFields(l, args)
		l.Info(s)
		return
	}

	af, ok := args.(map[string]interface{})
	if ok {
		l := log.WithFields(af)
		l.Info(s)
		return
	}

	log.Info(s)
}

func Debug(msg string) {
	pc, file, line, ok := runtime.Caller(1)
	if ok {
		log.WithFields(getAdditionalFields(pc, file, line)).Debug(msg)
		return
	}

	log.Debug(msg)
}

func Debugf(s string, args ...interface{}) {
	pc, file, line, ok := runtime.Caller(1)
	if ok {
		log.WithFields(getAdditionalFields(pc, file, line)).Debugf(s, args...)
		return
	}

	log.Debugf(s, args...)
}

func DebugWithAdditionalFields(s string, args interface{}) {
	pc, file, line, ok := runtime.Caller(1)
	if ok {
		l := log.WithFields(getAdditionalFields(pc, file, line))
		l = mergeAddionalFields(l, args)
		l.Debug(s)
		return
	}

	af, ok := args.(map[string]interface{})
	if ok {
		l := log.WithFields(af)
		l.Debug(s)
		return
	}

	log.Debug(s)
}

func Warn(msg string) {
	pc, file, line, ok := runtime.Caller(1)
	if ok {
		log.WithFields(getAdditionalFields(pc, file, line)).Warn(msg)
		return
	}

	log.Warn(msg)
}

func Warnf(s string, args ...interface{}) {
	pc, file, line, ok := runtime.Caller(1)
	if ok {
		log.WithFields(getAdditionalFields(pc, file, line)).Warnf(s, args...)
		return
	}

	log.Warnf(s, args...)
}

func WarnWithAdditionalFields(s string, args interface{}) {
	pc, file, line, ok := runtime.Caller(1)
	if ok {
		l := log.WithFields(getAdditionalFields(pc, file, line))
		l = mergeAddionalFields(l, args)
		l.Warn(s)
		return
	}

	af, ok := args.(map[string]interface{})
	if ok {
		l := log.WithFields(af)
		l.Warn(s)
		return
	}

	log.Warn(s)
}

func Error(msg string) {
	pc, file, line, ok := runtime.Caller(1)
	if ok {
		log.WithFields(getAdditionalFields(pc, file, line)).Error(msg)
		return
	}

	log.Error(msg)
}

func Errorf(s string, args ...interface{}) {
	pc, file, line, ok := runtime.Caller(1)
	if ok {
		log.WithFields(getAdditionalFields(pc, file, line)).Errorf(s, args...)
		return
	}

	log.Errorf(s, args...)
}

func ErrorWithAdditionalFields(s string, args interface{}) {
	pc, file, line, ok := runtime.Caller(1)
	if ok {
		l := log.WithFields(getAdditionalFields(pc, file, line))
		l = mergeAddionalFields(l, args)
		l.Error(s)
		return
	}

	af, ok := args.(map[string]interface{})
	if ok {
		l := log.WithFields(af)
		l.Error(s)
		return
	}

	log.Error(s)
}

func getAdditionalFields(pc uintptr, file string, line int) log.Fields {
	details := runtime.FuncForPC(pc)
	if details != nil {
		fn := strings.Split(details.Name(), ".")
		fields := log.Fields{
			"package":  fn[0],
			"function": fn[1],
			"file":     fmt.Sprintf("%s:%d", file, line),
		}
		return fields
	}

	return log.Fields{}
}

func mergeAddionalFields(l *log.Entry, args interface{}) *log.Entry {
	af, ok := args.(map[string]interface{})
	if ok {
		l = l.WithFields(af)
		return l
	}

	return l
}
