package ginrestaurant

import (
	"go-food-delivery/common"
	"go-food-delivery/component/appctx"
	restaurantbiz "go-food-delivery/module/restaurant/biz"
	restaurantstorage "go-food-delivery/module/restaurant/storage"
	"net/http"

	"github.com/gin-gonic/gin"
)

func FindDataWithCondition(appCtx appctx.AppContext) func(c *gin.Context) {
	return func(c *gin.Context) {
		uuid, err := common.FromBase58(c.Param("id"))
		if err != nil {
			c.JSON(http.StatusBadRequest, common.ErrInvalidRequest(err))
			return
		}
		store := restaurantstorage.NewSQLStore(appCtx.GetMaiDBConnection())
		biz := restaurantbiz.NewFindRestaurantBiz(store)
		data, err := biz.FindDataWithCondition(c.Request.Context(), int(uuid.GetLocalID()))
		if err != nil {
			c.JSON(http.StatusBadRequest, err)

			return
		}

		data.Mask(true)

		c.JSON(http.StatusOK, common.SimpleSuccessResponse(data))
	}
}
