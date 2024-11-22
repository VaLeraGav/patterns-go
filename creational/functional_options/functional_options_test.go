package functional_options

import (
	"fmt"
	"testing"
	"time"
)

func TestFunctionalOptions(t *testing.T) {
	cfg := NewServer(WithPort(9090), WithLogs(true))
	fmt.Println(cfg)
	// &{9090 60ns true}
}

func TestFunctionalOptionsTwo(t *testing.T) {
	serverOptions := NewServerOptionsTwo()
	serverOptions.Port = 9090
	serverOptions.Timeout = 30 * time.Second // Установка таймаута
	serverOptions.EnableLogs = true

	cfg := NewServerTwo(serverOptions)
	fmt.Println(cfg)
}
