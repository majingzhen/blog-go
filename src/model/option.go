package model

type Option struct {
	Id    int    `json:"id" gorm:"type:int;primary key"`
	Key   string `json:"key" gorm:"type:varchar(256);not null"`
	Value string `json:"value" gorm:"type:varchar(256)"`
}
