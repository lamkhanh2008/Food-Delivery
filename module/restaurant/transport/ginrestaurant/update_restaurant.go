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

func UpdateData(appCtx appctx.AppContext) func(c *gin.Context) {
	return func(c *gin.Context) {
		var data restaurantmodel.RestaurantUpdate
		if err := c.ShouldBind(&data); err != nil {
			c.JSON(http.StatusBadRequest, common.ErrInvalidRequest(err))
			return
		}
		uuid, err := common.FromBase58(c.Param("id"))
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}
		store := restaurantstorage.NewSQLStore(appCtx.GetMaiDBConnection())
		biz := restaurantbiz.NewUpdateRestaurantBiz(store)
		if err := biz.UpdateData(c.Request.Context(), int(uuid.GetLocalID()), &data); err != nil {

			c.JSON(http.StatusBadRequest, err)
			return

		}
		c.JSON(http.StatusOK, gin.H{
			"data": data,
		})
	}
}
