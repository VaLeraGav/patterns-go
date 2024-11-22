package functional_options

import "time"

type server struct {
	Port       int
	Timeout    time.Duration
	EnableLogs bool
}

type serverOption func(*server)

func WithPort(port int) serverOption {
	return func(s *server) {
		s.Port = port
	}
}

func WithTimeout(timeout time.Duration) serverOption {
	return func(s *server) {
		s.Timeout = timeout
	}
}

func WithLogs(enabled bool) serverOption {
	return func(s *server) {
		s.EnableLogs = enabled
	}
}

func NewServer(opts ...serverOption) *server {
	server := &server{
		Port:       8080,
		Timeout:    60,
		EnableLogs: false,
	}

	for _, opt := range opts {
		opt(server)
	}

	return server
}
