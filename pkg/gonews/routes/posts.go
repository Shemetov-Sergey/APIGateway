package routes

import (
	"context"
	"net/http"
	"strconv"

	"github.com/Shemetov-Sergey/APIGateway/pkg/gonews/pb"
	"github.com/gin-gonic/gin"
)

func Posts(ctx *gin.Context, c pb.GoNewsServiceClient) {
	newsCount, _ := strconv.ParseInt(ctx.Param("n"), 10, 64)

	res, err := c.Posts(context.Background(), &pb.PostsRequest{
		NewsCount: newsCount,
	})

	if err != nil {
		ctx.AbortWithError(http.StatusBadGateway, err)
		return
	}

	ctx.JSON(http.StatusOK, &res)
}

func NewsFullDetailed(ctx *gin.Context, c pb.GoNewsServiceClient) {
	newsId, _ := strconv.ParseInt(ctx.Param("news_id"), 10, 64)

	res, err := c.NewsFullDetailed(context.Background(), &pb.OneNewsRequest{
		NewsId: newsId,
	})

	if err != nil {
		ctx.AbortWithError(http.StatusBadGateway, err)
		return
	}

	ctx.JSON(http.StatusOK, &res)
}

func NewsShortDetailed(ctx *gin.Context, c pb.GoNewsServiceClient) {
	newsId, _ := strconv.ParseInt(ctx.Param("news_id"), 10, 64)

	res, err := c.NewsShortDetailed(context.Background(), &pb.OneNewsRequest{
		NewsId: newsId,
	})

	if err != nil {
		ctx.AbortWithError(http.StatusBadGateway, err)
		return
	}

	ctx.JSON(http.StatusOK, &res)
}

func FilterNews(ctx *gin.Context, c pb.GoNewsServiceClient) {
	filterValue := ctx.Param("filter_value")

	userId, err := strconv.ParseInt(ctx.Param("user_id"), 10, 64)
	if err != nil {
		ctx.AbortWithError(http.StatusBadGateway, err)
		return
	}

	pageSize, err := strconv.ParseInt(ctx.Param("page_size"), 10, 64)
	if err != nil {
		ctx.AbortWithError(http.StatusBadGateway, err)
		return
	}

	page, err := strconv.ParseInt(ctx.Param("page"), 10, 64)
	if err != nil {
		ctx.AbortWithError(http.StatusBadGateway, err)
		return
	}

	res, err := c.FilterNews(context.Background(), &pb.FilterNewsRequest{
		UserId:      userId,
		FilterValue: filterValue,
		PageSize:    int32(pageSize),
		Page:        int32(page),
	})

	if err != nil {
		ctx.AbortWithError(http.StatusBadGateway, err)
		return
	}

	ctx.JSON(http.StatusOK, &res)
}

func ListNews(ctx *gin.Context, c pb.GoNewsServiceClient) {
	newsCount, err := strconv.ParseInt(ctx.Param("news_count"), 10, 64)
	if err != nil {
		ctx.AbortWithError(http.StatusBadGateway, err)
		return
	}

	userId, err := strconv.ParseInt(ctx.Param("user_id"), 10, 64)
	if err != nil {
		ctx.AbortWithError(http.StatusBadGateway, err)
		return
	}

	pageSize, err := strconv.ParseInt(ctx.Param("page_size"), 10, 64)
	if err != nil {
		ctx.AbortWithError(http.StatusBadGateway, err)
		return
	}

	page, err := strconv.ParseInt(ctx.Param("page"), 10, 64)
	if err != nil {
		ctx.AbortWithError(http.StatusBadGateway, err)
		return
	}

	res, err := c.ListNews(context.Background(), &pb.ListPostsRequest{
		NewsCountGet: newsCount,
		UserId:       userId,
		PageSize:     int32(pageSize),
		Page:         int32(page),
	})

	if err != nil {
		ctx.AbortWithError(http.StatusBadGateway, err)
		return
	}

	ctx.JSON(http.StatusOK, &res)
}
