# Строитель (Builder)

Паттерн проектирования "`Строитель`" (Builder) позволяет создавать сложные объекты пошагово. Он отделяет конструкцию объекта от его представления, что позволяет создавать различные представления одного и того же типа объекта.

Есть 2 типа паттерн Строитель

1. Последовательно указывать методы построения объекта

```go
builder := NewConcretePizzaBuilder()
pizza := builder.SetSize("Large").
    SetDough("Thin").
    AddTopping("Olives").
    Build()

```

2. Метод в котором указываются одинаковые этапы построения

- Класс Director, который будет распоряжаться строителем и отдавать ему команды в нужном порядке, а строитель будет их выполнять

```go
type IBuilder interface {
	setWindowType()
	setNumFloor()
	getHouse() House
}

type Director struct {
	builder IBuilder
}

// Частный случай

func (d *Director) buildHouse() House {
	d.builder.setWindowType()
	d.builder.setNumFloor()
	return d.builder.getHouse()
}
```

- Базовый абстрактный класс `Builder`, который описывает интерфейс строителя, те команды, которые он обязан выполнять (в моем случае выполняемые команды определяются в `Director`)

```go
// ---- NormalBuilder ----
type NormalBuilder struct {
	windowType string
	floor      int
}

// ...

func (b *NormalBuilder) getHouse() House {
	return House{
		windowType: b.windowType,
		floor:      b.floor,
	}
}
```

```go
director := newDirector(normalBuilder)
normalHouse := director.buildHouse()
```
