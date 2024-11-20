# Приспособленец (Flyweight)

Паттерн `Flyweight` относится к структурным паттернам уровня объекта.

Паттерн `Flyweight` используется для эффективной поддержки большого числа мелких объектов, он позволяет повторно использовать мелкие объекты в различном контексте.

Требуется для реализации:

Класс `FlyweightFactory`, являющейся модифицированным паттерном фабрики, для создания приспособленцев

```go
type FlyweightFactory struct {
	pool map[string]Flyweighter
}

// GetFlyweight creates or returns a suitable Flyweighter by state.
func (f *FlyweightFactory) GetFlyweight(filename string) Flyweighter {
	if f.pool == nil {
		f.pool = make(map[string]Flyweighter)
	}
	if _, ok := f.pool[filename]; !ok {
		f.pool[filename] = &ConcreteFlyweight{filename: filename}
	}
	return f.pool[filename]
}
```

`Flyweight` описания общего интерфейса приспособленцев

```go
type Flyweighter interface {
	Draw(width, height int, opacity float64) string
}
```

Класс `ConcreteFlyweight` реализующий приспособленца, который будет замещать собой одинаковые мелкие объекты.

```go
type ConcreteFlyweight struct {
	filename string
}

func (f *ConcreteFlyweight) Draw(width, height int, opacity float64) string {
	return fmt.Sprintf("draw image: %s, width: %d, height: %d, opacity: %.2f", f.filename, width, height, opacity)
}
```

Суть в том, что мы можем запрашивать приспособленцев у `фабрики` по запросу, в свою очередь она будет отдавать те объекты, которые уже были созданы, или создавать новые. Это означает, что мы будем использовать уже созданные объекты, а не создавать ещё больше, если объекты под наши нужны уже имеются. Также стоит обратить внимание, что приспособленцы имеют внутреннее и внешние состояние. Фабрика находит приспособленцев по внутреннему состоянию, а внешнее состояние передается в его методы.
