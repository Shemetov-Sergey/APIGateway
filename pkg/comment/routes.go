package comment

import (
	"github.com/Shemetov-Sergey/APIGateway/pkg/auth"
	"github.com/Shemetov-Sergey/APIGateway/pkg/comment/routes"
	"github.com/Shemetov-Sergey/APIGateway/pkg/config"
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(r *gin.Engine, c *config.Config, authSvc *auth.ServiceClient) {
	a := auth.InitAuthMiddleware(authSvc)

	svc := &ServiceClient{
		Client: InitServiceClient(c),
	}

	routesUrls := r.Group("/comment")
	routesUrls.Use(a.AuthRequired)
	routesUrls.GET("/:news_id", svc.Comments)
}

func (svc *ServiceClient) Comments(ctx *gin.Context) {
	routes.Comments(ctx, svc.Client)
}
