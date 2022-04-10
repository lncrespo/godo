package godo

import (
	"errors"
	"strconv"
)

type addCommandFlags struct {
	title       *string
	description *string
	priority    *int
	project     string
	dueAt       *string
}

type listCommandFlags struct {
	showProjects *bool
	showAll      *bool
	project      string
}

type completeCommandFlags struct {
	id int64
}

type removeCommandFlags struct {
	id int64
}

type overviewCommandFlags struct {
	showAll *bool
}

type infoCommandFlags struct {
	id int64
}

func getIdFromArgument(argument string) (int64, error) {
	if argument == "" {
		return -1, errors.New("Missing argument")
	}

	id, err := strconv.Atoi(argument)

	if err != nil {
		return -1, err
	}

	return int64(id), nil
}
