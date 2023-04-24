package core

import (
	"runtime"

	log "github.com/sirupsen/logrus"
)

const (
	// PanicLevel level, highest level of severity. Logs and then calls panic with the
	// message passed to Debug, Info, ...
	PanicLevel log.Level = iota
	// FatalLevel level. Logs and then calls `logger.Exit(1)`. It will exit even if the
	// logging level is set to Panic.
	FatalLevel
	// ErrorLevel level. Logs. Used for errors that should definitely be noted.
	// Commonly used for hooks to send errors to an error tracking service.
	ErrorLevel
	// WarnLevel level. Non-critical entries that deserve eyes.
	WarnLevel
	// InfoLevel level. General operational entries about what's going on inside the
	// application.
	InfoLevel
	// DebugLevel level. Usually only enabled when debugging. Very verbose logging.
	DebugLevel
	// TraceLevel level. Designates finer-grained informational events than the Debug.
	TraceLevel
)

func SetLevel(level log.Level) {
	log.SetLevel(level)
}

func GetLevel() log.Level {
	return log.GetLevel()
}

func WithFields(fields Fields) *entry {
	e.Entry = log.NewEntry(log.StandardLogger())
	e.Entry = e.WithFields((log.Fields)(fields))
	return e
}

func Info(s string) {
	e.Entry = log.NewEntry(log.StandardLogger())
	pc, file, line, ok := runtime.Caller(1)
	if ok {
		e.WithFields(getAdditionalFields(pc, file, line)).Info(s)
		return
	}
	e.Info(s)
}

func Infof(s string, args ...interface{}) {
	e.Entry = log.NewEntry(log.StandardLogger())
	pc, file, line, ok := runtime.Caller(1)
	if ok {
		e.WithFields(getAdditionalFields(pc, file, line)).Infof(s, args...)
		return
	}
	e.Infof(s, args...)
}

func Debug(s string) {
	e.Entry = log.NewEntry(log.StandardLogger())
	pc, file, line, ok := runtime.Caller(1)
	if ok {
		e.WithFields(getAdditionalFields(pc, file, line)).Debug(s)
		return
	}
	e.Debug(s)

}

func Debugf(s string, args ...interface{}) {
	e.Entry = log.NewEntry(log.StandardLogger())
	pc, file, line, ok := runtime.Caller(1)
	if ok {
		e.WithFields(getAdditionalFields(pc, file, line)).Debugf(s, args...)
		return
	}
	e.Debugf(s, args...)
}

func Warn(s string) {
	e.Entry = log.NewEntry(log.StandardLogger())
	pc, file, line, ok := runtime.Caller(1)
	if ok {
		e.WithFields(getAdditionalFields(pc, file, line)).Warn(s)
		return
	}
	e.Warn(s)
}

func Warnf(s string, args ...interface{}) {
	e.Entry = log.NewEntry(log.StandardLogger())
	pc, file, line, ok := runtime.Caller(1)
	if ok {
		e.WithFields(getAdditionalFields(pc, file, line)).Warnf(s, args...)
		return
	}
	e.Warnf(s, args...)
}

func Error(s string) {
	e.Entry = log.NewEntry(log.StandardLogger())
	pc, file, line, ok := runtime.Caller(1)
	if ok {
		e.WithFields(getAdditionalFields(pc, file, line)).Error(s)
		return
	}
	e.Error(s)
}

func Errorf(s string, args ...interface{}) {
	e.Entry = log.NewEntry(log.StandardLogger())
	pc, file, line, ok := runtime.Caller(1)
	if ok {
		e.WithFields(getAdditionalFields(pc, file, line)).Errorf(s, args...)
		return
	}
	e.Errorf(s, args...)
}

func Panic(s string) {
	e.Entry = log.NewEntry(log.StandardLogger())
	pc, file, line, ok := runtime.Caller(1)
	if ok {
		e.WithFields(getAdditionalFields(pc, file, line)).Panic(s)
		return
	}
	e.Panic(s)
}

func Panicf(s string, args ...interface{}) {
	e.Entry = log.NewEntry(log.StandardLogger())
	pc, file, line, ok := runtime.Caller(1)
	if ok {
		e.WithFields(getAdditionalFields(pc, file, line)).Panicf(s, args...)
		return
	}
	e.Panicf(s, args...)
}
