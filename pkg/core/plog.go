package core

import (
	"fmt"
	"runtime"
	"strings"

	log "github.com/sirupsen/logrus"
)

type Fields map[string]interface{}
type entry struct {
	*log.Entry
}

var Entry *entry

func Init() {
	log.SetFormatter(&log.JSONFormatter{})
	Entry = new(entry)
	Entry.Entry = log.NewEntry(log.StandardLogger())
	// log.SetLevel(log.DebugLevel) //Note this must be enabled from the application
}

func WithFields(fields Fields) *entry {
	return &entry{log.WithFields((log.Fields)(fields))}
}

func (entry *entry) Info(s string) {
	pc, file, line, ok := runtime.Caller(1)
	if ok {
		entry.WithFields(getAdditionalFields(pc, file, line)).Info(s)
		return
	}
	entry.Info(s)
}

func (entry *entry) Infof(s string, args ...interface{}) {
	pc, file, line, ok := runtime.Caller(1)
	if ok {
		entry.WithFields(getAdditionalFields(pc, file, line)).Infof(s, args...)
		return
	}
	entry.Infof(s, args...)
}

func (entry *entry) Debug(s string) {
	pc, file, line, ok := runtime.Caller(1)
	if ok {
		entry.WithFields(getAdditionalFields(pc, file, line)).Debug(s)
		return
	}
	entry.Debug(s)
}

func (entry *entry) Debugf(s string, args ...interface{}) {
	pc, file, line, ok := runtime.Caller(1)
	if ok {
		entry.WithFields(getAdditionalFields(pc, file, line)).Debugf(s, args...)
		return
	}
	entry.Debugf(s, args...)
}

func (entry *entry) Warn(s string) {
	pc, file, line, ok := runtime.Caller(1)
	if ok {
		entry.WithFields(getAdditionalFields(pc, file, line)).Warn(s)
		return
	}
	entry.Warn(s)
}

func (entry *entry) Warnf(s string, args ...interface{}) {
	pc, file, line, ok := runtime.Caller(1)
	if ok {
		entry.WithFields(getAdditionalFields(pc, file, line)).Warnf(s, args...)
		return
	}
	entry.Warnf(s, args...)
}

func (entry *entry) Error(s string) {
	pc, file, line, ok := runtime.Caller(1)
	if ok {
		entry.WithFields(getAdditionalFields(pc, file, line)).Error(s)
		return
	}
	entry.Error(s)
}

func (entry *entry) Errorf(s string, args ...interface{}) {
	pc, file, line, ok := runtime.Caller(1)
	if ok {
		entry.WithFields(getAdditionalFields(pc, file, line)).Errorf(s, args...)
		return
	}
	entry.Warnf(s, args...)
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
