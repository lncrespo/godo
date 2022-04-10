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
		projectTodos, err := project.GetTodos(false)

		if err != nil {
			log.Println("Error during cleanup: " + err.Error())

			continue
		}

		inactiveTodos, err := project.GetTodos(true)

		if err != nil {
			log.Println("Error during cleanup: " + err.Error())

			continue
		}

		projectTodos = append(projectTodos, inactiveTodos...)

		if len(projectTodos) == 0 {
			err = project.Remove()

			if err != nil {
				log.Println("Error during cleanup: " + err.Error())
			}
		}
	}
}
