package gonews

import (
	"github.com/Shemetov-Sergey/APIGateway/pkg/auth"
	"github.com/Shemetov-Sergey/APIGateway/pkg/config"
	"github.com/Shemetov-Sergey/APIGateway/pkg/gonews/routes"
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(r *gin.Engine, c *config.Config, authSvc *auth.ServiceClient) {
	a := auth.InitAuthMiddleware(authSvc)

	svc := &ServiceClient{
		Client: InitServiceClient(c),
	}

	routesGroup := r.Group("/gonews")
	routesGroup.Use(a.AuthRequired)
	routesGroup.GET("/:n", svc.Posts)
	fullNews := routesGroup.Group("/news_full")
	fullNews.GET("/:news_id", svc.NewsFullDetailed)
	shortNews := routesGroup.Group("/news_short")
	shortNews.GET("/:news_id", svc.NewsShortDetailed)
	filteredNews := routesGroup.Group("/filtered_news")
	filteredNews.GET("/:filter_value", svc.FilterNews)
	ListNews := routesGroup.Group("/list_news_pages")
	ListNews.GET("/news_count=:news_count/user=:user_id/page_size=:page_size/page=:page", svc.ListNews)
}

func (svc *ServiceClient) Posts(ctx *gin.Context) {
	routes.Posts(ctx, svc.Client)
}

func (svc *ServiceClient) NewsFullDetailed(ctx *gin.Context) {
	routes.NewsFullDetailed(ctx, svc.Client)
}

func (svc *ServiceClient) NewsShortDetailed(ctx *gin.Context) {
	routes.NewsShortDetailed(ctx, svc.Client)
}

func (svc *ServiceClient) FilterNews(ctx *gin.Context) {
	routes.FilterNews(ctx, svc.Client)
}

func (svc *ServiceClient) ListNews(ctx *gin.Context) {
	routes.ListNews(ctx, svc.Client)
}
