package restaurantbiz

import (
	"context"
	restaurantmodel "go-food-delivery/module/restaurant/model"
)

type UpdateRestaurantStore interface {
	UpdateData(ctx context.Context, id int, data *restaurantmodel.RestaurantUpdate) error
}

type updateRestaurantBiz struct {
	store UpdateRestaurantStore
}

func NewUpdateRestaurantBiz(store UpdateRestaurantStore) *updateRestaurantBiz {
	return &updateRestaurantBiz{store: store}
}

func (biz *updateRestaurantBiz) UpdateData(ctx context.Context, id int, data *restaurantmodel.RestaurantUpdate) error {
	if err := biz.store.UpdateData(ctx, id, data); err != nil {
		return err
	}
	return nil
}
