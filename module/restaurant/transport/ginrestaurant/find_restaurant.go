package ginrestaurant

import (
	"go-food-delivery/component/appctx"
	restaurantbiz "go-food-delivery/module/restaurant/biz"
	restaurantstorage "go-food-delivery/module/restaurant/storage"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func FindDataWithCondition(appCtx appctx.AppContext) func(c *gin.Context) {
	return func(c *gin.Context) {
		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}
		store := restaurantstorage.NewSQLStore(appCtx.GetMaiDBConnection())
		biz := restaurantbiz.NewFindRestaurantBiz(store)
		data, err := biz.FindDataWithCondition(c.Request.Context(), id)
		if err != nil {
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
