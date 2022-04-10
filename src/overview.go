package godo

import (
	"log"
	"os"
	"strings"

	"github.com/lncrespo/godo/src/dbal"
)

func overview(overviewFlags overviewCommandFlags) {
	projects, err := dbal.GetProjects()

	if err != nil {
		log.Fatalln(err)
	}

	emulatedListFlags := listCommandFlags{}
	projectBorderPart := strings.Repeat("─", len("Global Todos")+2)
	emulatedListFlags.project = ""
	emulatedListFlags.showAll = overviewFlags.showAll

	os.Stdout.WriteString(
		projectBorderPart + "┐\n Global Todos │\n" + projectBorderPart + "┘\n")

	list(emulatedListFlags)

	for _, project := range projects {
		emulatedListFlags.project = project.Name

		projectTodos, err := project.GetTodos(false)

		if err != nil {
			continue
		}

		if *overviewFlags.showAll {
			inactiveTodos, err := project.GetTodos(true)

			if err == nil {
				projectTodos = append(projectTodos, inactiveTodos...)
			}
		}

		if len(projectTodos) == 0 {
			continue;
		}

		projectBorderPart = strings.Repeat("─", len(project.Name)+2)

		os.Stdout.WriteString(
			projectBorderPart + "┐\n " + project.Name + " │\n" + projectBorderPart + "┘\n")

		list(emulatedListFlags)
	}
}
