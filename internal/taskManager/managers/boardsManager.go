package managers

import (
	"github.com/endevour-code-writer/taskManager/internal/taskManager/entities"
	"github.com/endevour-code-writer/taskManager/internal/taskManager/providers"
)

type BoardsManager struct {
	boardProvider  *providers.BoardsProvider
	columnsManager *ColumnsManager
}

func (boardsManager BoardsManager) Create(board *entities.Board) (*entities.Board, error) {
	tx, err := boardsManager.boardProvider.DataBase.DB.Begin()

	if err != nil {
		return nil, err
	}

	defer func() {
		_ = tx.Rollback()
	}()

	board, err = boardsManager.boardProvider.Save(board, tx)

	if err != nil {
		return nil, err
	}

	column := boardsManager.columnsManager.MakeDefaultColumn()

	_, err = boardsManager.columnsManager.columnsProvider.Save(board, column, tx)

	if err != nil {
		return nil, err
	}

	board.Columns = []*entities.Column{column}
	if err = tx.Commit(); err != nil {
		return nil, err
	}

	return board, nil
}
