package dbal

import (
	"database/sql"
	"errors"
	"log"
	"os"

	_ "github.com/mattn/go-sqlite3"
)

var dbStoragePath string = ""
var Db *sql.DB = nil

func init() {
	initializeStorage()

	if dbStoragePath == "" {
		log.Fatalln("Database storage path is not yet initialized.")
	}

	connection, err := sql.Open("sqlite3", dbStoragePath)

	if err != nil {
		log.Fatalln("Unable to open the database connection: " + err.Error())
	}

	Db = connection
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
	var err error
	Db, err = sql.Open("sqlite3", dbStoragePath)

	if err != nil {
		log.Fatalln("Unable to open the database connection: " + err.Error())
	}

	createTodoTable();
	createProjectTable();
	createMigrationTable();
}

func createTodoTable() {
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

	_, err := Db.Exec(tableQuery)

	if err != nil {
		log.Fatalln("Unable to create table `todo`: " + err.Error())
	}
}

func createProjectTable() {
	tableQuery := "CREATE TABLE `project` ("
	tableQuery += "`id` INTEGER PRIMARY KEY AUTOINCREMENT, "
	tableQuery += "`name` VARCHAR(255) NOT NULL, "
	tableQuery += "`created_at` DATETIME DEFAULT CURRENT_TIMESTAMP, "
	tableQuery += "UNIQUE(`name`));"

	_, err := Db.Exec(tableQuery)

	if err != nil {
		log.Fatalln("Unable to create table `project`: " + err.Error())
	}
}

func createMigrationTable() {
	tableQuery := "CREATE TABLE `migration_version` ("
	tableQuery += "`version` INTEGER NOT NULL DEFAULT 0);"

	_, err := Db.Exec(tableQuery)

	if err != nil {
		log.Fatalln("Unable to create table `migration_version`: " + err.Error())
	}

	tableQuery = "INSERT INTO `migration_version` ("
	tableQuery += "`version`) VALUES (0);"

	_, err = Db.Exec(tableQuery)

	if err != nil {
		log.Fatalln("Unable to insert into table `migration_version`: " + err.Error())
	}
}
