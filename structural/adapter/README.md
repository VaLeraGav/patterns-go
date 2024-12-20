# Адаптер (Adapter)

Паттерн `Adapter` относится к структурным паттернам уровня класса.

Часто в новом проекте разработчики хотят повторно использовать уже существующий код. Например, имеющиеся классы могут обладать нужной функциональностью и иметь при этом несовместимые интерфейсы. В таких случаях следует использовать паттерн `Adapter`.

Смысл работы этого паттерна в том, что если у вас есть класс и его интерфейс не совместим с кодом вашей системы, то что бы разрешить этот конфликт, мы не изменяем код этого класса, а пишем для него адаптер. Другими словами `Adapter` адаптирует существующий код к требуемому интерфейсу (является переходником).
