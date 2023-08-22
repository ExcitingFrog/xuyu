package main

import (
	"github.com/ExcitingFrog/go-core-common/grpc"
	"github.com/ExcitingFrog/go-core-common/grpc_gateway"
	"github.com/ExcitingFrog/go-core-common/jaeger"
	"github.com/ExcitingFrog/go-core-common/mongodb"
	"github.com/ExcitingFrog/go-core-common/pprof"
	"github.com/ExcitingFrog/go-core-common/provider"
	"github.com/ExcitingFrog/xuyu/internal/server"
)

func main() {
	stack := provider.NewProviders()

	// init mongodb
	mongodbProvider := mongodb.NewMongoDB(nil)
	stack.AddProvider(mongodbProvider)

	// init pprof
	pprofProvider := pprof.NewPprof(nil)
	stack.AddProvider(pprofProvider)

	// init jaeger
	jaegerProvider := jaeger.NewJaeger(nil)
	stack.AddProvider(jaegerProvider)

	// init grpc
	grpcProvider := grpc.NewGRpc(nil)
	grpcProvider.Init()
	stack.AddProvider(grpcProvider)

	// init grpc gateway
	gatewayProvider := grpc_gateway.NewGataway(nil, grpcProvider)
	gatewayProvider.Init()
	stack.AddProvider(gatewayProvider)

	// init server
	serverProvider := server.NewServer(grpcProvider, gatewayProvider, mongodbProvider, jaegerProvider)
	serverProvider.Init()
	stack.AddProvider(serverProvider)

	stack.Run()
}
