package chain_of_responsibility

// Конкретный обработчик для уровня важности "низкий"
type LowPriorityHandler struct {
	BaseHandler
}

func (h *LowPriorityHandler) Handle(request string) string {
	if request == "low" {
		return "LowPriorityHandler: Handling low priority request."
	}
	return h.BaseHandler.Handle(request)
}
