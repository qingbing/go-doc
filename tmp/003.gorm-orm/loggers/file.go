package loggers

import (
	"gorm.io/gorm/logger"
	"log"
	"os"
	"time"
)

func formatToday(sep string) string {
	return time.Now().Format("2006" + sep + "01" + sep + "02")
}

func New() (logger.Interface, error) {
	if file, err := os.OpenFile(formatToday("")+".out", os.O_CREATE|os.O_WRONLY|os.O_APPEND, os.ModePerm); err != nil {
		return nil, err
	} else {
		return logger.New(log.New(file, "", log.LstdFlags), logger.Config{
			SlowThreshold:             time.Second,  // 慢 SQL 阈值
			LogLevel:                  logger.Error, // 日志级别
			IgnoreRecordNotFoundError: true,         // 忽略ErrRecordNotFound（记录未找到）错误
			Colorful:                  false,        // 禁用彩色打印
		}), nil
	}
}
