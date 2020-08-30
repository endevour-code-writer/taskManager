package db

import (
	"database/sql"
	"fmt"
	"github.com/endevour-code-writer/taskManager/configs"
	_ "github.com/lib/pq"
	"log"
)

type Database struct {
	DB *sql.DB
	dbConfig configs.DatabaseConfig
}

func NewDataBase(dbConfig configs.DatabaseConfig) Database {
	return Database{
		dbConfig: dbConfig,
	}
}

func (db Database) Open() Database {
	dsn :=
		fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
			db.dbConfig.Host,
			db.dbConfig.Port,
			db.dbConfig.User,
			db.dbConfig.Password,
			db.dbConfig.Name)

	openedDB, err := sql.Open(db.dbConfig.Driver, dsn)

	if err != nil {
		log.Fatal(err)
	}

	db.DB = openedDB
	err = db.DB.Ping()

	if err != nil {
		fmt.Println("Unpingable")
	} else {
		fmt.Println("Pingable")
	}

	return db
}

func (db Database) Close() {
	db.DB.Close()
}

func (db Database) GetDriver() string  {
	return db.dbConfig.Driver
}

func (db Database) GetMigrationPath() string  {
	return db.dbConfig.MigrationsPath
}

