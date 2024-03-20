package restaurantbiz

import (
	"context"
	"go-food-delivery/common"
	restaurantmodel "go-food-delivery/module/restaurant/model"
)

type ListRestaurantWithoutDelete interface {
	ListDataWithCondition(ctx context.Context, filter *restaurantmodel.Filter, paging *common.Paging, moreKeys ...string) ([]restaurantmodel.Restaurant, error)
}

type listRestaurantBiz struct {
	store ListRestaurantWithoutDelete
}

func NewListRestaurantBiz(store ListRestaurantWithoutDelete) listRestaurantBiz {
	return listRestaurantBiz{store: store}
}

func (biz *listRestaurantBiz) ListDataWithCondition(ctx context.Context, filter *restaurantmodel.Filter, paging *common.Paging, moreKeys ...string) ([]restaurantmodel.Restaurant, error) {
	list_result, err := biz.store.ListDataWithCondition(ctx, filter, paging)
	if err != nil {
		return nil, common.ErrCannotListEnity(restaurantmodel.EntityName, err)
	}
	return list_result, nil
}
