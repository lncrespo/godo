package dbal

import (
	"errors"
	"time"
)

type Project struct {
	Id        int64
	Name      string
	CreatedAt time.Time
}

func GetProjectByName(name string) (Project, error) {
	project := Project{}

	if db == nil {
		return project, errors.New("Database connection is not established.")
	}

	statement, err := db.Prepare(
		"SELECT `id`, `name`, `created_at` FROM `project` WHERE `name` = ?")

	if err != nil {
		return project, err
	}

	err = statement.QueryRow(name).Scan(&project.Id, &project.Name, &project.CreatedAt)

	return project, err
}

func AddProject(project Project) (int64, error) {
	if db == nil {
		return -1, errors.New("Database connection is not established.")
	}

	statement, err := db.Prepare(
		"INSERT INTO `project` (`name`) VALUES (?)")

	if err != nil {
		return -1, err
	}

	result, err := statement.Exec(project.Name)

	if err != nil {
		return -1, err
	}

	lastInsertedId, err := result.LastInsertId()

	if err != nil {
		return -1, err
	}

	return lastInsertedId, nil
}
