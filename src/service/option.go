package service

import (
	"blog-go/src/dao"
	"blog-go/src/model"
)

type OptionService struct {
	optionDao *dao.OptionDao
}

// GetByKey 根据key查询
func (s *OptionService) GetByKey(key string) (model.Option, error) {
	return s.optionDao.GetByKey(key)
}

// GetById 根据id查询User
func (s *OptionService) GetById(id int) model.Option {
	return s.optionDao.GetById(id)
}
