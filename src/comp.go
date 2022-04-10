package godo

import (
	"log"

	"github.com/lncrespo/godo/src/dbal"
)

func complete(compFlags completeCommandFlags) {
	todo, err := dbal.GetTodoById(int64(compFlags.id))

	if err != nil {
		log.Fatalln(err)
	}

	err = todo.ChangeState(0)

	if err != nil {
		log.Fatalln(err)
	}
}
