package routes

//import (
//	"context"
//	"github.com/Shemetov-Sergey/APIGateway/pkg/models"
//	"net/http"
//
//	"github.com/Shemetov-Sergey/APIGateway/pkg/comment/pb"
//	"github.com/gin-gonic/gin"
//)
//
//func AddComment(ctx *gin.Context, c pb.CommentServiceClient) {
//	body := models.CreateCommentRequestBody{}
//
//	if err := ctx.BindJSON(&body); err != nil {
//		ctx.AbortWithError(http.StatusBadRequest, err)
//		return
//	}
//
//	userId, _ := ctx.Value("userId").(uint64)
//	in <- body
//	fullInfo := <-out
//
//	res, err := c.CreateComment(context.Background(), &pb.CreateCommentRequest{
//		NewsId:   fullInfo.NewsId,
//		ParentId: fullInfo.ParentId,
//		Text:     fullInfo.Text,
//		UserId:   userId,
//		Censored: fullInfo.Censored,
//	})
//
//	if err != nil {
//		ctx.AbortWithError(http.StatusBadGateway, err)
//		return
//	}
//
//	ctx.JSON(http.StatusCreated, &res)
//}
