package routes

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"server/controllers"
	"server/logger"
	"server/middleware"
)

func SetUp(mode string) *gin.Engine {
	if mode == gin.ReleaseMode {
		gin.SetMode(gin.ReleaseMode)
	}

	r := gin.New()

	r.Use(logger.GinLogger(), logger.GinRecovery(true))

	ds_visual := r.Group("/data")
	ds_visual.Use(middleware.Cors())

	{
		ds_visual.GET("/panel", controllers.Panel)
		ds_visual.GET("/indexs", controllers.StockIndex)
		ds_visual.GET("/history", controllers.History)
		ds_visual.GET("/current", controllers.Current)
		ds_visual.GET("/rank", controllers.Rank)
		ds_visual.GET("/total_asset", controllers.Asset)
	}

	r.NoRoute(func(c *gin.Context) {
		c.JSONP(http.StatusOK, gin.H{
			"msg": "404",
		})
	})

	return r

}
