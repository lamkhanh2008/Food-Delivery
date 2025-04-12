package restaurantbiz

import (
	"context"
	"go-food-delivery/common"
	restaurantmodel "go-food-delivery/module/restaurant/model"
)

type DeleteRestaurantStore interface {
	DeleteRestaurant(ctx context.Context, id int) error
	FindDataWithCondition(ctx context.Context, condition map[string]interface{}, moreKeys ...string) (*restaurantmodel.Restaurant, error)
}

type deleteRestaurantBiz struct {
	store DeleteRestaurantStore
}

func NewDeleteRestaurantBiz(store DeleteRestaurantStore) *deleteRestaurantBiz {
	return &deleteRestaurantBiz{store: store}
}

func (biz *deleteRestaurantBiz) DeleteRestaurant(ctx context.Context, id int) error {
	olddata, err := biz.store.FindDataWithCondition(ctx, map[string]interface{}{"id": id})
	if err != nil {
		return common.ErrEntityNotFound(restaurantmodel.EntityName, err)
	}
	if olddata.Status == 0 {
		return common.ErrEntityDeleted(restaurantmodel.EntityName, err)
	}
	if err := biz.store.DeleteRestaurant(ctx, id); err != nil {
		return common.ErrCannotCreateEntity(restaurantmodel.EntityName, err)
	}
	return nil
}
