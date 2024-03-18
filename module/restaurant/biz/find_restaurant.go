package restaurantbiz

import (
	"context"
	restaurantmodel "go-food-delivery/module/restaurant/model"
)

type FindRestaurantWithCondition interface {
	FindDataWithCondition(ctx context.Context, condition map[string]interface{}, moreKeys ...string) (*restaurantmodel.Restaurant, error)
}

type findRestaurantBiz struct {
	store FindRestaurantWithCondition
}

func NewFindRestaurantBiz(store FindRestaurantWithCondition) *findRestaurantBiz {
	return &findRestaurantBiz{store: store}
}

func (biz *findRestaurantBiz) FindDataWithCondition(ctx context.Context, id int) (*restaurantmodel.Restaurant, error) {
	data, err := biz.store.FindDataWithCondition(ctx, map[string]interface{}{"id": id})
	if err != nil {

		return nil, err
	}

	return data, nil
}
