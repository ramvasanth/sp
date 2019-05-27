package controllers

import "github.com/gin-gonic/gin"

//LoadRoutes - loads the routes
func LoadRoutes(g *gin.Engine) {

	campaign := g.Group("/campaign")
	{
		campaign.POST("", Campaign{}.Create)
		campaign.PUT("", Campaign{}.Update)
		campaign.GET("/:id", Campaign{}.Get)
		campaign.DELETE("/:id", Campaign{}.Delete)
	}
	g.GET("/campaigns", Campaign{}.List)

	promo := g.Group("/promo")
	{

		promo.GET("", Promo{}.Create)
		promo.DELETE("", Promo{}.Delete)
	}

	g.GET("/participation", Participation{}.Update)
	g.GET("/participations", Participation{}.List)

	g.GET("/healthcheck", healthCheck)

}

func healthCheck(gCxt *gin.Context) {

	gCxt.String(200, "ok")
}