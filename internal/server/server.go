package server

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gin-gonic/gin"
)

type Server struct {
	httpServer *http.Server
}

func NewServer(router *gin.Engine, addr string) *Server {
	return &Server{
		httpServer: &http.Server{
			Addr:    addr,
			Handler: router,
		},
	}
}

func (s *Server) Start() {
	go func() {
		log.Println("Starting server on", s.httpServer.Addr)
		if err := s.httpServer.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("Could not listen on %s: %v\n", s.httpServer.Addr, err)
		}
	}()
}

func (s *Server) GracefulShutdown(timeout time.Duration) {
	// Create a channel to listen for OS signals.
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, syscall.SIGINT, syscall.SIGTERM)

	<-stop // Block until a signal is received.

	log.Println("Shutting down server...")

	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()

	if err := s.httpServer.Shutdown(ctx); err != nil {
		log.Fatalf("Server forced to shutdown: %v", err)
	}

	log.Println("Server exited gracefully.")
}
