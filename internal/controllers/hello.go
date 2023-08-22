package controllers

import (
	"context"

	"github.com/ExcitingFrog/go-core-common/jaeger"
	pb "github.com/ExcitingFrog/xuyu/proto/gen/go/proto/api"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (c *Controllers) Hello(ctx context.Context, req *pb.HelloRequest) (*pb.HelloResponse, error) {
	ctx, span := jaeger.StartSpanFromContext(ctx, "Controllers:Hello")
	defer span.End()

	err := c.services.Hello(ctx)
	if err != nil {
		return nil, status.Errorf(codes.Internal, err.Error())
	}
	return &pb.HelloResponse{Message: "Hello"}, nil
}
