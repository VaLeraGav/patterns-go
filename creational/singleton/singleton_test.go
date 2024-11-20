package singleton

import (
	"fmt"
	"testing"
)

func TestSingletonLight(t *testing.T) {

	instance1 := GetInstanceLight()
	instance2 := GetInstanceLight()

	if instance1 != instance2 {
		t.Error("Objects are not equal!\n")
	}
}

func TestSingletonLightTest(t *testing.T) {

	for i := 0; i < 30; i++ {
		go GetInstanceLightTest()
	}

	// Scanln is similar to Scan, but stops scanning at a newline and
	// after the final item there must be a newline or EOF.
	fmt.Scanln()
}

func TestSingletonMutex(t *testing.T) {

	for i := 0; i < 30; i++ {
		go GetInstanceMutex()
	}
	fmt.Scanln()
}
