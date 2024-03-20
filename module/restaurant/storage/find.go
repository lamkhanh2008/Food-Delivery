package restaurantstorage

import (
	"context"
	"errors"
	"go-food-delivery/common"
	restaurantmodel "go-food-delivery/module/restaurant/model"

	"gorm.io/gorm"
)

func (s *sqlStore) FindDataWithCondition(ctx context.Context, condition map[string]interface{}, moreKeys ...string) (*restaurantmodel.Restaurant, error) {
	var data restaurantmodel.Restaurant

	if err := s.db.Where(condition).First(&data).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, errors.New("Record not found")
		}
		return nil, common.ErrDB(err)
	}

	return &data, nil
}
