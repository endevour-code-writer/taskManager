package managers

import (
	"github.com/endevour-code-writer/taskManager/internal/taskManager/entities"
	"github.com/endevour-code-writer/taskManager/internal/taskManager/providers"
	"time"
)

const (
	DefaultColumnName     = "Default"
	DefaultColumnPosition = 100
)

type ColumnsManager struct {
	columnsProvider *providers.ColumnsProvider
	tasksManager    *TasksManager
}

func (columnManager *ColumnsManager) MakeDefaultColumn() *entities.Column {
	defaultColumn := &entities.Column{
		Name:      DefaultColumnName,
		Position:  DefaultColumnPosition,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now()}

	return defaultColumn
}

//func (columnManager *ColumnsManager) Save(board *entities.Board, column *entities.Column) (*entities.Column, error) {
//
//}
