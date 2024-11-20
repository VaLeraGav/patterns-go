package composite

import (
	"fmt"
	"testing"
)

func TestComposite(t *testing.T) {
	// Создаем листья
	leaf1 := NewLeaf("Menu Item 1")
	leaf2 := NewLeaf("Menu Item 2")
	leaf3 := NewLeaf("Submenu Item 1")

	// Создаем составной элемент
	submenu := NewComposite("Submenu")
	submenu.Add(leaf3)

	// Создаем корневой элемент
	root := NewComposite("Main Menu")
	root.Add(leaf1)
	root.Add(leaf2)
	root.Add(submenu)

	// Выполняем операцию
	fmt.Println(root.Operation())

	// Composite: Main Menu
	// 		Leaf: Menu Item 1
	// 		Leaf: Menu Item 2
	// 		Composite: Submenu
	// 		Leaf: Submenu Item 1
}
