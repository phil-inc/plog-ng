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
	level *int
}

var e *entry

func Init() {
	log.SetFormatter(&log.JSONFormatter{})
	e = new(entry)
	// log.SetLevel(log.DebugLevel) //Note this must be enabled from the application
}

func (entry *entry) WithCallerLevel(l int) *entry {
	entry.level = &l
	return entry
}

func (entry *entry) Info(s string) {
	pc, file, line, ok := runtime.Caller(getCallerLevel(entry))
	if ok {
		entry.WithFields(getAdditionalFields(pc, file, line)).Info(s)
		return
	}
	entry.Info(s)
}

func (entry *entry) Infof(s string, args ...interface{}) {
	pc, file, line, ok := runtime.Caller(getCallerLevel(entry))
	if ok {
		entry.WithFields(getAdditionalFields(pc, file, line)).Infof(s, args...)
		return
	}
	entry.Infof(s, args...)
}

func (entry *entry) Debug(s string) {
	pc, file, line, ok := runtime.Caller(getCallerLevel(entry))
	if ok {
		entry.WithFields(getAdditionalFields(pc, file, line)).Debug(s)
		return
	}
	entry.Debug(s)
}

func (entry *entry) Debugf(s string, args ...interface{}) {
	pc, file, line, ok := runtime.Caller(getCallerLevel(entry))
	if ok {
		entry.WithFields(getAdditionalFields(pc, file, line)).Debugf(s, args...)
		return
	}
	entry.Debugf(s, args...)
}

func (entry *entry) Warn(s string) {
	pc, file, line, ok := runtime.Caller(getCallerLevel(entry))
	if ok {
		entry.WithFields(getAdditionalFields(pc, file, line)).Warn(s)
		return
	}
	entry.Warn(s)
}

func (entry *entry) Warnf(s string, args ...interface{}) {
	pc, file, line, ok := runtime.Caller(getCallerLevel(entry))
	if ok {
		entry.WithFields(getAdditionalFields(pc, file, line)).Warnf(s, args...)
		return
	}
	entry.Warnf(s, args...)
}

func (entry *entry) Error(s string) {
	pc, file, line, ok := runtime.Caller(getCallerLevel(entry))
	if ok {
		entry.WithFields(getAdditionalFields(pc, file, line)).Error(s)
		return
	}
	entry.Error(s)
}

func (entry *entry) Errorf(s string, args ...interface{}) {
	pc, file, line, ok := runtime.Caller(getCallerLevel(entry))
	if ok {
		entry.WithFields(getAdditionalFields(pc, file, line)).Errorf(s, args...)
		return
	}
	entry.Errorf(s, args...)
}

func (entry *entry) Panic(s string) {
	pc, file, line, ok := runtime.Caller(getCallerLevel(entry))
	if ok {
		entry.WithFields(getAdditionalFields(pc, file, line)).Panic(s)
		return
	}
	entry.Panic(s)
}

func (entry *entry) Panicf(s string, args ...interface{}) {
	pc, file, line, ok := runtime.Caller(getCallerLevel(entry))
	if ok {
		entry.WithFields(getAdditionalFields(pc, file, line)).Panicf(s, args...)
		return
	}
	entry.Panicf(s, args...)
}

func getAdditionalFields(pc uintptr, file string, line int) log.Fields {
	details := runtime.FuncForPC(pc)
	if details != nil {
		caller := strings.Split(details.Name(), "/")
		file := strings.Split(file, "/")
		fields := log.Fields{
			"caller": caller[len(caller)-1],
			"file":   fmt.Sprintf("%s:%d", file[len(file)-1], line),
		}
		return fields
	}

	return log.Fields{}
}

func getCallerLevel(e *entry) int {
	if e.level == nil {
		return 1
	}

	return *e.level
}
