package model

import "github.com/jinzhu/gorm"

type Member struct {
	gorm.Model
	Email    string `gorm:"type:char(32);comment:'邮箱'"`
	Nickname string `gorm:"type:char(32);comment:'用户昵称'"`
	Level    byte   `gorm:"type:tinyint(4);comment:'等级'"`
}
