package grpc

import (
	"grpcTest/internal/controller/grpc/spec"
	"grpcTest/internal/usecase"
	"log/slog"
)

type Service struct {
	spec.UnimplementedGrpcTestAppServiceServer
	useCase usecase.UseCase
	logger  *slog.Logger
}

func New(useCase *usecase.UseCase, logger *slog.Logger) *Service {
	return &Service{
		UnimplementedGrpcTestAppServiceServer: spec.UnimplementedGrpcTestAppServiceServer{},
		useCase:                               *useCase,
		logger:                                logger,
	}
}
