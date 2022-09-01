package test

import (
	"testing"

	logger "github.com/phil-inc/plog-ng/pkg/core"
)

func TestConnectivity(t *testing.T) {
	logger.Init()
	logger.Info("test message")
	logger.Debug("test message")
	logger.Error("test message")
	logger.Warn("test message")
}
