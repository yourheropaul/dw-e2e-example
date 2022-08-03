package http

import (
	"context"
	"log"
	"net"
	"net/http"
	"sync"
)

type Handler = http.Handler

type Server struct {
	server *http.Server
	wg     *sync.WaitGroup
}

func NewServer(addr string, handler http.Handler) *Server {
	return &Server{
		server: &http.Server{
			Addr:    addr,
			Handler: handler,
		},
		wg: &sync.WaitGroup{},
	}
}

func (h *Server) Start() error {
	h.wg.Add(1)

	listener, err := net.Listen("tcp", h.server.Addr)
	if err != nil {
		return err
	}

	go func() {
		defer h.wg.Done()
		if err := h.server.Serve(listener); err != http.ErrServerClosed {
			log.Fatalf("Serve: %v", err)
		}
	}()

	return nil
}

func (h *Server) Stop() error {
	if h.server == nil {
		return nil
	}

	if err := h.server.Shutdown(context.Background()); err != nil {
		return err
	}

	// Wait until ListenAndServe completes
	h.wg.Wait()

	return nil
}
