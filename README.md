## Паттерны

[Порождающие (Creational)](creational) - сфокусированы на процессе инстанцирования объектов или групп связанных объектов. Сосредоточены на процессе внедрения объектов или групп объектов

- [Абстрактная фабрика (Abstract Factory)](creational/abstract_factory)
- [Строитель (Builder)](creational/builder)
- [Фабричный метод (Factory Method)](creational/factory_method)
- [Прототип (Prototype)](creational/prototype)
- [Одиночка (Singleton)](creational/singleton)

[Структурные (Structural)](structural) - это шаблоны проектирования, которые упрощают проектирование, определяя простой способ реализации взаимосвязей между объектами. Они помогают ответить на вопрос «Как создать программный компонент?»

- [Адаптер (Adapter)](structural/adapter)
- [Мост (Bridge)](structural/bridge)
- [Компоновщик (Composite)](structural/composite)
- [Декоратор (Decorator)](structural/decorator)
- [Фасад (Facade)](structural/facade)
- [Приспособленец (Flyweight)](structural/flyweight)
- [Заместитель (Proxy)](structural/proxy)

[Поведенческие (Behavioral)](behavioral) - это связано с распределением обязанностей между объектами. Что отличает их от структурных паттернов, так это то, что они не только определяют структуру, но и обрисовывают в общих чертах паттерны для передачи/общения между ними. Или, другими словами, они помогают ответить на вопрос «Как запустить поведение в программном компоненте?»

- [Цепочка ответственности (Chain Of Responsibility)](behavioral/chain_of_responsibility)
- [Команда (Command)](behavioral/command)
- [Итератор (Iterator)](behavioral/iterator)
- [Посредник (Mediator)](behavioral/mediator)
- [Хранитель (Memento)](behavioral/memento)
- [Наблюдатель (Observer)](behavioral/observer)
- [Хранилище (Repository)](behavioral/repository)
- [Состояние (State)](behavioral/state)
- [Стратегия (Strategy)](behavioral/strategy)
- [Шаблонный метод (Template Method)](behavioral/template_method)
- [Посетитель (Visitor)](behavioral/visitor)

> [!] В описании паттернов применяются общие понятия, такие как Класс, Объект, Абстрактный класс. Применимо к языку Go, это Пользовательский Тип, Значение этого Типа и Интерфейс. Также в языке Go за место общепринятого наследования используется агрегирование и встраивание.

`Ассоциация` - отношение, когда объекты двух классов могут ссылаться один на другой. Например, свойство класса содержит экземпляр другого класса.

`Агрегация` – частная форма ассоциации. Агрегация применяется, когда один объект должен быть контейнером для других объектов и время существования этих объектов никак не зависит от времени существования объекта контейнера. Вообщем, если контейнер будет уничтожен, то входящие в него объекты не пострадают. Например, мы создали объект, а потом передали его в объект контейнер, каким-либо образом, можно в метод объекта контейнера передать или присвоить сразу свойству контейнера извне. Значит, при удалении контейнера мы ни как не затронем наш созданный объект, который может взаимодействовать и с другими контейнерами.

`Композиция` – Тоже самое, что и агрегация, но составные объекты не могут существовать отдельно от объекта контейнера и если контейнер будет уничтожен, то всё его содержимое будет уничтожено тоже. Например, мы создали объект в методе объекта контейнера и присвоили его свойству объекта контейнера. Из вне, о нашем созданном объекте никто не знает, значит, при удалении контейнера, созданный объект будет удален так же, т.к. на него нет ссылки извне.

[Паттерны Параллельного программирования](parallel) -

## Как различать ?

Различать паттерны проектирования можно по нескольким критериям. Вот основные аспекты, которые помогут вам идентифицировать и классифицировать паттерны:

1. Назначение паттерна, повторим еще раз

- Создающие паттерны: Отвечают за создание объектов. Примеры: Фабрика (`Factory`), Абстрактная фабрика (`Abstract Factory`), Синглтон (`Singleton`), Прототип (`Prototype`).
- Структурные паттерны: Определяют, как классы и объекты композируются для формирования больших структур. Примеры: Фасад (`Facade`), Декоратор (`Decorator`), Компоновщик (`Composite`), Адаптер (`Adapter`).
- Поведенческие паттерны: Определяют, как объекты взаимодействуют и распределяют ответственность. Примеры: Наблюдатель (`Observer`), Стратегия (`Strategy`), Команда (`Command`), Итератор (`Iterator`).

2. Способ реализации

- Инкапсуляция: Некоторые паттерны инкапсулируют сложную логику (например, `Facade`).
- Инъекция зависимостей: Паттерны, такие как `Factory` или `Builder`, могут использовать инъекцию зависимостей для создания объектов с необходимыми параметрами.

3. Уровень абстракции

- Высокий уровень абстракции: Паттерны, которые предоставляют общий подход к решению проблем (например, `Abstract Factory`).
- Низкий уровень абстракции: Паттерны, которые решают более конкретные задачи (например, `Decorator`).

Вот пример часто встречаемого кода, и объяснение какой из паттернов реализуется

```go
type server struct {
	logger       *logrus.Logger
	store        store.Store
	sessionStore sessions.Store
}

func newServer(store store.Store, sessionStore sessions.Store) *server {
	s := &server{
		logger:       logrus.New(),
		store:        store,
		sessionStore: sessionStore,
	}

	s.configureRouting()

	return s
}
```

- Паттерн `Фабрика` (Factory) - Функция `newServer` отвечает за создание экземпляра структуры `server`, инициализируя её поля и настраивая необходимые зависимости. Это соответствует паттерну Фабрика, который используется для создания объектов с определённой конфигурацией.
- Паттерн `Строитель` (Builder) - Хотя в вашем примере не видно явного поэтапного создания объекта, функция newServer может рассматриваться как часть паттерна Строитель, если бы у вас было больше параметров или более сложная логика настройки. Она инкапсулирует процесс создания и настройки объекта.
- Паттерн `Инъекции Зависимостей` (Dependency Injection) - Передача store и sessionStore в качестве параметров функции newServer является примером инъекции зависимостей. Это позволяет легко заменять зависимости (например, для тестирования) и способствует слабой связанности между компонентами.

Да, реализация вашего кода действительно может напоминать паттерн `Фасад`, но давайте разберёмся, почему это не совсем так.

Почему это может напоминать `Фасад`:

- **Скрытие сложности**: Функция `newServer` инициализирует объект server и настраивает его зависимости, что может восприниматься как скрытие сложности создания и настройки объекта.
- **Упрощение интерфейса**: Если server инкапсулирует логику работы с store и sessionStore, то это может восприниматься как упрощение взаимодействия с этими компонентами.

Однако, ключевые отличия:

1. Назначение:

- Фасад: Предоставляет простой интерфейс для работы с сложной системой или набором интерфейсов. Он скрывает внутренние детали реализации и предоставляет упрощённый доступ к функциональности.
- Ваш код: Функция newServer просто создаёт и настраивает объект. Она не предоставляет упрощённый интерфейс для работы с несколькими подсистемами, как это делает Фасад.

2. Структура
   В паттерне Фасад обычно есть один класс (фасад), который объединяет несколько подсистем и предоставляет методы для их взаимодействия. В вашем коде server не является фасадом для других компонентов, а скорее представляет собой конкретный объект с определёнными зависимостями.

Еще один пример часто используемого кода

```go
type Config struct {
	BindAddr    string `toml:"bind_addr"`
	LogLevel    string `toml:"log_level"`
}


func NewConfig() *Config {
	return &Config{
		BindAddr: ":8080",
		LogLevel: "debug",
	}
}
```

Данный код соответствует паттерну `Фабрика`, так как он предоставляет метод для создания и настройки нового объекта с определёнными значениями по умолчанию. Она инкапсулирует логику создания объекта, что позволяет избежать дублирования кода при создании конфигурации в разных частях программы.

---

## Использованные источники

1. Большая часть примеров и теории было взято из https://github.com/AlexanderGrom/go-patterns, выражаю благодарность
2. [5 паттернов параллельного программирования в GO, которые сделают ваш следующий проект лучше](https://habr.com/ru/companies/otus/articles/722880/)
3. https://refactoringguru.cn/ru/design-patterns/go