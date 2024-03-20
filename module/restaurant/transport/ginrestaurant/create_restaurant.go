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

func CreateRestaurantdb(appCtx appctx.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		var data restaurantmodel.RestaurantCreate
		if err := c.ShouldBind(&data); err != nil {

			c.JSON(http.StatusBadRequest, common.ErrInvalidRequest(err))

			return
		}

		store := restaurantstorage.NewSQLStore(appCtx.GetMaiDBConnection())
		biz := restaurantbiz.NewCreateRestaurantBiz(store)
		if err := biz.CreateRestaurant(c.Request.Context(), &data); err != nil {
			c.JSON(http.StatusBadRequest, err)
			return
		}
		data.Mask(true)
		c.JSON(http.StatusOK, common.SimpleSuccessResponse(data.FakeId.String()))

	}
}
