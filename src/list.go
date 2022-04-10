package godo

import (
	"errors"
	"fmt"
	"log"
	"os"
	"sort"
	"text/tabwriter"
	"time"

	"github.com/lncrespo/godo/src/dbal"
)

var writer *tabwriter.Writer

func init() {
	writer = tabwriter.NewWriter(os.Stdout, 1, 1, 5, ' ', 0)
}

func list(listFlags listCommandFlags) {
	project, err := dbal.GetProjectByName(listFlags.project)

	if listFlags.showProjects != nil && *listFlags.showProjects {
		projects, err := dbal.GetProjects()

		if err != nil {
			log.Fatalln(err)
		}

		printProjects(projects)

		return
	}

	if err != nil && project != (dbal.Project{}) {
		log.Fatalln(err)
	} else if err != nil && listFlags.project != "" && project == (dbal.Project{}) {
		log.Fatalln("The given project does not exist")
	}

	todos, err := project.GetTodos(false)

	if err != nil {
		log.Fatalln(err)
	}

	if listFlags.showAll != nil && *listFlags.showAll {
		inactiveTodos, err := project.GetTodos(true)

		if err == nil {
			todos = append(todos, inactiveTodos...)
		}
	}

	printTodos(todos, listFlags.showAll)
}

func printTodos(todos []dbal.Todo, showAll *bool) error {
	if writer == nil {
		return errors.New("Writer is not initialized")
	}

	sort.Slice(todos, func(i, j int) bool {
		return todos[i].Priority < todos[j].Priority
	})

	checkbox := ""

	if showAll != nil && *showAll {
		fmt.Fprintln(writer, "ID\tCompleted\tTitle\tPriority\tCreated at\tCompleted at")
		fmt.Fprintln(writer, "--\t---------\t-----\t--------\t----------\t------------")

		for _, todo := range todos {
			checkbox = "[ ]"

			if todo.State == 0 {
				checkbox = "[✓]"
			}

			completedAtTime := ""

			if todo.State == 0 {
				completedAtTime = todo.CompletedAt.Local().Format(time.RFC1123)
			}

			fmt.Fprintf(
				writer,
				"%d\t%s\t%s\t%d\t%s\t%s\n",
				todo.Id,
				checkbox,
				todo.Title,
				todo.Priority,
				todo.CreatedAt.Local().Format(time.RFC1123),
				completedAtTime)
		}
	} else {
		fmt.Fprintln(writer, "ID\tTitle\tPriority\tCreated at")
		fmt.Fprintln(writer, "--\t-----\t--------\t----------")

		for _, todo := range todos {
			fmt.Fprintf(
				writer,
				"%d\t%s\t%d\t%s\n",
				todo.Id,
				todo.Title,
				todo.Priority,
				todo.CreatedAt.Local().Format(time.RFC1123))
		}
	}

	writer.Flush()

	return nil
}

func printProjects(projects []dbal.Project) error {
	if writer == nil {
		return errors.New("Writer is not initialized")
	}

	fmt.Fprintln(writer, "ID\tName\tCreated at")
	fmt.Fprintln(writer, "--\t-----\t----------")

	for _, project := range projects {
		fmt.Fprintf(
			writer,
			"%d\t%s\t%s\n",
			project.Id,
			project.Name,
			project.CreatedAt.Local().Format(time.RFC1123))
	}

	writer.Flush()

	return nil
}
