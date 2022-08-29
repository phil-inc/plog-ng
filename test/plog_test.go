package test

import (
	"testing"

	logger "github.com/phil-inc/plog-ng/pkg/core"
)

func TestConnectivity(t *testing.T) {
	logger.Init()
	logger.Info("test message %s", "special value")
	logger.Debug("test message %s", "special value")
	logger.Error("test message %s", "special value")
	logger.Warn("test message %s", "special value")
}
