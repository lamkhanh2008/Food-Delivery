package ginrestaurant

import (
	"go-food-delivery/component/appctx"
	restaurantbiz "go-food-delivery/module/restaurant/biz"
	restaurantmodel "go-food-delivery/module/restaurant/model"
	restaurantstorage "go-food-delivery/module/restaurant/storage"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func UpdateData(appCtx appctx.AppContext) func(c *gin.Context) {
	return func(c *gin.Context) {
		var data restaurantmodel.RestaurantUpdate
		if err := c.ShouldBind(&data); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}
		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}
		store := restaurantstorage.NewSQLStore(appCtx.GetMaiDBConnection())
		biz := restaurantbiz.NewUpdateRestaurantBiz(store)
		if err := biz.UpdateData(c.Request.Context(), id, &data); err != nil {

			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return

		}
		c.JSON(http.StatusOK, gin.H{
			"data": data,
		})
	}
}
