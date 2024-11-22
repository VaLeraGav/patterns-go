package functional_options

import "time"

type serverOptionsTwo struct {
	Port       int
	Timeout    time.Duration
	EnableLogs bool
}

func NewServerOptionsTwo() *serverOptionsTwo {
	return &serverOptionsTwo{
		Port:       8080,
		Timeout:    60 * time.Second,
		EnableLogs: false,
	}
}

type Server struct {
	options *serverOptionsTwo
}

// NewServer создает новый сервер с заданными параметрами
func NewServerTwo(options *serverOptionsTwo) *Server {
	return &Server{options: options}
}
