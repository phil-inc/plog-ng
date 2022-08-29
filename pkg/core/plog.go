package core

import (
	log "github.com/sirupsen/logrus"
)

func Init() {
	log.SetFormatter(&log.TextFormatter{
		DisableColors: true,
		FullTimestamp: true,
	})
	//log.SetLevel(log.DebugLevel) //Note this must be enabled from the application
}

func Info(args ...interface{}) {
	log.Info(args)
}

func Debug(args ...interface{}) {
	log.Debug(args)
}

func Warn(args ...interface{}) {
	log.Warn(args)
}

func Error(args ...interface{}) {
	log.Error(args)
}
