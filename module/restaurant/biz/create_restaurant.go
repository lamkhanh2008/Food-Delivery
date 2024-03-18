package restaurantbiz

import (
	"context"
	restaurantmodel "go-food-delivery/module/restaurant/model"
)

type CreateRestaurantStore interface {
	CreateRestaurant(ctx context.Context, data *restaurantmodel.RestaurantCreate) error
}

type createRestaurantBiz struct {
	store CreateRestaurantStore
}

func NewCreateRestaurantBiz(store CreateRestaurantStore) *createRestaurantBiz {
	return &createRestaurantBiz{store: store}
}

func (biz *createRestaurantBiz) CreateRestaurant(ctx context.Context, data *restaurantmodel.RestaurantCreate) error {
	if err := biz.store.CreateRestaurant(ctx, data); err != nil {
		return err
	}
	return nil
}
