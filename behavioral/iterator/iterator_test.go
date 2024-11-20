package iterator

import (
	"fmt"
	"testing"
)

func TestIterator(t *testing.T) {
	collection := NewCollection()
	collection.Add("Элемент 1")
	collection.Add("Элемент 2")
	collection.Add("Элемент 3")

	iterator := collection.CreateIterator()

	for iterator.HasNext() {
		item := iterator.Next()
		fmt.Println(item)
	}
}
