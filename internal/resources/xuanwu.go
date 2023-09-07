package resources

import (
	"context"
	"net/http"

	"github.com/ExcitingFrog/go-core-common/provider"
	"github.com/ExcitingFrog/go-core-common/utrace"
	xw "github.com/ExcitingFrog/go-proto-lib/rest/xuanwu/swagger/client/client"
	"github.com/ExcitingFrog/go-proto-lib/rest/xuanwu/swagger/client/client/operations"
	"github.com/ExcitingFrog/xuyu/configs"
	httptransport "github.com/go-openapi/runtime/client"
	"go.opentelemetry.io/contrib/instrumentation/net/http/otelhttp"
	trace_codes "go.opentelemetry.io/otel/codes"
)

type Xuanwu struct {
	provider.IProvider

	client *xw.XuanWuService
}

func NewXuanwu() *Xuanwu {
	xuanwu := &Xuanwu{}
	config := configs.GetConfig()

	transport := httptransport.New(config.XuyuHost, "xuanwu/v1", []string{config.XuyuSchema})
	transport.Transport = otelhttp.NewTransport(http.DefaultTransport)

	xuanwu.client = xw.New(transport, nil)

	return xuanwu
}

func (x *Xuanwu) Hello(ctx context.Context) error {
	ctx, span, logger := utrace.StartSpanAndLogFromContext(ctx, "Resources:Hello")
	defer span.End()

	params := operations.NewHelloParams().WithContext(ctx)
	_, err := x.client.Operations.Hello(params)
	if err != nil {
		logger.Error(err.Error())
		span.RecordError(err)
		span.SetStatus(trace_codes.Error, err.Error())
		return err
	}
	return err
}
