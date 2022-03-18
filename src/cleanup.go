package godo

import (
	"log"

	"github.com/lncrespo/godo/src/dbal"
)

func cleanup() {
	removeEmptyProjects()
}

func removeEmptyProjects() {
	projects, err := dbal.GetProjects()

	if err != nil {
		log.Println("Error during cleanup: " + err.Error())
	}

	for _, project := range projects {
		projectTodos, err := dbal.GetTodosByProject(project, false)

		if err != nil {
			log.Println("Error during cleanup: " + err.Error())

			continue
		}

		inactiveTodos, err := dbal.GetTodosByProject(project, true)

		if err != nil {
			log.Println("Error during cleanup: " + err.Error())

			continue
		}

		projectTodos = append(projectTodos, inactiveTodos...)

		if len(projectTodos) == 0 {
			err = dbal.RemoveProject(project)

			if err != nil {
				log.Println("Error during cleanup: " + err.Error())
			}
		}
	}
}
