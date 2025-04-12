package ginrestaurant

import (
	"go-food-delivery/common"
	"go-food-delivery/component/appctx"
	restaurantbiz "go-food-delivery/module/restaurant/biz"
	restaurantmodel "go-food-delivery/module/restaurant/model"
	restaurantstorage "go-food-delivery/module/restaurant/storage"
	"net/http"

	"github.com/gin-gonic/gin"
)

func ListDataWithCondition(appCtx appctx.AppContext) func(c *gin.Context) {
	return func(c *gin.Context) {
		var list_result []restaurantmodel.Restaurant
		var fil restaurantmodel.Filter
		if err := c.ShouldBind(&fil); err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		var pagging common.Paging
		if err := c.ShouldBind(&pagging); err != nil {
			panic(common.ErrInvalidRequest(err))
		}
		pagging.Fullfill()
		store := restaurantstorage.NewSQLStore(appCtx.GetMaiDBConnection())
		biz := restaurantbiz.NewListRestaurantBiz(store)
		list_result, err := biz.ListDataWithCondition(c.Request.Context(), &fil, &pagging)
		if err != nil {
			panic(err)
		}
		for i := range list_result {
			list_result[i].Mask(true)
		}
		c.JSON(http.StatusOK, common.NewSuccessResponse(list_result, pagging, fil))
	}
}
