package v1

import (
	"context"
	"fmt"
	grpc_opentracing "github.com/grpc-ecosystem/go-grpc-middleware/tracing/opentracing"
	"go.opencensus.io/plugin/ocgrpc"
	gerror "golang.swit.dev/grpc/interceptor/error"
	"google.golang.org/grpc"
)

type Address struct {
	IP   string
	Port string
}

func UserClientProvider() AuthServiceClient {
	a := Address{
		IP:   "localhost",
		Port: "50053",
	}

	add := fmt.Sprintf("%s:%s", a.IP, a.Port)
	conn, err := grpc.DialContext(context.Background(), add,
		grpc.WithChainUnaryInterceptor(
			gerror.NewClientErrorInterceptor(),
			grpc_opentracing.UnaryClientInterceptor(),
		),
		grpc.WithDefaultCallOptions(grpc.MaxCallRecvMsgSize(104857600), grpc.MaxCallSendMsgSize(104857600)),
		grpc.WithInsecure(),
		grpc.WithStatsHandler(&ocgrpc.ClientHandler{}),
	)

	if err != nil {
		panic(err)
	}

	return NewAuthServiceClient(conn)
}
