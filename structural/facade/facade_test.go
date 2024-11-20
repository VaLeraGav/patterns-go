package facade

import (
	"fmt"
	"testing"
)

func TestFacade(t *testing.T) {

	house := &House{}
	tree := &Tree{}
	man := NewMan(house, tree)

	result := man.Todo()

	fmt.Println(result)
	//	Build house
	// Tree grow
}
