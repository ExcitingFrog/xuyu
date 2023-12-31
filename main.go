package main

import (
	"github.com/ExcitingFrog/go-core-common/grpc"
	"github.com/ExcitingFrog/go-core-common/grpc_gateway"
	"github.com/ExcitingFrog/go-core-common/log"
	"github.com/ExcitingFrog/go-core-common/mongodb"
	"github.com/ExcitingFrog/go-core-common/pprof"
	"github.com/ExcitingFrog/go-core-common/provider"
	"github.com/ExcitingFrog/go-core-common/utrace"
	"github.com/ExcitingFrog/xuyu/internal/server"
)

func main() {
	stack := provider.NewProviders()

	// init log
	logProvider := log.NewLog(nil)
	stack.AddProvider(logProvider)

	// init mongodb
	mongodbProvider := mongodb.NewMongoDB(nil)
	stack.AddProvider(mongodbProvider)

	// init pprof
	pprofProvider := pprof.NewPprof(nil)
	stack.AddProvider(pprofProvider)

	// init utrace
	traceProvider := utrace.NewUTrace(nil)
	stack.AddProvider(traceProvider)

	// init grpc
	grpcProvider := grpc.NewGRpc(nil)
	stack.AddProvider(grpcProvider)

	// init grpc gateway
	gatewayProvider := grpc_gateway.NewGataway(nil, grpcProvider)
	stack.AddProvider(gatewayProvider)

	// init server
	serverProvider := server.NewServer(grpcProvider, gatewayProvider, mongodbProvider)
	stack.AddProvider(serverProvider)

	stack.Run()
}
