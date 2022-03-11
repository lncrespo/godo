package godo

import (
	"fmt"
	"log"
	"os"
	"sort"
	"text/tabwriter"
	"time"

	"github.com/lncrespo/godo/src/dbal"
)

func list(projectName string) {
	project, err := dbal.GetProjectByName(projectName)

	if err != nil && project != (dbal.Project{}) {
		log.Fatalln(err)
	} else if err != nil && projectName != "" && project == (dbal.Project{}) {
		log.Fatalln("The given project does not exist")
	}

	todos, err := dbal.GetTodosByProject(project)

	sort.Slice(todos, func(i, j int) bool {
		return todos[i].Priority < todos[j].Priority
	})

	if err != nil {
		log.Fatalln(err)
	}

	writer := tabwriter.NewWriter(os.Stdout, 1, 1, 5, ' ', 0)

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

	writer.Flush()
}
