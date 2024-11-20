package house

import (
	"fmt"
	"testing"
)

func TestHouseBuilder(t *testing.T) {
	normalBuilder := NewNormalBuilder()

	director := newDirector(normalBuilder)
	normalHouse := director.buildHouse()

	fmt.Printf("Normal House Door Type: %s\n", normalHouse.doorType)
	fmt.Printf("Normal House Window Type: %s\n", normalHouse.windowType)
	fmt.Printf("Normal House Num Floor: %d\n", normalHouse.floor)

	//	Normal House Door Type: Wooden Door
	// Normal House Window Type: Wooden Window
	// Normal House Num Floor: 2

	// Можно указать свой Builder
	// director.setBuilder(iglooBuilder)
}
