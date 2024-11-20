package abstract_factory

import (
	"fmt"
	"testing"
)

func TestAbstractFactory(t *testing.T) {
	adidasFactory := CreateAdidasFactory()
	nikeFactory := CreateNikeFactory()

	nikeShoe := nikeFactory.MakeShoe()
	nikeShirt := nikeFactory.MakeShirt()

	adidasShoe := adidasFactory.MakeShoe()
	adidasShirt := adidasFactory.MakeShirt()

	printShoeDetails(nikeShoe)
	printShirtDetails(nikeShirt)
	// Logo: nike
	// Size: 14
	// Size: 14

	printShoeDetails(adidasShoe)
	printShirtDetails(adidasShirt)
	// Logo: adidas
	// Size: 16
	// Size: 16
}

func printShoeDetails(s AbstractShoe) {
	fmt.Printf("Logo: %s \n", s.GetLogo())
	fmt.Printf("Size: %d \n", s.GetSize())

}

func printShirtDetails(s AbstractShirt) {
	fmt.Printf("Size: %d \n", s.GetSize())
}
