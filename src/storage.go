package godo

import (
	"database/sql"
	"errors"
	"log"
	"os"

	_ "github.com/mattn/go-sqlite3"
)

type Todo struct {
	title       string
	description string
	priority    int16
}

var dbStoragePath string = ""

func initializeStorage() {
	homeDir, err := os.UserHomeDir()

	if err != nil {
		log.Fatalln("Unable to get the user home directory: " + err.Error())
	}

	dbStoragePath = homeDir + "/.local/share/.godo.sqlite"

	if _, err := os.Stat(homeDir + "/.local/share"); errors.Is(err, os.ErrNotExist) {
		os.Stderr.WriteString(
			"Could not find \".local/share\", falling back to home directory for database file.\n")

		dbStoragePath = homeDir + "/.godo.sqlite"
	}

	if _, err := os.Stat(dbStoragePath); errors.Is(err, os.ErrNotExist) {
		os.Stdout.WriteString(
			"Database file not found, creating database file and initializing tables.\n")

		initializeDatabase()
	}
}

func initializeDatabase() {
	db, err := sql.Open("sqlite3", dbStoragePath)

	if err != nil {
		log.Fatalln("Unable to open the database connection: " + err.Error())
	}

	tableQuery := "CREATE TABLE `todo` ("
	tableQuery += "`id` INTEGER PRIMARY KEY AUTOINCREMENT, "
	tableQuery += "`title` VARCHAR(255) NOT NULL, "
	tableQuery += "`description` TEXT NULL DEFAULT NULL,"
	tableQuery += "`priority` INTEGER NOT NULL DEFAULT 9, "
	tableQuery += "`created_at` DATETIME DEFAULT CURRENT_TIMESTAMP);"

	_, err = db.Exec(tableQuery)

	if err != nil {
		log.Fatalln("Unable to create table `todo`: " + err.Error())
	}

	tableQuery = "CREATE TABLE `project` ("
	tableQuery += "`id` INTEGER PRIMARY KEY AUTOINCREMENT, "
	tableQuery += "`name` VARCHAR(255) NOT NULL, "
	tableQuery += "`created_at` DATETIME DEFAULT CURRENT_TIMESTAMP);"

	_, err = db.Exec(tableQuery)

	if err != nil {
		log.Fatalln("Unable to create table `project`: " + err.Error())
	}

	tableQuery = "CREATE TABLE `todo_project` ("
	tableQuery += "`todo_id` INTEGER, "
	tableQuery += "`project_id` INTEGER, "
	tableQuery += " PRIMARY KEY(`todo_id`, `project_id`));"

	_, err = db.Exec(tableQuery)

	if err != nil {
		log.Fatalln("Unable to create table `todo_project`: " + err.Error())
	}
}
