package restaurantstorage

import (
	"context"
	"go-food-delivery/common"
	restaurantmodel "go-food-delivery/module/restaurant/model"
)

func (s *sqlStore) UpdateData(ctx context.Context, id int, data *restaurantmodel.RestaurantUpdate) error {
	_, err := s.FindDataWithCondition(ctx, map[string]interface{}{"id": id})
	if err != nil {
		return nil
	}

	if err := s.db.Where("id = ?", id).Updates(&data).Error; err != nil {
		return common.ErrDB(err)
	}
	return nil
}
