package proxy

import (
	"fmt"
	"testing"
)

func TestProxy(t *testing.T) {
	proxy := &Proxy{}
	result := proxy.Send()

	fmt.Printf("result: %s \n", result)
	// <strong>I’ll be back!</strong>
}
