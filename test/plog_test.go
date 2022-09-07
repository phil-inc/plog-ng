package test

import (
	"testing"

	logger "github.com/phil-inc/plog-ng/pkg/core"
)

func TestConnectivity(t *testing.T) {
	logger.Init()
	logger.Entry.Info("test message")
	logger.Entry.Debug("test message")
	logger.Entry.Error("test message")
	logger.Entry.Warn("test message")
	logger.Entry.Infof("test message %s", "special value")
	logger.Entry.Debugf("test message %s", "special value")
	logger.Entry.Errorf("test message %s", "special value")
	logger.Entry.Warnf("test message %s", "special value")
	logger.WithFields(logger.Fields{"Message": "Additional test message"}).Info("test message")
	logger.WithFields(logger.Fields{"Message": "Additional test message"}).Infof("test message %s", "special value")
	logger.WithFields(logger.Fields{"Message": "Additional test message"}).Debug("test message")
	logger.WithFields(logger.Fields{"Message": "Additional test message"}).Debugf("test message %s", "special value")
	logger.WithFields(logger.Fields{"Message": "Additional test message"}).Warn("test message")
	logger.WithFields(logger.Fields{"Message": "Additional test message"}).Warnf("test message %s", "special value")
	logger.WithFields(logger.Fields{"Message": "Additional test message"}).Error("test message")
	logger.WithFields(logger.Fields{"Message": "Additional test message"}).Errorf("test message %s", "special value")
}
