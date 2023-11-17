package http

import (
	"context"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"github.com/spf13/viper"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"grpcTest/internal/controller/grpc/spec"
	"log/slog"
	"time"
)

var Config *Conf

type Conf struct {
	Port     string
	HostPort string
	LogPath  string
	Duration time.Duration
}

func (h Conf) Init() {
	Config = &Conf{
		Port:     viper.GetString("server.port"),
		HostPort: viper.GetString("server.host_port"),
		LogPath:  viper.GetString("server.log_path"),
		Duration: viper.GetDuration("server.duration"),
	}
}

func New(endpoint string, logger *slog.Logger) *runtime.ServeMux {
	logger.Info("new http server")
	mux := runtime.NewServeMux()

	opts := []grpc.DialOption{grpc.WithTransportCredentials(insecure.NewCredentials())}

	err := spec.RegisterGrpcTestAppServiceHandlerFromEndpoint(context.Background(), mux, endpoint, opts)
	if err != nil {
		logger.Error("new http server - RegisterGrpcTestAppServiceHandlerFromEndpoint", err)
	}

	return mux
}
