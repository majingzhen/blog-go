package model

import "time"

type Menu struct {
	Id         int       `json:"id" gorm:"type:int;primary key"`
	Pid        int       `json:"pid" gorm:"type:int;"`
	Name       string    `json:"name" gorm:"type:varchar(128);not null"`
	Url        string    `json:"url" gorm:"type:varchar(128)"`
	Icon       string    `json:"icon" gorm:"type:varchar(64)"`
	Seq        int       `json:"seq" gorm:"type:int"`
	Type       int       `json:"type" gorm:"type:int"`
	Target     int       `json:"target" gorm:"type:int"`
	Status     int       `json:"status" gorm:"type:int"`
	CreateTime time.Time `json:"createTime" gorm:"not null"`
	UpdateTime time.Time `json:"updateTime" gorm:"default null"`
}
