package handlers

import (
	"encoding/json"
	"github.com/endevour-code-writer/taskManager/internal/taskManager/entities"
	"github.com/endevour-code-writer/taskManager/internal/taskManager/managers"
	"io/ioutil"
	"net/http"
)

type BoardsHandler struct {
	manager   *managers.BoardsManager
	responder Responder
}

var board *entities.Board

func NewBoardsHandler(boardsManager *managers.BoardsManager, responder Responder) *BoardsHandler {
	return &BoardsHandler{boardsManager, responder}
}

func (boardsHandler *BoardsHandler) Ping(w http.ResponseWriter, r *http.Request) {
	boardsHandler.responder.responseJSON(w, http.StatusOK, map[string]string{"message": "Pong"})
}

func (boardsHandler *BoardsHandler) Create(w http.ResponseWriter, r *http.Request) {
	request, err := ioutil.ReadAll(r.Body)

	if err != nil {
		//TODO: add log
		badResponse := "Request data wasn't read"
		boardsHandler.responder.responseError(w, http.StatusBadRequest, badResponse)
	}

	if err = json.Unmarshal(request, &board); err != nil {
		//TODO: add log)
		boardsHandler.responder.responseError(w, http.StatusBadRequest, err)
		return
	}

	board, err = boardsHandler.manager.Create(board)

	if err != nil {
		//TODO: add logs
		boardsHandler.responder.responseError(w, http.StatusBadRequest, err)
		return
	}

	boardsHandler.responder.responseJSON(w, http.StatusOK, board)
}
