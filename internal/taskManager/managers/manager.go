package managers

import (
	"github.com/endevour-code-writer/taskManager/internal/db"
	"github.com/endevour-code-writer/taskManager/internal/taskManager/providers"
)

func ArrangeManagers(database db.Database) (*BoardsManager, *ColumnsManager, *TasksManager, *CommentsManager) {
	boardsProvider := &providers.BoardsProvider{DataBase: database}
	columnsProvider := &providers.ColumnsProvider{DataBase: database}
	tasksProvider := &providers.TasksProvider{DataBase: database}
	commentsProvider := &providers.CommentsProvider{DataBase: database}
	commentsManager := &CommentsManager{commentsProvider}
	tasksManager := &TasksManager{tasksProvider, commentsManager}
	columnsManager := &ColumnsManager{columnsProvider, tasksManager}
	boardsManager := &BoardsManager{boardsProvider, columnsManager}

	return boardsManager, columnsManager, tasksManager, commentsManager
}
