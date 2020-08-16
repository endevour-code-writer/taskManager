package main

import (
	"github.com/endevour-code-writer/taskManager/internal/taskManager"
)

func main () {
	app := taskManager.Init()
	app.Run()
	defer app.CloseDB()
}