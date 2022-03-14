package dbal

import (
	"database/sql"
	"errors"
	"log"
	"os"

	_ "github.com/mattn/go-sqlite3"
)

var dbStoragePath string = ""
var db *sql.DB = nil

func init() {
	initializeStorage()

	if dbStoragePath == "" {
		log.Fatalln("Database storage path is not yet initialized.")
	}

	connection, err := sql.Open("sqlite3", dbStoragePath)

	if err != nil {
		log.Fatalln("Unable to open the database connection: " + err.Error())
	}

	db = connection
}

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
	tableQuery += "`state` INTEGER NOT NULL DEFAULT 1, "
	tableQuery += "`title` VARCHAR(255) NOT NULL, "
	tableQuery += "`description` TEXT NULL DEFAULT NULL,"
	tableQuery += "`priority` INTEGER NOT NULL DEFAULT 9, "
	tableQuery += "`created_at` DATETIME DEFAULT CURRENT_TIMESTAMP, "
	tableQuery += "`project_id` INTEGER NULL DEFAULT NULL, "
	tableQuery += "FOREIGN KEY (`project_id`) REFERENCES `project`(`id`), "
	tableQuery += "UNIQUE(`state`, `title`, `project_id`));"

	_, err = db.Exec(tableQuery)

	if err != nil {
		log.Fatalln("Unable to create table `todo`: " + err.Error())
	}

	tableQuery = "CREATE TABLE `project` ("
	tableQuery += "`id` INTEGER PRIMARY KEY AUTOINCREMENT, "
	tableQuery += "`name` VARCHAR(255) NOT NULL, "
	tableQuery += "`created_at` DATETIME DEFAULT CURRENT_TIMESTAMP, "
	tableQuery += "UNIQUE(`name`));"

	_, err = db.Exec(tableQuery)

	if err != nil {
		log.Fatalln("Unable to create table `project`: " + err.Error())
	}
}
