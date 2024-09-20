package comment

import (
	"fmt"

	"github.com/Shemetov-Sergey/APIGateway/pkg/comment/pb"
	"github.com/Shemetov-Sergey/APIGateway/pkg/config"
	"github.com/Shemetov-Sergey/APIGateway/pkg/models"
	"google.golang.org/grpc"
)

type ServiceClient struct {
	Client     pb.CommentServiceClient
	inComment  chan models.CreateCommentRequestBody
	outComment chan models.CreateCommentRequestBody
}

func InitServiceClient(c *config.Config) pb.CommentServiceClient {
	// using WithInsecure() because no SSL running
	cc, err := grpc.Dial(c.CommentSvcUrl, grpc.WithInsecure())

	if err != nil {
		fmt.Println("Could not connect:", err)
	}

	return pb.NewCommentServiceClient(cc)
}
