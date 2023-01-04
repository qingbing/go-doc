package pools

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
	"os"
	"time"
)

// 定义 db 连接池
// 特别注意: 使用连接池技术后，不能在使用完 db 后调用db.Close关闭数据库连接: 这样会导致整个数据库连接池关闭，导致连接池没有可用的连接

// 全局的 db 对象，实现单利模式
var _db *gorm.DB

func dailyCall(cb func()) {
	now := time.Now()
	nextSec := 86400 - now.Hour()*3600 - now.Minute()*60 - now.Second()
	if nextSec > 1 {
		cb()
	}
	time.AfterFunc(time.Second*time.Duration(nextSec), func() {
		go cb()
		for {
			select {
			case <-time.After(time.Hour * 24):
				go cb()
			}
		}
	})
}

func setLogger() {
	if file, err := os.OpenFile(time.Now().Format("20060102")+".out", os.O_CREATE|os.O_WRONLY|os.O_APPEND, os.ModePerm); err == nil {
		_db.Logger = logger.New(log.New(file, "", log.LstdFlags), logger.Config{
			SlowThreshold:             time.Second,  // 慢 SQL 阈值
			LogLevel:                  logger.Error, // 日志级别
			IgnoreRecordNotFoundError: true,         // 是否忽略ErrRecordNotFound（记录未找到）错误
			Colorful:                  false,        // 禁用彩色打印
		})
	} else {
		log.Fatal(err)
	}
}

// 初始化包，对 db 连接池进行初始化操纵
func init() {
	//配置MySQL连接参数
	username := "root"     //账号
	password := "123456"   //密码
	host := "127.0.0.1"    //数据库地址，可以是Ip或者域名
	port := 3306           //数据库端口
	Dbname := "beego_blog" //数据库名
	timeout := "10s"       //连接超时，10秒

	// 拼接 dsn 参数
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8&parseTime=True&loc=Local&timeout=%s", username, password, host, port, Dbname, timeout)
	// 连接 mysql, 获得DB类型实例，用于后面的数据库读写操作
	var err error
	_db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("连接数据库失败, error=" + err.Error())
	}
	if sqlDB, err := _db.DB(); err != nil {
		panic("获取sqlDB失败, error=" + err.Error())
	} else {
		// 设置数据库连接池参数
		sqlDB.SetMaxOpenConns(100) // 连接池最大连接数
		sqlDB.SetMaxIdleConns(20)  // 最大允许的空闲连接数，如果没有sql任务需要执行的连接数大于20，超过的连接会被连接池关闭
	}
	dailyCall(setLogger)
}

// GetDB 获取 gorm db 对象，在需要数据库操作时，直接使用该方法获取 db 连接对象即可，
// 不需要担心协程并发使用通用的 db 对象会共用同一个连接，db 对象在调用他的方法的时候会从数据库连接池中获取新的连接
func GetDB() *gorm.DB {
	return _db
}
