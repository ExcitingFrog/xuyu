package server

import (
	"context"
	"fmt"

	"github.com/ExcitingFrog/go-core-common/grpc"
	"github.com/ExcitingFrog/go-core-common/grpc_gateway"
	"github.com/ExcitingFrog/go-core-common/mongodb"
	"github.com/ExcitingFrog/go-core-common/provider"
	"github.com/ExcitingFrog/xuyu/internal/controllers"
	"github.com/ExcitingFrog/xuyu/internal/repository"
	"github.com/ExcitingFrog/xuyu/internal/resources"
	"github.com/ExcitingFrog/xuyu/internal/services"
	pb "github.com/ExcitingFrog/xuyu/proto/gen/go/proto/api"
)

type Server struct {
	provider.IProvider

	controllers *controllers.Controllers
	grpc        *grpc.GRpc
	gateway     *grpc_gateway.Gataway
	mongodb     *mongodb.MongoDB
}

func NewServer(grpc *grpc.GRpc, gateway *grpc_gateway.Gataway, mongodb *mongodb.MongoDB) *Server {

	repository := repository.NewRepository(mongodb)

	xuanwu := resources.NewXuanwu()

	services := services.NewServices(repository, xuanwu)

	controllers := controllers.NewControllers(services)

	return &Server{
		controllers: controllers,
		grpc:        grpc,
		gateway:     gateway,
		mongodb:     mongodb,
	}
}
func (s *Server) Init() error {

	pb.RegisterHelloAPIServer(s.grpc.Server, s.controllers)

	return nil
}

func (s *Server) Run() error {

	ctx := context.Background()
	// register gateway
	pb.RegisterHelloAPIHandlerFromEndpoint(ctx, s.gateway.Mux, fmt.Sprintf(":%d", s.grpc.Config.ServerPort), s.gateway.Options())

	return nil
}
