package model

import "time"

type User struct {
	Id         int       `json:"id" gorm:"type:int;primary_key"`
	Account    string    `json:"account" gorm:"type:varchar(32);not null"`
	UserName   string    `json:"userName" gorm:"type:varchar(32);not null"`
	Password   string    `json:"password" gorm:"type:varchar(32);not null"`
	Salt       string    `json:"salt" gorm:"type:varchar(32);not null"`
	Status     int       `json:"status" gorm:"type:int;not null"`
	Avatar     string    `json:"avatar" gorm:"type:varchar(256)"`
	RoleId     int       `json:"roleId" gorm:"type:int;not null"`
	Email      string    `json:"email" gorm:"type:varchar(128)"`
	Website    string    `json:"website" gorm:"type:varchar(256)"`
	CreateTime time.Time `json:"createTime" gorm:"not null"`
	UpdateTime time.Time `json:"updateTime" gorm:"default:null"`
}
