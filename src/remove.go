package godo

import "github.com/lncrespo/godo/src/dbal"

func remove(removeFlags removeCommandFlags) {
	dbal.RemoveTodoById(int64(removeFlags.id))
}
