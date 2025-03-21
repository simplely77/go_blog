package initialize

import (
	"github.com/robfig/cron/v3"
	"go.uber.org/zap"
	"os"
	"server/global"
	"server/task"
)

// ZapLogger 结构体实现了 cron.Logger 接口的 Info 和 Error 方法，这些方法用于接收 cron 包生成的日志并使用 zap 进行记录
type ZapLogger struct {
	logger *zap.Logger
}

func (l *ZapLogger) Info(msg string, keyAndValues ...interface{}) {
	l.logger.Info(msg, zap.Any("keyAndValues", keyAndValues))
}

func (l *ZapLogger) Error(err error, msg string, keyAndValues ...interface{}) {
	l.logger.Info(msg, zap.Error(err), zap.Any("keyAndValues", keyAndValues))
}

func NewZapLogger() *ZapLogger {
	return &ZapLogger{logger: global.Log}
}

// InitCron 初始化定时任务
func InitCorn() {
	c := cron.New(cron.WithLogger(NewZapLogger()))
	err := task.RegisterScheduledTasks(c)
	if err != nil {
		global.Log.Error("Error scheduling cron job:", zap.Error(err))
		os.Exit(1)
	}
	c.Start()
}
