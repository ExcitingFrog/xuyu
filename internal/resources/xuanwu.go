package resources

import (
	"context"
	"net/http"

	"github.com/ExcitingFrog/go-core-common/jaeger"
	"github.com/ExcitingFrog/go-core-common/provider"
	xw "github.com/ExcitingFrog/xuanwu/swagger/client/client"
	"github.com/ExcitingFrog/xuanwu/swagger/client/client/operations"
	"github.com/ExcitingFrog/xuyu/configs"
	httptransport "github.com/go-openapi/runtime/client"
	"go.opentelemetry.io/contrib/instrumentation/net/http/otelhttp"
)

type Xuanwu struct {
	provider.IProvider

	client *xw.XuanWuService
}

func NewXuanwu() *Xuanwu {
	xuanwu := &Xuanwu{}
	config := configs.GetConfig()

	// xuanwuCfg := xw.DefaultTransportConfig().WithHost(config.XuyuHost).WithSchemes([]string{config.XuyuSchema})

	transport := httptransport.New(config.XuyuHost, "xuanwu/v1", []string{config.XuyuSchema})
	transport.Transport = otelhttp.NewTransport(http.DefaultTransport)

	// xuanwu.client = xw.NewHTTPClientWithConfig(nil, xuanwuCfg)
	xuanwu.client = xw.New(transport, nil)

	return xuanwu
}

func (x *Xuanwu) Hello(ctx context.Context) error {
	ctx, span, logger := jaeger.StartSpanAndLogFromContext(ctx, "Resources:Hello")
	defer span.End()

	ctx = context.WithValue(ctx, "token", "token")
	params := operations.NewHelloParams().WithContext(ctx)
	_, err := x.client.Operations.Hello(params)
	if err != nil {
		logger.Error(err.Error())
		return err
	}
	return err
}
