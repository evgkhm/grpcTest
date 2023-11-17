package grpc

import (
	"context"
	grpcmiddleware "github.com/grpc-ecosystem/go-grpc-middleware"
	grpcrecovery "github.com/grpc-ecosystem/go-grpc-middleware/recovery"
	"github.com/grpc-ecosystem/go-grpc-middleware/v2/interceptors/logging"
	"github.com/spf13/viper"
	"google.golang.org/grpc"
	grpcservice "grpcTest/internal/controller/grpc"
	"grpcTest/internal/controller/grpc/spec"
	"log/slog"
)

type Conf struct {
	Port     string
	HostPort string
}

var Config *Conf

func (h Conf) Init() {
	Config = &Conf{
		Port:     viper.GetString("grpc.port"),
		HostPort: viper.GetString("grpc.host_port"),
	}
}

// InterceptorLogger adapts slog logger to interceptor logger.
// This code is simple enough to be copied and not imported.
func InterceptorLogger(l *slog.Logger) logging.Logger {
	return logging.LoggerFunc(func(ctx context.Context, lvl logging.Level, msg string, fields ...any) {
		l.Log(ctx, slog.Level(lvl), msg, fields...)
	})
}

func New(services *grpcservice.Service, logger *slog.Logger) *grpc.Server {
	logger.Info("new grpc server")

	var opts = []grpc.ServerOption{
		grpc.StreamInterceptor(grpcmiddleware.ChainStreamServer(
			logging.StreamServerInterceptor(InterceptorLogger(logger)),
			grpcrecovery.StreamServerInterceptor(),
		)),
		grpc.UnaryInterceptor(grpcmiddleware.ChainUnaryServer(
			logging.UnaryServerInterceptor(InterceptorLogger(logger)),
			grpcrecovery.UnaryServerInterceptor(),
		)),
	}

	srv := grpc.NewServer(opts...)
	spec.RegisterGrpcTestAppServiceServer(srv, services)

	return srv
}
