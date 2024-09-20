package censor

import (
	"fmt"

	"github.com/Shemetov-Sergey/APIGateway/pkg/censor/middleware"
	"github.com/Shemetov-Sergey/APIGateway/pkg/censor/pb"
	"github.com/Shemetov-Sergey/APIGateway/pkg/config"
	"google.golang.org/grpc"
)

type ServiceClient struct {
	Client pb.CensorServiceClient
}

func InitServiceClient(c *config.Config) pb.CensorServiceClient {
	// using WithInsecure() because no SSL running
	cc, err := grpc.Dial(c.CensoredSvcUrl, grpc.WithInsecure(), middleware.WithClientUnaryInterceptor())

	if err != nil {
		fmt.Println("Could not connect:", err)
	}

	return pb.NewCensorServiceClient(cc)
}
