# Фасад (Facade)

Паттерн `Facade` относится к структурным паттернам уровня объекта.

Паттерн `Facade` предоставляет высокоуровневый унифицированный интерфейс в виде набора имен методов к набору взаимосвязанных классов или объектов некоторой подсистемы, что облегчает ее использование.

Разбиение сложной системы на подсистемы позволяет упростить процесс разработки, а также помогает максимально снизить зависимости одной подсистемы от другой. Однако использовать такие подсистемы становиться довольно сложно. Один из способов решения этой проблемы является паттерн Facade. Наша задача, сделать простой, единый интерфейс, через который можно было бы взаимодействовать с подсистемами.

В качестве примера можно привести интерфейс автомобиля. Современные автомобили имеют унифицированный интерфейс для водителя, под которым скрывается сложная подсистема. Благодаря применению навороченной электроники, делающей большую часть работы за водителя, тот может с лёгкостью управлять автомобилем, не задумываясь, как там все работает.

Почему это можно назвать Фасадом:

- **Скрытие сложности**: Класс `Man` предоставляет простой интерфейс (метод `Todo`()), который скрывает сложность взаимодействия с другими классами (`House` и `Tree`). Пользователь `Man` не должен знать о внутреннем устройстве `House` и `Tree`, чтобы выполнить действия.
- **Упрощение интерфейса**: Вместо того чтобы взаимодействовать с `House` и `Tree` напрямую, пользователь работает только с `Man`, что упрощает использование системы.
- **Агрегация**: `Man` агрегирует объекты House и Tree, что позволяет управлять ими через единый интерфейс.

Примечание: Хороший фасад не содержит созданий экземпляров классов (new) внутри. Если внутри фасада создаются объекты для реализации каждого метода, это не Фасад, это Строитель или [Абстрактная|Статическая|Простая] Фабрика [или Фабричный Метод].
