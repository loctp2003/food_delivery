package ginrestaurant

import (
	"food/common"
	"food/component/appcontext"
	"food/module/restaurant/restaurantbusines"
	"food/module/restaurant/restaurentstorage"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func GetRestaurant(appCtx appcontext.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			c.JSON(401, gin.H{
				"error": err.Error(),
			})

			return
		}

		store := restaurentstorage.NewSQLStore(appCtx.GetMainDBConnection())
		biz := restaurantbiz.NewGetRestaurantBiz(store)

		data, err := biz.GetRestaurant(c.Request.Context(), id)

		if err != nil {
			c.JSON(401, map[string]interface{}{
				"error": err.Error(),
			})

			return
		}

		c.JSON(http.StatusOK, common.SimpleSuccessResponse(data))
	}
}
