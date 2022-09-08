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
	logger.WithFields(logger.Fields{"Message": "Additional test message"}).Info("test message")
	logger.WithFields(logger.Fields{"Message": "Additional test message"}).Infof("test message %s", "special value")
	logger.WithFields(logger.Fields{"Message": "Additional test message"}).Debug("test message")
	logger.WithFields(logger.Fields{"Message": "Additional test message"}).Debugf("test message %s", "special value")
	logger.WithFields(logger.Fields{"Message": "Additional test message"}).Warn("test message")
	logger.WithFields(logger.Fields{"Message": "Additional test message"}).Warnf("test message %s", "special value")
	logger.WithFields(logger.Fields{"Message": "Additional test message"}).Error("test message")
	logger.WithFields(logger.Fields{"Message": "Additional test message"}).Errorf("test message %s", "special value")
}
