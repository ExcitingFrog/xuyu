package controllers

import (
	"context"

	"github.com/ExcitingFrog/go-core-common/utrace"
	pb "github.com/ExcitingFrog/xuyu/proto/gen/go/proto/api"
	trace_codes "go.opentelemetry.io/otel/codes"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (c *Controllers) Hello(ctx context.Context, req *pb.HelloRequest) (*pb.HelloResponse, error) {
	ctx, span, logger := utrace.StartSpanAndLogFromContext(ctx, "Controllers:Hello")
	defer span.End()

	err := c.services.Hello(ctx)
	if err != nil {
		logger.Error(err.Error())
		span.RecordError(err)
		span.SetStatus(trace_codes.Error, err.Error())
		return nil, status.Errorf(codes.Internal, err.Error())
	}
	return &pb.HelloResponse{Message: "Hello"}, nil
}

func (c *Controllers) HelloTrace(ctx context.Context, req *pb.HelloTraceRequest) (*pb.HelloTraceResponse, error) {
	ctx, span, logger := utrace.StartSpanAndLogFromContext(ctx, "Controllers:HelloTrace")
	defer span.End()

	err := c.services.HelloTrace(ctx)
	if err != nil {
		logger.Error(err.Error())
		span.RecordError(err)
		span.SetStatus(trace_codes.Error, err.Error())
		return nil, status.Errorf(codes.Internal, err.Error())
	}
	return &pb.HelloTraceResponse{Message: "Hello"}, nil
}
