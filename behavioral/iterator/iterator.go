package iterator

// Интерфейс итератора
type Iterator interface {
	HasNext() bool
	Next() interface{}
}

// Итератор для коллекции
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
