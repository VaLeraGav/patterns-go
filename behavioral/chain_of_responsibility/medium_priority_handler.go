package chain_of_responsibility

// Конкретный обработчик для уровня важности "средний"
type MediumPriorityHandler struct {
	BaseHandler
}

func (h *MediumPriorityHandler) Handle(request string) string {
	if request == "medium" {
		return "MediumPriorityHandler: Handling medium priority request."
	}
	return h.BaseHandler.Handle(request)
}
