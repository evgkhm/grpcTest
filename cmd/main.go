package main

import (
	"context"
	"golang.org/x/sync/errgroup"
	"grpcTest/internal/config"
	"grpcTest/internal/controller/grpc"
	"grpcTest/internal/usecase"
	"grpcTest/pkg/logging"
	grpcserver "grpcTest/pkg/server/grpc"
	httpserver "grpcTest/pkg/server/http"
	"net"
	"net/http"
)

func init() {
	config.InitAll([]config.Config{
		httpserver.Conf{},
		grpcserver.Conf{},
	})
}

func main() {
	logger := logging.Logger()

	useCases := usecase.New(logger)

	services := grpc.New(useCases, logger)

	grpcServer := grpcserver.New(services, logger)

	listen, err := net.Listen("tcp", grpcserver.Config.HostPort)
	if err != nil {
		logger.Error("main - net.Listen", err)
		return
	}

	mux := httpserver.New(grpcserver.Config.Port, logger)

	g, _ := errgroup.WithContext(context.Background())
	g.Go(func() (err error) {
		return grpcServer.Serve(listen)
	})
	g.Go(func() (err error) {
		serv := http.Server{
			Addr:         httpserver.Config.HostPort,
			ReadTimeout:  httpserver.Config.Duration,
			WriteTimeout: httpserver.Config.Duration,
			Handler:      mux,
		}
		return serv.ListenAndServe()
	})

	err = g.Wait()
	if err != nil {
		logger.Error("main - g.Wait", err)
	}
}
