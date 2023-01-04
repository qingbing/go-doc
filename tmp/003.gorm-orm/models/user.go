package models

import (
	"time"
)

type User struct {
	Uid         uint   `gorm:"primaryKey"`
	Email       string `gorm:"unique"`
	Password    string
	Nickname    string `gorm:"unique"`
	RealName    string `gorm:"index"`
	IsEnable    bool
	RegisterIp  string
	RegisterAt  time.Time `gorm:"autoCreateTime"`
	LoginTimes  uint
	LastLoginIp string
	LastLoginAt time.Time `gorm:"default:1000-01-01 00:00:00"`
	UpdatedAt   time.Time `gorm:"autoUpdateTime"`
}
