package gonews

import (
	"log"

	"github.com/Shemetov-Sergey/APIGateway/pkg/config"
	"github.com/Shemetov-Sergey/APIGateway/pkg/gonews/middleware"
	"github.com/Shemetov-Sergey/APIGateway/pkg/gonews/pb"
	"google.golang.org/grpc"
)

type ServiceClient struct {
	Client pb.GoNewsServiceClient
}

func InitServiceClient(c *config.Config) pb.GoNewsServiceClient {
	// using WithInsecure() because no SSL running
	cc, err := grpc.Dial(c.GoNewsSvcUrl, grpc.WithInsecure(), middleware.WithClientUnaryInterceptor())

	if err != nil {
		log.Println("Could not connect:", err)
	}

	return pb.NewGoNewsServiceClient(cc)
}
