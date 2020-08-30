package handlers

import (
	"github.com/endevour-code-writer/taskManager/internal/taskManager/managers"
)

type TasksHandler struct {
	manager   *managers.TasksManager
	responder Responder
}

//var task *entities.Task

func NewTasksHandler(tasksManager *managers.TasksManager, responder Responder) *TasksHandler {
	return &TasksHandler{tasksManager, responder}
}
