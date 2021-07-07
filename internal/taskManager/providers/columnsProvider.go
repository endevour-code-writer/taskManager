package providers

import (
	"database/sql"
	"errors"
	"github.com/endevour-code-writer/taskManager/internal/db"
	"github.com/endevour-code-writer/taskManager/internal/taskManager/entities"
)

type ColumnsProvider struct {
	DataBase db.Database
}

func (columnsProvider *ColumnsProvider) Save(board *entities.Board, column *entities.Column, tx *sql.Tx) (sql.Result, error) {
	if column == nil {
		//TODO: add logs
		return nil, errors.New("column provider: Nil pointer was passed for saving")
	}

	query := "INSERT INTO columns (name, position, board_id) VALUES ($1, $2, $3)"
	stmt, err := tx.Prepare(query)

	if err != nil {
		return nil, err
	}

	result, err := stmt.Exec(column.Name, column.Position, board.Id)

	if err != nil {
		if err := stmt.Close(); err != nil {
			//TODO: add logs
		}
		return nil, err
	}

	return result, nil
}
