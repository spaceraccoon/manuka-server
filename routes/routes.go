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
		v1.GET("hit", controllers.GetHits)
		v1.POST("hit", controllers.CreateHit)
		v1.GET("hit/:id", controllers.GetHit)
		v1.DELETE("hit/:id", controllers.DeleteHit)
		v1.GET("honeypot", controllers.GetHoneypots)
		v1.GET("listener", controllers.GetListeners)
		v1.POST("listener", controllers.CreateListener)
		v1.GET("listener/:id", controllers.GetListener)
		v1.PUT("listener/:id", controllers.UpdateListener)
		v1.DELETE("listener/:id", controllers.DeleteListener)
		v1.GET("source", controllers.GetSources)
		v1.POST("source", controllers.CreateSource)
		v1.GET("source/:id", controllers.GetSource)
		v1.PUT("source/:id", controllers.UpdateSource)
		v1.DELETE("source/:id", controllers.DeleteSource)
	}

	return r
}
