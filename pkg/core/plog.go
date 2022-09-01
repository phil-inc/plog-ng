package core

import (
	"runtime"
	"strings"

	log "github.com/sirupsen/logrus"
)

func Init() {
	log.SetFormatter(&log.JSONFormatter{})
	// log.SetLevel(log.DebugLevel) //Note this must be enabled from the application
}

func Info(msg string) {
	pc, file, _, ok := runtime.Caller(1)
	if ok {
		log.WithFields(getAdditionalFields(pc, file)).Info(msg)
		return
	}

	log.Info(msg)
}

func Debug(msg string) {
	pc, file, _, ok := runtime.Caller(1)
	if ok {
		log.WithFields(getAdditionalFields(pc, file)).Debug(msg)
		return
	}

	log.Debug(msg)
}

func Warn(msg string) {
	pc, file, _, ok := runtime.Caller(1)
	if ok {
		log.WithFields(getAdditionalFields(pc, file)).Warn(msg)
		return
	}

	log.Warn(msg)
}

func Error(msg string) {
	pc, file, _, ok := runtime.Caller(1)
	if ok {
		log.WithFields(getAdditionalFields(pc, file)).Error(msg)
		return
	}

	log.Error(msg)
}

func getAdditionalFields(pc uintptr, file string) log.Fields {
	details := runtime.FuncForPC(pc)
	if details != nil {
		fn := strings.Split(details.Name(), ".")
		fields := log.Fields{
			"package":  fn[0],
			"function": fn[1],
			"file":     file,
		}
		return fields
	}

	return log.Fields{}
}
