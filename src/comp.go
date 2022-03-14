package godo

import (
	"log"

	"github.com/lncrespo/godo/src/dbal"
)

func complete(compFlags completeCommandFlags) {
	err := dbal.ChangeTodoStateById(int64(compFlags.id), 0)

	if err != nil {
		log.Fatalln(err)
	}
}
