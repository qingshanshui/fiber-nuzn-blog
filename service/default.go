package service

import (
	"fiber-nuzn-blog/models"
	"fiber-nuzn-blog/validator/form"
)

type Default struct{}

func NewDefaultService() *Default {
	return &Default{}
}

// Home 首页
func (t *Default) Home() ([]models.Course, error) {
	list, err := models.NewCourse().List()
	if err != nil {
		return nil, err
	}
	return list, nil
}

// Category 详情
func (t *Default) Category(c form.CategoryRequest) (*models.Course, error) {
	list, err := models.NewCourse().Category(c.Id)
	if err != nil {
		return nil, err
	}
	return list, nil
}
