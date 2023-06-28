package service

import (
	"blog-go/src/dao"
	"blog-go/src/model"
)

type MenuService struct {
	menuDao *dao.MenuDao
}

// GetById 根据id查询User
func (s *MenuService) GetById(id int) model.Menu {
	return s.menuDao.GetById(id)
}
