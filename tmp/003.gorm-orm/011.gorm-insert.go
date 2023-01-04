package main

import (
	"fmt"
	"go-doc/tmp/003.gorm-orm/models"
	"go-doc/tmp/003.gorm-orm/pools"
)

func main() {
	db := pools.GetDB()
	fmt.Println(db)

	user := models.User{
		Email:    "48004217@qq.com",
		Nickname: "1ingbing",
		RealName: "qb",
	}
	if err := db.Create(&user).Error; err != nil {
		fmt.Println("Create Error", err)
	} else {
		fmt.Printf("ID:%d Email:%s Nickname:%s RealName: %s", user.Uid, user.Email, user.Nickname, user.RealName)
	}

	//now := time.Now()
	//nextSec := 3600 - now.Minute()*60 + now.Second() + (24-now.Hour())*3600
	//fmt.Println(time.Now().Hour(), nextSec)
	//fmt.Println(86400 - now.Hour()*3600 - now.Minute()*60 - now.Second())
}
