package main

import (
	"fmt"
	"go-doc/tmp/003.gorm-orm/models"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
	"os"
	"time"
)

func main() {
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
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		//Logger: log,
	})

	if err != nil {
		panic("连接数据库失败, error=" + err.Error())
	}
	// 延时关闭数据库连接
	if sqlDb, err := db.DB(); err == nil {
		defer sqlDb.Close()
	}

	go func() {
		for {
			select {
			case <-time.After(time.Hour):
				if file, err := os.OpenFile(time.Now().Format("20060102030405")+".out", os.O_CREATE|os.O_WRONLY|os.O_APPEND, os.ModePerm); err == nil {
					db.Logger = logger.New(log.New(file, "", log.LstdFlags), logger.Config{
						SlowThreshold:             time.Second,  // 慢 SQL 阈值
						LogLevel:                  logger.Error, // 日志级别
						IgnoreRecordNotFoundError: true,         // 是否忽略ErrRecordNotFound（记录未找到）错误
						Colorful:                  false,        // 禁用彩色打印
					})
				}
			}
		}
	}()

	for {
		select {
		case <-time.After(time.Second * 2):
			// insert
			user := models.User{
				Email:    "480042175@qq.com",
				Nickname: "4ingbing",
				RealName: "qb",
			}
			if err := db.Create(&user).Error; err != nil {
				fmt.Println("Create Error", err)
			}

			// select
			user1 := models.User{}
			if err := db.Where("nickname = ?", "qinging").First(&user1).Error; err != nil {
				fmt.Println("error")
				// 处理错误...
			} else {
				fmt.Printf("%+v\n", user1)
			}
		}
	}
}
