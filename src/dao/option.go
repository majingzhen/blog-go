package dao

import (
	"blog-go/src/model"
	"log"
)

type OptionDao struct {
}

// GetByKey 根据key查询
func (o *OptionDao) GetByKey(key string) (model.Option, error) {
	var option model.Option
	err := db.Model(&model.Option{}).Where(&model.Option{Key: key}).First(&option).Error
	if err != nil {
		log.Fatalf("OptionDao GetByKey, error: %v", err)
	}
	return option, err
}

// GetById 根据id查询
func (o *OptionDao) GetById(id int) model.Option {
	var option model.Option
	if err := db.Model(&model.Option{}).Where(&model.Option{Id: id}).First(&option).Error; err != nil {
		log.Fatalf("OptionDao GetById, error: %v", err)
	}
	return option
}

// Save 保存信息
func (o *OptionDao) Save(option *model.Option) bool {
	if err := db.Create(option).Error; err != nil {
		log.Fatalf("OptionDao Save, error: %v", err)
		return false
	}
	return true
}

// Update 跟新信息
func (o OptionDao) Update(option *model.Option) bool {
	if err := db.Model(&model.Option{}).Where(&model.Option{Id: option.Id}).Updates(&option).Error; err != nil {
		log.Fatalf("OptionDao Update, error: %v", err)
		return false
	}
	return true
}
