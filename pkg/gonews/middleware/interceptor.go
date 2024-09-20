package middleware

import (
	"context"
	"log"
	"time"

	"github.com/Shemetov-Sergey/APIGateway/pkg/config"
	"github.com/Shemetov-Sergey/APIGateway/pkg/gonews/pb"
	"github.com/renstrom/shortuuid"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
)

const (
	DefaultXRequestIDKey = "x-request-id"
	DefaultXRequestURL   = "x-service-address"
)

func clientInterceptor(
	ctx context.Context,
	method string,
	req interface{},
	reply interface{},
	cc *grpc.ClientConn,
	invoker grpc.UnaryInvoker,
	opts ...grpc.CallOption,
) error {
	c, err := config.LoadConfig()

	if err != nil {
		log.Fatalln("Failed at config", err)
	}
	// Logic before invoking the invoker
	start := time.Now()
	meta := metadata.New(map[string]string{})
	meta.Set(DefaultXRequestURL, c.GateWayAddr+c.Port)
	newCtx := SetRequestId(ctx, meta)
	// Calls the invoker to execute RPC
	err = invoker(newCtx, method, req, reply, cc, opts...)
	// Logic after invoking the invoker
	var responseStatus int64
	switch reply.(type) {
	case *pb.ListPostsResponse:
		replyValues := reply.(*pb.ListPostsResponse)
		responseStatus = replyValues.Status
	case *pb.PostsResponse:
		replyValues := reply.(*pb.PostsResponse)
		responseStatus = replyValues.Status
	case *pb.OnePostResponse:
		replyValues := reply.(*pb.OnePostResponse)
		responseStatus = replyValues.Status
	}
	log.Printf("Invoked RPC method=%s; destination=%s; requestId=%s; status=%d; Duration=%s; Error=%v",
		method,
		cc.Target(),
		meta.Get(DefaultXRequestIDKey)[0],
		responseStatus,
		time.Since(start),
		err)

	return err
}

func SetRequestId(ctx context.Context, meta metadata.MD) context.Context {
	requestId := HandleRequestID(ctx)
	ctx = metadata.NewOutgoingContext(ctx, meta)
	meta.Set(DefaultXRequestIDKey, requestId)
	return ctx
}

func WithClientUnaryInterceptor() grpc.DialOption {
	return grpc.WithUnaryInterceptor(clientInterceptor)
}

func HandleRequestID(ctx context.Context) string {
	md, ok := metadata.FromOutgoingContext(ctx)
	if !ok {
		return newRequestID()
	}

	header, ok := md[DefaultXRequestIDKey]
	if !ok || len(header) == 0 {
		return newRequestID()
	}

	requestID := header[0]
	if requestID == "" {
		return newRequestID()
	}

	return requestID
}

func newRequestID() string {
	return shortuuid.New()
}
