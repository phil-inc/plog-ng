package core

import (
	log "github.com/sirupsen/logrus"
)

func Init() {
	log.SetFormatter(&log.TextFormatter{
		DisableColors: true,
		FullTimestamp: true,
	})
}

func Info(args interface{}) {
	log.Info(args)
}

func Debug(args interface{}) {
	log.Debug(args)
}

func Warn(args interface{}) {
	log.Warn(args)
}

func Error(args interface{}) {
	log.Error(args)
}
