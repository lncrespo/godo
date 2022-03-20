package godo

import (
	"fmt"
	"log"
	"time"

	"github.com/lncrespo/godo/src/dbal"
)

func info(infoFlags infoCommandFlags) {
	todo, err := dbal.GetTodoById(infoFlags.id)

	if err != nil {
		log.Fatalln(err)
	}

	fmt.Printf("Showing info for todo with ID %d\n\n", todo.Id)
	fmt.Printf("Title:\n%s\n\n", todo.Title)
	fmt.Printf("Description:\n%s\n\n", todo.Description)
	fmt.Printf("Priority:\n%d\n\n", todo.Priority)
	fmt.Printf("Created at:\n%s\n", todo.CreatedAt.Local().Format(time.RFC1123))
}
