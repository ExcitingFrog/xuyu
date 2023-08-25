package services

import (
	"context"

	"github.com/ExcitingFrog/go-core-common/jaeger"
	"github.com/ExcitingFrog/xuyu/internal/schema"
	uuid "github.com/satori/go.uuid"
)

type IHello interface {
	Hello(ctx context.Context) error
	HelloTrace(ctx context.Context) error
}

func (s *Services) Hello(ctx context.Context) error {
	ctx, span, logger := jaeger.StartSpanAndLogFromContext(ctx, "Service:Hello")
	defer span.End()

	err := s.repository.SaveHello(ctx, &schema.Hello{
		ID: uuid.NewV4().String(),
	})
	if err != nil {
		logger.Error(err.Error())
		return err
	}
	return nil
}

func (s *Services) HelloTrace(ctx context.Context) error {
	ctx, span, logger := jaeger.StartSpanAndLogFromContext(ctx, "Service:HelloTrace")
	defer span.End()

	err := s.xuanwu.Hello(ctx)
	if err != nil {
		logger.Error(err.Error())
		return err
	}

	return err
}
