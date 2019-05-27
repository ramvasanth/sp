package controllers

import "github.com/gin-gonic/gin"

//LoadRoutes - loads the routes
func LoadRoutes(g *gin.Engine) {

	relation := g.Group("/relation")
	{
		relation.POST("/create", Relation{}.Create)
		relation.POST("/get", Relation{}.GetFriend)
		relation.POST("/common", Relation{}.GetCommonFriends)
	}

	subscription := g.Group("/subscription")
	{
		subscription.POST("/create", Subscription{}.Create)
		subscription.DELETE("/block", Subscription{}.Block)
	}
}

func healthCheck(gCxt *gin.Context) {

	gCxt.String(200, "ok")
}
