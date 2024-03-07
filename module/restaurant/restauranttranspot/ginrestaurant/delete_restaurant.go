package ginrestaurant

import (
	"food/common"
	"food/component/appcontext"
	"food/module/restaurant/restaurantbusines"
	"food/module/restaurant/restaurantmodel"
	"food/module/restaurant/restaurentstorage"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func DeleteRestaurant(appCtx appcontext.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			c.JSON(401, map[string]interface{}{
				"error": err.Error(),
			})
		}
		var data restaurantmodel.RestaurantUpdate

		if err := c.ShouldBind(&data); err != nil {
			c.JSON(401, gin.H{
				"error": err.Error(),
			})
			return
		}
		store := restaurentstorage.NewSQLStore(appCtx.GetMainDBConnection())
		biz := restaurantbiz.NewDeleteRestaurantBiz(store)

		if err := biz.DeleteRestaurant(c.Request.Context(), id); err != nil {
			c.JSON(401, gin.H{
				"error": err.Error(),
			})
			return
		}

		c.JSON(http.StatusOK, common.SimpleSuccessResponse(true))
	}

}
