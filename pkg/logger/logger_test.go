package logger

import (
	"com.gientech/equipment-data-collection/pkg/config"
	"go.uber.org/zap"
	"testing"
)

// TestZap zap 日志框架
func TestZap(t *testing.T) {
	var config = config.NewConfig("../../config")
	zapLog := newZapLog(config)
	zapLog.Debug("debug message")
	zapLog.Info("info message")
	zapLog.Warn("warn message")
	zapLog.Error("error message")
}

func TestLogger(t *testing.T) {
	var config = config.NewConfig("../../config")
	InitLogger(config)
	S().Debugf("debug message")
	S().Infof("info message")
	S().Warnf("warn message")
	S().Errorf("error message:%s", "this is message")
	zap.S().Infof("xxxxxxxxxx:%s","hello")
}
