package godo

import (
	"log"

	"github.com/lncrespo/godo/src/dbal"
)

func remove(removeFlags removeCommandFlags) {
	todo, err := dbal.GetTodoById(removeFlags.id)

	if err != nil {
		log.Fatalln("Could not fetch the todo from the database")
	}

	err = todo.Remove()

	if err != nil {
		log.Fatalln(err)
	}
}
