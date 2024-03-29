package ginrestaurant

import (
	"food/common"
	"food/component/appcontext"
	"food/module/restaurant/restaurantbusines"
	"food/module/restaurant/restaurantmodel"
	"food/module/restaurant/restaurentstorage"
	"github.com/gin-gonic/gin"
	"net/http"
)

func ListRestaurant(appCtx appcontext.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		var filter restaurantmodel.Filter
		if err := c.ShouldBind(&filter); err != nil {
			c.JSON(401, gin.H{
				"error": err.Error(),
			})
			return
		}
		var paging common.Paging
		if err := c.ShouldBind(&paging); err != nil {
			c.JSON(401, gin.H{
				"error": err.Error(),
			})
			return
		}
		paging.Fulfill()
		store := restaurentstorage.NewSQLStore(appCtx.GetMainDBConnection())
		biz := restaurantbiz.NewListRestaurantBiz(store)
		result, err := biz.ListRestaurant(c.Request.Context(), &filter, &paging)
		if err != nil {
			c.JSON(401, gin.H{
				"error": err.Error(),
			})
			return
		}

		c.JSON(http.StatusOK, common.NewSuccessResponse(result, paging, filter))
	}

}
