package main

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"io"
	"os"
)

/**
自定义 formatter， 只要实现以下即可
type Formatter interface {
	Format(*Entry) ([]byte, error)
}
*/
type myFormatter struct {
}

func (formatter myFormatter) Format(entry *logrus.Entry) ([]byte, error) {
	tString := entry.Time.Format("2006-01-02 15:04:05 ") + entry.Message + "\n"
	return []byte(tString), nil
}

type myHook struct {
}

// Levels 只接受 hook 的日志级别
func (hook myHook) Levels() []logrus.Level {
	return []logrus.Level{
		logrus.WarnLevel,
		logrus.ErrorLevel,
	}
}

// Fire 日志输出前调用的方法
func (hook myHook) Fire(entry *logrus.Entry) error {
	entry.Message = "{Hook}: " + entry.Message
	return nil
}

func main() {
	fmt.Println(logrus.GetLevel()) // 获取日志级别
	//logrus.SetLevel(logrus.WarnLevel) // 设置日志输出级别
	//logrus.SetReportCaller(true) // 日志输出调用者

	// 默认输出到 os.Stderr, 重定向输出到文件等地方，可以设置多个
	writer1 := os.Stdout
	writer2, _ := os.OpenFile("test.log", os.O_CREATE|os.O_APPEND|os.O_WRONLY, os.ModePerm)
	logrus.SetOutput(io.MultiWriter(writer1, writer2)) // 将在 标准输出的同时，也记录日志文件
	//logrus.SetFormatter(&logrus.JSONFormatter{})       // 设置日志的格式为 JSON
	logrus.SetFormatter(&myFormatter{})

	logrus.AddHook(&myHook{})

	logrus.Trace("trace msg")
	logrus.Println("print msg")
	logrus.Debug("debug msg")
	logrus.Info("info msg")
	logrus.Warn("warn msg")
	logrus.Error("error msg")
	//logrus.Fatal("fatal msg")
	//logrus.Panic("panic msg")
	logrus.WithField("name", "qingbing").WithField("age", 19).Warnln("WithField Warning") // 打印带字段信息
	logrus.WithFields(logrus.Fields{
		"name": "qingbing",
		"age":  18,
	}).Warnln("WithFields Warning")
}
