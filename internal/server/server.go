package server

import (
	"context"
	"fmt"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"time"
)

// Server ...
type Server struct {
	httpServer *http.Server
	wait       time.Duration
}

// NewServer new http server instance
func NewServer(port string, router *mux.Router) *Server {
	corsOpts := handlers.AllowedOrigins([]string{"*"})

	return &Server{
		httpServer: &http.Server{
			Addr: fmt.Sprintf(":%s", port),
			// good practice to set timeouts to avoid slowloris attacks.
			WriteTimeout: time.Second * 15,
			ReadTimeout:  time.Second * 15,
			Handler:      handlers.CORS(corsOpts)(router),
		},
		wait: 10 * time.Second,
	}
}

// Serve starts http server
func (s *Server) Serve() error {
	err := s.httpServer.ListenAndServe()
	if err != nil {
		return err
	}
	return nil
}

// Shutdown closes http server
func (s *Server) Shutdown() {
	// Create a deadline to wait for.
	ctx, cancel := context.WithTimeout(context.Background(), s.wait)
	defer cancel()
	// Doesn't block if no connections, but will otherwise wait
	// until the timeout deadline.
	err := s.httpServer.Shutdown(ctx)
	if err != nil {
		log.Printf("could not shutdown server: %v\n", err)
	}
}
