package singleton

import (
	"fmt"
	"sync"
)

// Singleton implementation.
type Singleton struct {
}

var (
	instance *Singleton
	once     sync.Once
)

func GetInstanceLight() *Singleton {
	once.Do(func() {
		instance = &Singleton{}
	})
	return instance
}

func GetInstanceLightTest() *Singleton {
	if instance == nil {
		once.Do(
			func() {
				fmt.Println("Creating single instance now.")
				instance = &Singleton{}
			})
	} else {
		fmt.Println("Single instance already created.")
	}
	return instance
}
