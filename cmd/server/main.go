package main

import (
	"context"
	"go-elasticsearch/config"
	"go-elasticsearch/internal/rpc"
	"go-elasticsearch/pkg/logger"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

var (
	log = logger.New()
)

type JSONServer struct{}

func main() {
	var cfg config.ServerConfig
	config.MustLoad(&cfg)
	log.Debug().Interface("config", cfg).Msg("ServerConfig loaded")

	srv := rpc.NewJSONRPCServer(cfg)
	go func() {
		if err := srv.Serve(); err != nil && err != http.ErrServerClosed {
			log.Fatal().Err(err).Msg("serve rpc failed")
		}
	}()

	ctx, _ := signal.NotifyContext(context.Background(), os.Kill, os.Interrupt, syscall.SIGTERM)
	<-ctx.Done()

	cancelCtx, cc := context.WithTimeout(context.Background(), 15*time.Second)
	defer cc()
	if err := srv.Shutdown(cancelCtx); err != nil {
		log.Err(err).Msg("shutdown rpc failed")
	}
}
