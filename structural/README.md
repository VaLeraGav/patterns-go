# Структурные паттерны (Structural)

**Структурные паттерны** - это шаблоны проектирования, которые упрощают проектирование, определяя простой способ реализации взаимосвязей между объектами. Они помогают ответить на вопрос «Как создать программный компонент?»

Структурные паттерны делятся на два типа:

- **Паттерны уровня класса** - описывают взаимодействия между классами и их подклассами. Такие отношения выражаются путем наследования и реализации классов. Тут базовый класс определяет интерфейс, а подклассы - реализацию.
- **Паттерны уровня объекта** - описывают взаимодействия между объектами. Такие отношения выражаются связями - `ассоциацией`, `агрегацией` и `композицией`. Тут структуры строятся путем объединения объектов некоторых классов.

К паттернам уровня класса относится только «`Адаптер`». Смысл его работы в том, что если у вас есть класс и его интерфейс не совместим с библиотеками вашей системы, то что бы разрешить этот конфликт, мы не изменяем код этого класса, а пишем для него адаптер.

Все структурные паттерны отвечают за создание правильной структуры системы, в которой без труда смогут взаимодействовать между собой уже имеющиеся классы и объекты.

- [Адаптер (Adapter)](adapter)
- [Мост (Bridge)](bridge)
- [Компоновщик (Composite)](composite)
- [Декоратор (Decorator)](decorator)
- [Фасад (Facade)](facade)
- [Приспособленец (Flyweight)](flyweight)
- [Заместитель (Proxy)](proxy)