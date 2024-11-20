package chain_of_responsibility

// Интерфейс обработчика
type Handler interface {
	SetNext(handler Handler) Handler
	Handle(request string) string
}

// Базовый класс для обработчиков
type BaseHandler struct {
	next Handler
}

// Устанавливает следующий обработчик в цепочке
func (h *BaseHandler) SetNext(handler Handler) Handler {
	h.next = handler
	return handler
}

// Метод для передачи запроса следующему обработчику
func (h *BaseHandler) Handle(request string) string {
	if h.next != nil {
		return h.next.Handle(request)
	}
	return ""
}
