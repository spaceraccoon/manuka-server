package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/spaceraccoon/manuka-server/controllers"
)

// SetupRouter creates the gin router
func SetupRouter() *gin.Engine {
	r := gin.Default()

	v1 := r.Group("/v1")
	{
		v1.GET("campaign", controllers.GetCampaigns)
		v1.POST("campaign", controllers.CreateCampaign)
		v1.GET("campaign/:id", controllers.GetCampaign)
		v1.PUT("campaign/:id", controllers.UpdateCampaign)
		v1.DELETE("campaign/:id", controllers.DeleteCampaign)
	}

	return r
}
