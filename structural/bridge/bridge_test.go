package bridge

import (
	"fmt"
	"testing"
)

func TestBridge(t *testing.T) {
	suzukiCar := NewCar(&EngineSuzuki{})
	hondaCar := NewCar(&EngineHonda{})
	ladaCar := NewCar(&EngineLada{})

	fmt.Println("Suzuki Car Sound:", suzukiCar.Rase())
	fmt.Println("Honda Car Sound:", hondaCar.Rase())
	fmt.Println("Lada Car Sound:", ladaCar.Rase())

	// 	Suzuki Car Sound: SssuuuuZzzuuuuKkiiiii
	// Honda Car Sound: HhoooNnnnnnnnnDddaaaaaaa
	// Lada Car Sound: PhhhhPhhhhPhPhPhPhPh
}
