package service

import (
	"blog-go/src/dao"
	"blog-go/src/model"
)

type UserService struct {
	userDao *dao.UserDao
}

// GetByAccount 根据账户查询User
func (s *UserService) GetByAccount(account string) (model.User, error) {
	return s.userDao.GetByAccount(account)
}

// GetById 根据id查询User
func (s *UserService) GetById(id int) model.User {
	return s.userDao.GetById(id)
}
