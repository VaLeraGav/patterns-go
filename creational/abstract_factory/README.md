# Абстрактная фабрика (Abstract Factory)

Паттерн `Abstract Factory` относится к порождающим паттернам уровня объекта.

Паттерн `Abstract Factory` предоставляет общий интерфейс для создания семейства взаимосвязанных объектов. Это позволяет отделить функциональность системы от внутренней реализации каждого класса, а обращение к этим классам становится возможным через абстрактные интерфейсы

В общем виде абстрактная фабрика выглядит следующим образом. Для каждого из семейств объектов, создается конкретная фабрика (наследник абстрактной), посредством которой создаются продукты этого семейства.

## Концептуальный пример

Представим, что вам нужно купить спортивную форму, состоящую из двух разных вещей: пара обуви и футболка. Вы хотите приобрести полный набор от одного бренда, чтобы вещи сочитались между собой.

- `abstract_sports_factory.go`: Интерфейс абстрактной фабрики

```go
type AbstractSportsFactory interface {
	MakeShoe() AbstractShoe
	MakeShirt() AbstractShirt
}

func CreateAdidasFactory() AbstractSportsFactory {
	return &Adidas{}
}
func CreateNikeFactory() AbstractSportsFactory {
	return &Nike{}
}
```

- `adidas.go` и `nike.go`: Конкретная фабрика

```go
type Nike struct {
}

func (n *Nike) MakeShoe() AbstractShoe {
	return &NikeShoe{
        logo: "nike",
        size: 14,
	}
}

func (n *Nike) MakeShirt() AbstractShirt {
	return &NikeShirt{
        size: 14,
	}
}

```

- `abstract_shoe.go` и `abstract_shirt.go`: Абстрактный продукт — это интерфейс или абстрактный класс, который определяет общие характеристики и поведение для группы связанных объектов. Он не содержит конкретных реализаций, а только описывает, какие методы должны быть реализованы в конкретных продуктах.

```go
type AbstractShirt interface {
	setSize(size int)
	GetSize() int
}

type AbstractShoe interface {
	setLogo(logo string)
	setSize(size int)
	GetLogo() string
	GetSize() int
}
```

- `adidasCertain.go` и `nikeCertain.go`: Конкретный продукт — это класс, который реализует интерфейс абстрактного продукта. Он содержит **конкретные реализации** методов, определенных в **абстрактном продукте**, и представляет собой реальный объект, который можно создать и использовать в программе.
