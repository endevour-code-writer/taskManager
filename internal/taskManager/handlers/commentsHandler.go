package handlers

import (
	"github.com/endevour-code-writer/taskManager/internal/taskManager/managers"
)

type CommentsHandler struct {
	manager   *managers.CommentsManager
	responder Responder
}

//var comment *entities.Comment

func NewCommentsHandler(commentsManager *managers.CommentsManager, responder Responder) *CommentsHandler {
	return &CommentsHandler{commentsManager, responder}
}
