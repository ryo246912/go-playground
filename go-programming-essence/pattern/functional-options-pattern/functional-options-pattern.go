package main

import (
	"log"
	"os"
	"time"
)

type Server struct {
	host    string
	port    int
	timeout time.Duration
	logger  *log.Logger
}

func New(host string, port int, opts ...Option) *Server {
	s := &Server{
		host: host,
		port: port,
	}

	for _, opt := range opts {
		opt(s)
	}

	return s
}

func (s *Server) Start() error {
	s.logger.Printf("Starting server on %s:%d with timeout %s", s.host, s.port, s.timeout)
	// Simulate server start logic
	time.Sleep(1 * time.Second) // Simulating some startup delay
	s.logger.Println("Server started successfully")
	return nil
}

func WithTimeout(timeout time.Duration) func(*Server) {
	return func(s *Server) {
		s.timeout = timeout
	}
}
func WithLogger(logger *log.Logger) func(*Server) {
	return func(s *Server) {
		s.logger = logger
	}
}

type Option func(*Server)

func main() {
	srv := New("localhost", 8080,
		WithTimeout(10*time.Second),
		WithLogger(log.New(os.Stdout, "INFO: ", log.LstdFlags)),
	)
	if err := srv.Start(); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}

}
