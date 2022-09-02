package core

import (
	"fmt"
	"runtime"
	"strings"

	log "github.com/sirupsen/logrus"
)

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

func Debug(msg string) {
	pc, file, line, ok := runtime.Caller(1)
	if ok {
		log.WithFields(getAdditionalFields(pc, file, line)).Debug(msg)
		return
	}

	log.Debug(msg)
}

func Warn(msg string) {
	pc, file, line, ok := runtime.Caller(1)
	if ok {
		log.WithFields(getAdditionalFields(pc, file, line)).Warn(msg)
		return
	}

	log.Warn(msg)
}

func Error(msg string) {
	pc, file, line, ok := runtime.Caller(1)
	if ok {
		log.WithFields(getAdditionalFields(pc, file, line)).Error(msg)
		return
	}

	log.Error(msg)
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
