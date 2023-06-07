package dao

import (
	"blog-go/src/model"
	"log"
)

type UserDao struct {
}

// GetByAccount 根据账户查询User
func (u *UserDao) GetByAccount(account string) (model.User, error) {
	var user model.User
	err := db.Model(&model.User{}).Where(&model.User{Account: account}).First(&user).Error
	if err != nil {
		log.Fatalf("UserDao GetByAccount, error: %v", err)
	}
	return user, err
}

// GetById 根据id查询User
func (u *UserDao) GetById(id int) model.User {
	var user model.User
	if err := db.Model(&model.User{}).Where(&model.User{Id: id}).First(&user).Error; err != nil {
		log.Fatalf("UserDao GetById, error: %v", err)
	}
	return user
}

// Save 保存用户信息
func (u *UserDao) Save(user *model.User) bool {
	if err := db.Create(user).Error; err != nil {
		log.Fatalf("UserDao Save, error: %v", err)
		return false
	}
	return true
}

// Update 跟新用户信息
func (u UserDao) Update(user *model.User) bool {
	if err := db.Model(&model.User{}).Where(&model.User{Id: user.Id}).Updates(&user).Error; err != nil {
		log.Fatalf("UserDao Update, error: %v", err)
		return false
	}
	return true
}
