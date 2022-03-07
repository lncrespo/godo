package godo

import (
	"bufio"
	"database/sql"
	"log"
	"os"
	"strconv"

	_ "github.com/mattn/go-sqlite3"
)

func add(title string, description string, priority int, project string) {
	if title == "" {
		title, description, priority, project = getInteractiveValues()
	}

	todo := Todo{
		title,
		description,
		int16(priority)}

	addTodo(todo)
}

func addTodo(todo Todo) {
	if dbStoragePath == "" {
		log.Fatalln("Database storage path is not yet initialized.")
	}

	db, err := sql.Open("sqlite3", dbStoragePath)

	if err != nil {
		log.Fatalln("Unable to open the database connection: " + err.Error())
	}

	addStatement, err := db.Prepare(
		"INSERT INTO `todo`(`title`, `description`, `priority`) VALUES (?, ?, ?)")

	if err != nil {
		log.Fatalln("Unable to insert into database: " + err.Error())
	}

	_, err = addStatement.Exec(todo.title, todo.description, todo.priority)

	if err != nil {
		log.Fatalln("Unable to insert into database: " + err.Error())
	}
}

func getInteractiveValues() (string, string, int, string) {
	os.Stdout.WriteString("Please enter the title of your todo:\n")

	reader := bufio.NewReader(os.Stdin)
	title, err := reader.ReadString('\n')

	if err != nil {
		log.Fatal("Could not read from stdin: " + err.Error())
	}

	os.Stdout.WriteString("Please enter a description for your todo: (optional)\n")
	description := ""

	for {
		line, err := reader.ReadString('\n')

		if err != nil {
			log.Fatalln("Could not read from stdin: " + err.Error())
		}

		description += line

		if len(line) == 1 {
			break
		}
	}

	if err != nil {
		log.Fatal("Could not read from stdin.")
	}

	os.Stdout.WriteString("Please enter a priority for your todo: (0-9, defaults to 9)\n")
	priorityArg, err := reader.ReadString('\n')

	if err != nil {
		log.Fatal("Could not read from stdin.")
	}

	priority, err := strconv.Atoi(priorityArg)

	if err != nil || priority < 0 || priority > 9 {
		os.Stdout.WriteString("Invalid priority entered, defaulting to 9\n" + err.Error())
		priority = 9
	}

	os.Stdout.WriteString("Please enter the project for your todo: (Leave empty for global)")
	project, err := reader.ReadString('\n')

	if err != nil {
		log.Fatal("Could not read from stdin: " + err.Error())
	}

	return title, description, priority, project
}
