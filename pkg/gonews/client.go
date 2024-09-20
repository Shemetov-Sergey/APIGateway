package gonews

import (
	"fmt"

	"github.com/Shemetov-Sergey/APIGateway/pkg/config"
	"github.com/Shemetov-Sergey/APIGateway/pkg/gonews/pb"
	"google.golang.org/grpc"
)

type ServiceClient struct {
	Client pb.GoNewsServiceClient
}

func InitServiceClient(c *config.Config) pb.GoNewsServiceClient {
	// using WithInsecure() because no SSL running
	cc, err := grpc.Dial(c.GoNewsSvcUrl, grpc.WithInsecure())

	if err != nil {
		fmt.Println("Could not connect:", err)
	}

	return pb.NewGoNewsServiceClient(cc)
}
