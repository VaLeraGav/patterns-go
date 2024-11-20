package iterator

// Коллекция
type Collection struct {
	items []interface{}
}

// Создание новой коллекции
func NewCollection() *Collection {
	return &Collection{}
}

// Добавление элемента в коллекцию
func (c *Collection) Add(item interface{}) {
	c.items = append(c.items, item)
}

// Создание итератора для коллекции
func (c *Collection) CreateIterator() Iterator {
	return &ConcreteIterator{collection: c, index: 0}
}
