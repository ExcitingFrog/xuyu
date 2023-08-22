package services

import (
	"context"

	"github.com/ExcitingFrog/go-core-common/jaeger"
	"github.com/ExcitingFrog/xuyu/internal/schema"
	uuid "github.com/satori/go.uuid"
)

type IHello interface {
	Hello(ctx context.Context) error
}

func (s *Services) Hello(ctx context.Context) error {
	ctx, span := jaeger.StartSpanFromContext(ctx, "Service:Hello")
	defer span.End()

	s.repository.SaveHello(ctx, &schema.Hello{
		ID: uuid.NewV4().String(),
	})
	return nil
}
