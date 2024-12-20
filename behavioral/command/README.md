# Команда (Command)

Паттерн `Command` относится к поведенческим паттернам уровня объекта.

Паттерн `Command` позволяет представить запрос в виде объекта. Из этого следует, что команда - это объект. Такие запросы, например, можно ставить в очередь, отменять или возобновлять.

В этом паттерне мы оперируем следующими понятиями:

- `Command` - запрос в виде объекта на выполнение;
- `Receiver` - объект-получатель запроса, который будет обрабатывать нашу команду;
- `Invoker` - объект-инициатор запроса.
