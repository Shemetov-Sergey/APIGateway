package auth

import (
	"github.com/Shemetov-Sergey/APIGateway/pkg/auth/routes"
	"github.com/Shemetov-Sergey/APIGateway/pkg/config"
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(r *gin.Engine, c *config.Config) *ServiceClient {
	svc := &ServiceClient{
		Client: InitServiceClient(c),
	}

	authRoutes := r.Group("/auth")
	authRoutes.POST("/register", svc.Register)
	authRoutes.POST("/login", svc.Login)

	return svc
}

func (svc *ServiceClient) Register(ctx *gin.Context) {
	routes.Register(ctx, svc.Client)
}

func (svc *ServiceClient) Login(ctx *gin.Context) {
	routes.Login(ctx, svc.Client)
}
