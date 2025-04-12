package restaurantstorage

import (
	"context"
	"go-food-delivery/common"
	restaurantmodel "go-food-delivery/module/restaurant/model"
)

func (s *sqlStore) CreateRestaurant(ctx context.Context, data *restaurantmodel.RestaurantCreate) error {
	if err := s.db.Create(&data).Error; err != nil {
		return common.ErrDB(err)
	}
	return nil
}
