package api

import (
	"fmt"
	grpc_proto "kaspar/api/grpcprotos"
	handlers "kaspar/api/handlers"
	"kaspar/configuration"
	"kaspar/repository"
	usecase "kaspar/usecase/implementation"
	"net"

	"github.com/rs/zerolog/log"
	"google.golang.org/grpc"
)

func GrpcServe() {
	log.Debug().Msgf("Starting GRPC Server on %s", ":8081")

	cache := repository.NewRedis()
	stockRedditApi := usecase.NewStockRedditApi(cache)

	GRPC_PORT := configuration.GetEnvAsString("GRPC_PORT", "8081")
	lis, err := net.Listen("tcp", fmt.Sprintf(":%s", GRPC_PORT))
	if err != nil {
		log.Error().Msgf("Fail to Listen: %s", err)
	}
	s := grpc.NewServer()
	grpc_proto.RegisterStockHandleServer(s, &handlers.GrpcServer{StockRedditApi: *stockRedditApi})

	err = s.Serve(lis)

	if err != nil {
		log.Error().Msgf("Error while serving: %s", err)
	}
}
