package managers

import "github.com/endevour-code-writer/taskManager/internal/taskManager/providers"

type TasksManager struct {
	taskProvider    *providers.TasksProvider
	commentsManager *CommentsManager
}
