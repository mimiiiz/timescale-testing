package service

import (
	"github.com/jinzhu/gorm"
	"github.com/mimiiiz/timescale-testing/model"
)

type conditionService struct {
	db *gorm.DB
}
type ConditionService interface {
	CreateCondition(condition *model.Condition) error
	ListCondition() ([]*model.Condition, error)
}

func NewConditionService(db *gorm.DB) ConditionService {
	return &conditionService{db}
}

func (s *conditionService) CreateCondition(condition *model.Condition) error {
	return s.db.Create(&condition).Error

}

func (s *conditionService) ListCondition() (conditions []*model.Condition, err error) {
	if err := s.db.Find(&conditions).Error; err != nil {
		return conditions, err
	}
	return conditions, nil
}
