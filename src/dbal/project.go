package dbal

import (
	"database/sql"
	"errors"
	"time"
)

type Project struct {
	Id        int64
	Name      string
	CreatedAt time.Time
}

func GetProjectById(id int64) (Project, error) {
	project := Project{}

	if Db == nil {
		return project, errors.New("Database connection is not established.")
	}

	statement, err := Db.Prepare(
		"SELECT `id`, `name`, `created_at` FROM `project` WHERE `id` = ?")

	if err != nil {
		return project, err
	}

	err = statement.QueryRow(id).Scan(&project.Id, &project.Name, &project.CreatedAt)

	return project, err
}

func GetProjectByName(name string) (Project, error) {
	project := Project{}

	if Db == nil {
		return project, errors.New("Database connection is not established.")
	}

	statement, err := Db.Prepare(
		"SELECT `id`, `name`, `created_at` FROM `project` WHERE `name` = ?")

	if err != nil {
		return project, err
	}

	err = statement.QueryRow(name).Scan(&project.Id, &project.Name, &project.CreatedAt)

	return project, err
}

func GetProjects() ([]Project, error) {
	projects := []Project{}

	if Db == nil {
		return projects, errors.New("Database connection is not established.")
	}

	query := "SELECT `id`, `name`, `created_at` FROM `project`"

	statement, err := Db.Prepare(query)

	if err != nil {
		return projects, err
	}

	var rows *sql.Rows

	rows, err = statement.Query()

	if err != nil {
		return projects, err
	}

	for rows.Next() {
		project := Project{}

		err := rows.Scan(&project.Id, &project.Name, &project.CreatedAt)

		if err != nil {
			continue
		}

		projects = append(projects, project)
	}

	return projects, nil
}

func (project Project) Add() (int64, error) {
	if Db == nil {
		return -1, errors.New("Database connection is not established.")
	}

	statement, err := Db.Prepare(
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

func (project Project) Remove() error {
	if Db == nil {
		return errors.New("Database connection is not established.")
	}

	project, err := GetProjectById(project.Id)

	if err != nil {
		return errors.New("Could not fetch project from database.")
	}

	statement, err := Db.Prepare("DELETE FROM `project` WHERE `id` = ?")

	if err != nil {
		return err
	}

	_, err = statement.Exec(project.Id)

	return err
}
