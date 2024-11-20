package chain_of_responsibility

// Конкретный обработчик для уровня важности "высокий"
type HighPriorityHandler struct {
	BaseHandler
}

func (h *HighPriorityHandler) Handle(request string) string {
	if request == "high" {
		return "HighPriorityHandler: Handling high priority request."
	}
	return h.BaseHandler.Handle(request)
}
