package handlers

import (
	"github.com/endevour-code-writer/taskManager/internal/taskManager/managers"
)

func ArrangeHandlers(
	bordersManager *managers.BoardsManager,
	columnsManager *managers.ColumnsManager,
	tasksManager *managers.TasksManager,
	commentsManager *managers.CommentsManager,
	responder Responder) (*BoardsHandler, *ColumnsHandler, *TasksHandler, *CommentsHandler) {
	boardsHandler := NewBoardsHandler(bordersManager, responder)
	columnsHandler := NewColumnsHandler(columnsManager, responder)
	tasksHandler := NewTasksHandler(tasksManager, responder)
	commentsHandler := NewCommentsHandler(commentsManager, responder)

	return boardsHandler, columnsHandler, tasksHandler, commentsHandler
}
