package test

import (
	"fmt"
	"testing"

	logger "github.com/phil-inc/plog-ng/pkg/core"
)

func TestConnectivity(t *testing.T) {
	logger.Init()
	logger.Info(fmt.Sprintf("test message %s", "special value"))
	logger.Debug(fmt.Sprintf("test message %s", "special value"))
	logger.Error(fmt.Sprintf("test message %s", "special value"))
	logger.Warn(fmt.Sprintf("test message %s", "special value"))
}
