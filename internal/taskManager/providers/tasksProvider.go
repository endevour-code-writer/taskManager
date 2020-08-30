package providers

import (
	_ "database/sql"
	_ "errors"
	"github.com/endevour-code-writer/taskManager/internal/db"
	_ "github.com/endevour-code-writer/taskManager/internal/taskManager/entities"
)

type TasksProvider struct {
	DataBase db.Database
}
