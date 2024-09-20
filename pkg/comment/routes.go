package comment

import (
	"github.com/Shemetov-Sergey/APIGateway/pkg/auth"
	"github.com/Shemetov-Sergey/APIGateway/pkg/comment/routes"
	"github.com/Shemetov-Sergey/APIGateway/pkg/config"
	"github.com/Shemetov-Sergey/APIGateway/pkg/models"
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(r *gin.Engine, c *config.Config, authSvc *auth.ServiceClient, in chan models.CreateCommentRequestBody, out chan models.CreateCommentRequestBody) {
	a := auth.InitAuthMiddleware(authSvc)

	svc := &ServiceClient{
		Client:     InitServiceClient(c),
		inComment:  in,
		outComment: out,
	}

	routesUrls := r.Group("/comment")
	routesUrls.Use(a.AuthRequired)
	routesUrls.POST("/", svc.AddComment)
	routesUrls.GET("/:news_id", svc.Comments)
}

func (svc *ServiceClient) AddComment(ctx *gin.Context) {
	routes.AddComment(ctx, svc.Client, svc.inComment, svc.outComment)
}

func (svc *ServiceClient) Comments(ctx *gin.Context) {
	routes.Comments(ctx, svc.Client)
}
