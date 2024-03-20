package ginrestaurant

import (
	"go-food-delivery/common"
	"go-food-delivery/component/appctx"
	restaurantbiz "go-food-delivery/module/restaurant/biz"
	restaurantstorage "go-food-delivery/module/restaurant/storage"
	"net/http"

	"github.com/gin-gonic/gin"
)

func DeleteRestaurant(appCtx appctx.AppContext) func(c *gin.Context) {
	return func(c *gin.Context) {

		uuid, err := common.FromBase58(c.Param("id"))
		if err != nil {
			panic(common.ErrInvalidRequest(err))
		}
		store := restaurantstorage.NewSQLStore(appCtx.GetMaiDBConnection())
		biz := restaurantbiz.NewDeleteRestaurantBiz(store)

		if err := biz.DeleteRestaurant(c.Request.Context(), int(uuid.GetLocalID())); err != nil {
			c.JSON(http.StatusBadRequest, err)
			return
		}
		c.JSON(http.StatusOK, common.SimpleSuccessResponse(true))
	}
}
