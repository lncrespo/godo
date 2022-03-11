package godo

import (
	"bufio"
	"log"
	"os"
	"strconv"
	"strings"

	"github.com/lncrespo/godo/src/dbal"
	_ "github.com/mattn/go-sqlite3"
)

func add(title string, description string, priority int, projectName string) {
	if title == "" {
		title, description, priority, projectName = getTodoInteractively()
	}

	project, err := dbal.GetProjectByName(projectName)

	todo := dbal.Todo{
		Title: title,
		Description: description,
		Priority: int16(priority)}

	if err != nil && projectName != "" {
		os.Stdout.WriteString("Project does not exist. Creating project and adding todo.\n")
		project = dbal.Project{Name: projectName}

		// Declare projectId seperately, so `err` gets overwritten instead of redeclared.
		// That way `err` can be handled after the entire if-block without logging the same
		// `err` from above (`dbal.GetProjectByName(projectName)`)
		var projectId int64
		projectId, err = dbal.AddProject(project)

		if err != nil {
			log.Fatalln("Failed to create project" + err.Error())
		}

		project.Id = projectId

		todo.Project = project

		_, err = dbal.AddTodo(todo)
	} else if projectName == "" {
		_, err = dbal.AddTodo(todo)
	} else {
		todo.Project = project

		_, err = dbal.AddTodo(todo)
	}

	if err != nil {
		log.Fatalln(err)
	}
}

func getTodoInteractively() (string, string, int, string) {
	os.Stdout.WriteString("Please enter the title of your todo:\n")

	reader := bufio.NewReader(os.Stdin)
	title, err := reader.ReadString('\n')
	title = strings.TrimRight(title, "\n")

	if err != nil {
		log.Fatalln("Could not read from stdin: " + err.Error())
	}

	os.Stdout.WriteString("Please enter a description for your todo: (optional)\n")
	description := ""
	description = strings.TrimRight(description, "\n")

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
		log.Fatalln("Could not read from stdin.")
	}

	os.Stdout.WriteString("Please enter a priority for your todo: (0-9, defaults to 9)\n")
	priorityArg, err := reader.ReadString('\n')

	if err != nil {
		log.Fatalln("Could not read from stdin.")
	}

	priority, err := strconv.Atoi(strings.TrimRight(priorityArg, "\n"))

	if err != nil || priority < 0 || priority > 9 {
		os.Stdout.WriteString("Invalid priority entered, defaulting to 9\n" + err.Error())
		priority = 9
	}

	os.Stdout.WriteString("Please enter the project for your todo: (Leave empty for global)")
	project, err := reader.ReadString('\n')
	project = strings.TrimRight(project, "\n")

	if err != nil {
		log.Fatalln("Could not read from stdin: " + err.Error())
	}

	return title, description, priority, project
}
