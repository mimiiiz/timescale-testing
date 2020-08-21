package service

import (
	"github.com/jinzhu/gorm"
	"github.com/mimiiiz/timescale-testing/model"
)

type locationService struct {
	db *gorm.DB
}
type LocationService interface {
	CreateLocation(location *model.Location) error
	ListLocation() ([]*model.Location, error)
}

func NewLocationService(db *gorm.DB) LocationService {
	return &locationService{db}
}

func (s *locationService) CreateLocation(location *model.Location) error {
	return s.db.Create(&location).Error

}

func (s *locationService) ListLocation() (locations []*model.Location, err error) {
	if err := s.db.Find(&locations).Error; err != nil {
		return locations, err
	}
	return locations, nil
}
