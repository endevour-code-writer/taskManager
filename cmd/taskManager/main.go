package main

import "github.com/endevour-code-writer/taskManager/internal/taskManager"

func main() {
	defer app.CloseDB()
	app := taskManager.Init()
	app.Run()
}
