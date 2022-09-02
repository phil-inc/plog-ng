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
	logger.Infof("test message %s", "special value")
	logger.Debugf("test message %s", "special value")
	logger.Errorf("test message %s", "special value")
	logger.Warnf("test message %s", "special value")
}
