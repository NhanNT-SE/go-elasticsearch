package rpc

import (
	"context"
	"fmt"
	"marketplace-backend/config"
	"marketplace-backend/internal/handler/vmp"
	"net/http"
	"sync"

	"github.com/gorilla/mux"
	"github.com/gorilla/rpc/v2"
)

type JSONRPCServer struct {
	cfg    config.ServerConfig
	server *http.Server

	mu sync.Mutex
}

func NewJSONRPCServer(cfg config.ServerConfig) *JSONRPCServer {
	return &JSONRPCServer{
		cfg: cfg,
	}
}

func (s *JSONRPCServer) Serve() error {
	// init handlers
	vmpHandler := vmp.NewRPCHandler(s.cfg)

	rpcServer := rpc.NewServer()
	rpcServer.RegisterCodec(newCustomCodec(), "application/json")
	if err := rpcServer.RegisterService(vmpHandler, "vmp"); err != nil {
		return fmt.Errorf("register service: %w", err)
	}

	r := mux.NewRouter()
	r.HandleFunc("/health", s.Health)
	r.Handle("/rpc", rpcServer)

	s.mu.Lock()
	s.server = &http.Server{
		Addr:    s.cfg.HTTPAddress,
		Handler: r,
	}
	s.mu.Unlock()

	return s.server.ListenAndServe()
}

func (s *JSONRPCServer) Shutdown(ctx context.Context) error {
	s.mu.Lock()
	srv := s.server
	s.mu.Unlock()

	if srv == nil {
		return nil
	}

	return srv.Shutdown(ctx)
}

func (s *JSONRPCServer) Health(w http.ResponseWriter, r *http.Request) {
	_, _ = fmt.Fprintln(w, "OK")
}
