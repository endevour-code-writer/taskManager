package taskManager

import (
	"fmt"
	"github.com/endevour-code-writer/taskManager/configs"
	"github.com/endevour-code-writer/taskManager/internal/db"
	"github.com/endevour-code-writer/taskManager/internal/router"
	"github.com/endevour-code-writer/taskManager/internal/taskManager/handlers"
	"github.com/endevour-code-writer/taskManager/internal/taskManager/managers"
	"github.com/endevour-code-writer/taskManager/internal/taskManager/server"
	"github.com/go-chi/chi"
	"github.com/golang-migrate/migrate/v4"
	migratePostgresSchema "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	_ "github.com/lib/pq"
	"log"
)

var app App

type App struct {
	config          configs.Configuration
	db              db.Database
	router          router.Router
	boardsHandler   *handlers.BoardsHandler
	columnsHandler  *handlers.ColumnsHandler
	tasksHandler    *handlers.TasksHandler
	commentsHandler *handlers.CommentsHandler
}

func Init() App {
	return app
}

func (app App) Run() {
	app.registerRoutesWithHandlers()
	address := ":" + fmt.Sprintf("%d", app.config.Server.Port)
	app.migrateDBSchema()
	httpServer := server.MakeHttpServer(app.router.RoutesMultiplexer)

	if err := httpServer.Start(address); err != nil {
		log.Fatalln(err)
	}
}

func (app App) CloseDB() {
	app.db.Close()
}

func init() {
	config := configs.GetInitConfig()
	dataBase := db.NewDataBase(config.Database).Open()
	appRouter := router.Init()
	responder := handlers.NewResponder()
	boardsManager, columnsManager, tasksManager, commentsManager := managers.ArrangeManagers(dataBase)
	boardsHandler, columnsHandler, tasksHandler, commentsHandler := handlers.ArrangeHandlers(boardsManager, columnsManager, tasksManager, commentsManager, responder)
	app = App{
		config,
		dataBase,
		appRouter,
		boardsHandler,
		columnsHandler,
		tasksHandler,
		commentsHandler}
}

func (app App) migrateDBSchema() {
	driver, err := migratePostgresSchema.WithInstance(app.db.DB, &migratePostgresSchema.Config{})

	if err != nil {
		//TODO: add logs
		log.Fatalf("DB migration: driver instance is failed: %v", err)
	}

	newMigrate, err := migrate.NewWithDatabaseInstance(app.db.GetMigrationPath(), app.db.GetDriver(), driver)

	if err != nil {
		//TODO : add logs
		log.Fatalf("DB migration: instance failed: %v", err)
	}

	err = newMigrate.Up()

	switch err {
	case migrate.ErrNoChange:
		//TODO: add logs
		fmt.Println("DB migration: database schema is already up to date")
	case nil:
		//TODO: add logs
		fmt.Println("DB migration: changes applied")
	default:
		log.Fatalf("DB migration: up failed: %v", err)
	}

}

func (app App) registerRoutesWithHandlers() {
	app.router.RoutesMultiplexer.Route("/board", func(r chi.Router) {
		r.Post("/new", app.boardsHandler.Create)
	})
}
