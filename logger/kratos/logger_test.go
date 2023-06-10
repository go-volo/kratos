package kratos

import (
	"testing"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-volo/logger"
)

func TestDefaultLogger_Log(t *testing.T) {
	// 创建一个模拟的 Go Volo logger
	mockLogger := logger.DefaultLogger

	// 创建一个 defaultLogger 实例
	logger := New(mockLogger)

	// 设置预期的日志级别和消息
	level := log.LevelInfo
	message := "Test log message"

	// 调用 Log 方法进行日志记录
	err := logger.Log(level, message)
	if err != nil {
		t.Errorf("Log method returned an error: %v", err)
	}
}
