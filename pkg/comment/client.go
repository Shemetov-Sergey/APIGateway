package comment

import (
	"fmt"

	"github.com/Shemetov-Sergey/APIGateway/pkg/comment/middleware"
	"github.com/Shemetov-Sergey/APIGateway/pkg/comment/pb"
	"github.com/Shemetov-Sergey/APIGateway/pkg/config"
	"google.golang.org/grpc"
)

type ServiceClient struct {
	Client pb.CommentServiceClient
}

func InitServiceClient(c *config.Config) pb.CommentServiceClient {
	// using WithInsecure() because no SSL running
	cc, err := grpc.Dial(c.CommentSvcUrl, grpc.WithInsecure(), middleware.WithClientUnaryInterceptor())

	if err != nil {
		fmt.Println("Could not connect:", err)
	}

	return pb.NewCommentServiceClient(cc)
}
