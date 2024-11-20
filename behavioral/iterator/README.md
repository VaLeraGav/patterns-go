# Итератор (Iterator)

Паттерн `Iterator` относится к поведенческим паттернам уровня объекта.

Паттерн `Iterator` предоставляет механизм обхода коллекций объектов не раскрывая их внутреннего представления.

Зачастую этот паттерн используется вместо массива объектов, чтобы не только предоставить доступ к элементам, но и наделить некоторой логикой.

`Iterator` представляет собой общий интерфейс, позволяющий реализовать произвольную логику итераций.

Требуется для реализации:

- Интерфейс `Iterator` описывающий набор методов для доступа к коллекции

```go
// Интерфейс итератора
type Iterator interface {
	HasNext() bool
	Next() interface{}
}
```

- Класс Concrete`Iterator`, реализующий интерфейс `Iterator`. Следит за позицией текущего элемента при переборе коллекции (Aggregate)

```go
type ConcreteIterator struct {
	collection *Collection
	index      int
}

func (i *ConcreteIterator) HasNext() bool {
	return i.index < len(i.collection.items)
}

func (i *ConcreteIterator) Next() interface{} {
	if i.HasNext() {
		item := i.collection.items[i.index]
		i.index++
		return item
	}
	return nil
}
```

- Класс `NewCollection`, хранящий в себе элементы коллекции.

```go
type Collection struct {
	items []interface{}
}

func NewCollection() *Collection {
	return &Collection{}
}

func (c *Collection) Add(item interface{}) {
	c.items = append(c.items, item)
}

func (c *Collection) CreateIterator() Iterator {
	return &ConcreteIterator{collection: c, index: 0}
}
```
