package repository

import (
	"context"

	"github.com/ExcitingFrog/go-core-common/jaeger"
	"github.com/ExcitingFrog/xuyu/internal/schema"
)

type IHello interface {
	SaveHello(ctx context.Context, h *schema.Hello) error
}

const helloCollection = "hello"
const testDB = "test"

func (r *repository) SaveHello(ctx context.Context, h *schema.Hello) error {
	ctx, span := jaeger.StartSpanFromContext(ctx, "Repository:Hello")
	defer span.End()

	_, err := r.mongo.Client.Database(testDB).Collection(helloCollection).InsertOne(ctx, h)
	if err != nil {
		return err
	}
	return nil
}
