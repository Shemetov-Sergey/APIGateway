package censor

import (
	"github.com/Shemetov-Sergey/APIGateway/pkg/auth"
	"github.com/Shemetov-Sergey/APIGateway/pkg/censor/routes"
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
	routesUrls.POST("/", svc.AddComment)
}

func (svc *ServiceClient) AddComment(ctx *gin.Context) {
	routes.AddComment(ctx, svc.Client)
}
