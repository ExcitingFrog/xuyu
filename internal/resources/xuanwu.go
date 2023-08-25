package resources

import (
	"context"

	"github.com/ExcitingFrog/go-core-common/jaeger"
	"github.com/ExcitingFrog/go-core-common/provider"
	xw "github.com/ExcitingFrog/xuanwu/swagger/client/client"
	"github.com/ExcitingFrog/xuanwu/swagger/client/client/operations"
	"github.com/ExcitingFrog/xuyu/configs"
)

type Xuanwu struct {
	provider.IProvider

	client *xw.XuanWuService
}

func NewXuanwu() *Xuanwu {
	xuanwu := &Xuanwu{}
	config := configs.GetConfig()

	xuanwuCfg := xw.DefaultTransportConfig().WithHost(config.XuyuHost).WithSchemes([]string{config.XuyuSchema})
	xuanwu.client = xw.NewHTTPClientWithConfig(nil, xuanwuCfg)

	return xuanwu
}

func (x *Xuanwu) Hello(ctx context.Context) error {
	ctx, span, logger := jaeger.StartSpanAndLogFromContext(ctx, "Resources:Hello")
	defer span.End()

	params := operations.NewHelloParams().WithContext(ctx)
	_, err := x.client.Operations.Hello(params)
	if err != nil {
		logger.Error(err.Error())
		return err
	}
	return err
}
