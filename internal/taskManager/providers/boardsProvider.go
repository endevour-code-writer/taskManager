package providers

import (
	"database/sql"
	"errors"
	"github.com/endevour-code-writer/taskManager/internal/db"
	"github.com/endevour-code-writer/taskManager/internal/taskManager/entities"
)

type BoardsProvider struct {
	DataBase db.Database
}

func (boardsProvider *BoardsProvider) Save(board *entities.Board, tx *sql.Tx) (*entities.Board, error) {
	if board == nil {
		//TODO: add logs
		return nil, errors.New("board: Nil pointer was passed for saving")
	}

	query := "INSERT INTO boards (name, description) VALUES ($1, $2) RETURNING id, name, description, created_at, updated_at;"
	stmt, err := tx.Prepare(query)

	if err != nil {
		return nil, err
	}

	err = stmt.QueryRow(board.Name, board.Description).Scan(
		&board.Id,
		&board.Name,
		&board.Description,
		&board.CreatedAt,
		&board.UpdatedAt)

	if err != nil {
		if err := stmt.Close(); err != nil {
			//TODO: add logs
		}
		return nil, err
	}

	return board, nil
}
