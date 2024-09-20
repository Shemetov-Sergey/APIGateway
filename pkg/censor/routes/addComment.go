package routes

import (
	"context"
	"net/http"

	"github.com/Shemetov-Sergey/APIGateway/pkg/censor/pb"
	"github.com/Shemetov-Sergey/APIGateway/pkg/models"
	"github.com/gin-gonic/gin"
)

func AddComment(ctx *gin.Context, c pb.CensorServiceClient) {
	body := models.CreateCommentRequestBody{}

	if err := ctx.BindJSON(&body); err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	userId, _ := ctx.Value("userId").(uint64)

	res, err := c.CreateComment(context.Background(), &pb.CreateCommentRequest{
		NewsId:   body.NewsId,
		ParentId: body.ParentId,
		Text:     body.Text,
		UserId:   userId,
	})

	if err != nil {
		ctx.AbortWithError(http.StatusBadGateway, err)
		return
	}

	ctx.JSON(http.StatusCreated, &res)
}
