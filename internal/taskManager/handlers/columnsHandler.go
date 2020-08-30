package handlers

import (
	"github.com/endevour-code-writer/taskManager/internal/taskManager/managers"
)

type ColumnsHandler struct {
	manager   *managers.ColumnsManager
	responder Responder
}

//var column *entities.Column

func NewColumnsHandler(columnsManager *managers.ColumnsManager, responder Responder) *ColumnsHandler {
	return &ColumnsHandler{columnsManager, responder}
}
