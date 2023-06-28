package dao

import (
	"blog-go/src/model"
	"log"
)

type MenuDao struct {
}

// GetById 根据id查询
func (m *MenuDao) GetById(id int) model.Menu {
	var menu model.Menu
	if err := db.Model(&model.Menu{}).Where(&model.Menu{Id: id}).First(&menu).Error; err != nil {
		log.Fatalf("MenuDao GetById, error: %v", err)
	}
	return menu
}

// Save 保存信息
func (m *MenuDao) Save(menu *model.Menu) bool {
	if err := db.Create(menu).Error; err != nil {
		log.Fatalf("MenuDao Save, error: %v", err)
		return false
	}
	return true
}

// Update 跟新信息
func (u MenuDao) Update(menu *model.Menu) bool {
	if err := db.Model(&model.Menu{}).Where(&model.Menu{Id: menu.Id}).Updates(&menu).Error; err != nil {
		log.Fatalf("MenuDao Update, error: %v", err)
		return false
	}
	return true
}

// List 数据列表
func (m MenuDao) List(menu *model.Menu) []model.Menu {
	var models []model.Menu
	if err := db.Where(menu).Find(&models).Error; err != nil {
		log.Fatalf("MenuDao List, error: %v", err)
	}
	return models
}
