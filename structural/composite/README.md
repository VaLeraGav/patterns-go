# Компоновщик (Composite)

Паттерн `Composite` относится к структурным паттернам уровня объекта.

Паттерн `Composite` группирует схожие объекты в древовидные структуры.

Для построения дерева будут использоваться массивы, представляющие ветви дерева.

- `component.go`: Содержит интерфейс `Component` - предоставляет интерфейс, как для ветвей, так и для листьев дерева, который определяет метод Operation. Все элементы меню (листья и составные) будут реализовывать этот интерфейс.

```go
type Component interface {
	Operation() string
}
```

- `leaf.go`: Реализует простой элемент меню (`Leaf`) - реализующий интерфейс `Component` и являющийся листом дерева. Он имеет имя и реализует метод Operation, который возвращает строку с названием элемента.

```go
type Leaf struct {
	name string
}

func NewLeaf(name string) *Leaf {
	return &Leaf{name: name}
}

func (l *Leaf) Operation() string {
	return fmt.Sprintf("Leaf: %s", l.name)
}
```

- `composite.go`: Реализует составной элемент меню (`Composite`), реализующий интерфейс `Component` и являющийся ветвью дерева. Он может содержать другие компоненты и реализует метод `Operation`, который рекурсивно вызывает `Operation` для всех дочерних элементов.

```go
type Composite struct {
	name     string
	children []Component
}

func NewComposite(name string) *Composite {
	return &Composite{name: name}
}

func (c *Composite) Add(component Component) {
	c.children = append(c.children, component)
}

func (c *Composite) Operation() string {
	result := fmt.Sprintf("Composite: %s\n", c.name)
	for _, child := range c.children {
		result += "  " + child.Operation() + "\n"
	}
	return result
}
```

Обратите внимание, что лист дерева является классом листовых узлов и **не может иметь потомков** (из листа не может вырасти ветвь или другой лист).

Ветви дерева задают поведение объектов, входящих в структуру дерева, у которых есть потомки, а также сами хранит в себе компоненты дерева. Другим словами ветви могут содержать другие ветви и листья.

Основным назначением паттерна, является обеспечение единого интерфейса как к составному (ветви) так и конечному (листу) объекту, что бы клиент не задумывался над тем, с каким объектом он работает.
