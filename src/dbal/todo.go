package dbal

import (
	"database/sql"
	"errors"
	"time"
)

type Todo struct {
	Id          int64
	State       int16
	Title       string
	Description string
	Priority    int16
	CreatedAt   time.Time
	CompletedAt time.Time
	DueAt       time.Time
	Project     Project
}

func GetTodoById(id int64) (Todo, error) {
	todo := Todo{}

	if Db == nil {
		return todo, errors.New("Database connection is not established.")
	}

	query := "SELECT `id`, `title`, `description`, `priority`, `created_at`, `completed_at`, `due_at` FROM `todo` WHERE `id` = ?"

	statement, err := Db.Prepare(query)

	if err != nil {
		return todo, err
	}

	err = statement.QueryRow(id).
		Scan(&todo.Id,
			&todo.Title,
			&todo.Description,
			&todo.Priority,
			&todo.CreatedAt,
			&todo.CompletedAt,
			&todo.DueAt)

	return todo, err
}

func GetTodoByTitle(title string, projectName string, checkInactive bool) (Todo, error) {
	todo := Todo{}

	if Db == nil {
		return todo, errors.New("Database connection is not established.")
	}

	query := "SELECT `id`, `title`, `description`, `priority`, `created_at`, `completed_at`, `due_at` FROM `todo` WHERE `title` = ?"

	if !checkInactive {
		query += " AND `state` = 1"
	}

	project, err := GetProjectByName(projectName)

	if err != nil && projectName == "" {
		query += " AND `project_id` IS NULL"

		statement, err := Db.Prepare(query)

		if err != nil {
			return todo, err
		}

		err = statement.QueryRow(title).Scan(
			&todo.Id,
			&todo.Title,
			&todo.Description,
			&todo.Priority,
			&todo.CreatedAt,
			&todo.CompletedAt,
			&todo.DueAt)

		return todo, err
	}

	query += " AND `project_id` = ?"

	statement, err := Db.Prepare(query)

	if err != nil {
		return todo, err
	}

	todo.Project = project
	err = statement.QueryRow(title, project.Id).Scan(
		&todo.Id,
		&todo.Title,
		&todo.Description,
		&todo.Priority,
		&todo.CreatedAt,
		&todo.CompletedAt,
		&todo.DueAt)

	return todo, err
}

func (todo Todo) Add() (int64, error) {
	if Db == nil {
		return -1, errors.New("Database connection is not established.")
	}

	_, err := GetTodoByTitle(todo.Title, todo.Project.Name, false)

	if err == nil {
		return -1, errors.New("Todo already exists in the given project")
	}

	statement, err := Db.Prepare(
		"INSERT INTO `todo` (`title`, `description`, `priority`, `project_id`, `due_at`) VALUES (?, ?, ?, ?, ?)")

	if err != nil {
		return -1, err
	}

	var result sql.Result

	// TODO DRY - Is there another way to write this a bit cleaner?
	if todo.Project == (Project{}) {
		result, err = statement.Exec(todo.Title, todo.Description, todo.Priority, nil, todo.DueAt)
	} else {
		result, err = statement.Exec(todo.Title, todo.Description, todo.Priority, todo.Project.Id, todo.DueAt)
	}

	if err != nil {
		return -1, err
	}

	lastInsertedId, err := result.LastInsertId()

	if err != nil {
		return -1, err
	}

	return lastInsertedId, nil
}

func (project Project) GetTodos(onlyCheckInactive bool) ([]Todo, error) {
	todos := []Todo{}

	if Db == nil {
		return todos, errors.New("Database connection is not established.")
	}

	query := "SELECT `id`, `state`, `title`, `description`, `priority`, `created_at`, `completed_at`, `due_at` FROM `todo` WHERE `project_id` "

	if project == (Project{}) {
		query += "IS NULL"
	} else {
		query += "= ?"
	}

	query += " AND `state` = "

	if onlyCheckInactive {
		query += "0"
	} else {
		query += "1"
	}

	statement, err := Db.Prepare(query)

	if err != nil {
		return todos, err
	}

	var rows *sql.Rows

	if project == (Project{}) {
		rows, err = statement.Query()
	} else {
		rows, err = statement.Query(project.Id)
	}

	if err != nil {
		return todos, err
	}

	for rows.Next() {
		todo := Todo{}

		err := rows.Scan(
			&todo.Id,
			&todo.State,
			&todo.Title,
			&todo.Description,
			&todo.Priority,
			&todo.CreatedAt,
			&todo.CompletedAt,
			&todo.DueAt)

		if err != nil {
			continue
		}

		todos = append(todos, todo)
	}

	return todos, nil
}

func (todo Todo) ChangeState(state int16) error {
	if Db == nil {
		return errors.New("Database connection is not established.")
	}

	todo, err := GetTodoById(todo.Id)

	if err != nil && todo == (Todo{}) {
		return errors.New("Could not fetch todo from database.")
	}

	statement, err := Db.Prepare(
		"UPDATE `todo` SET `state` = ?, `completed_at` = ? WHERE `id` = ?")

	if err != nil {
		return err
	}

	completedAt := time.Time{}

	if state == 0 {
		completedAt = time.Now().UTC()
	}

	_, err = statement.Exec(state, completedAt, todo.Id)

	return err
}

func (todo Todo) Remove() error {
	if Db == nil {
		return errors.New("Database connection is not established.")
	}

	todo, err := GetTodoById(todo.Id)

	if err != nil && todo == (Todo{}) {
		return errors.New("Could not fetch todo from database.")
	}

	statement, err := Db.Prepare("DELETE FROM `todo` WHERE `id` = ?")

	if err != nil {
		return err
	}

	_, err = statement.Exec(todo.Id)

	return err
}

func TruncateTodos() error {
	if Db == nil {
		return errors.New("Database connection is not established.")
	}

	statement, err := Db.Prepare("DELETE FROM `todo`")

	if err != nil {
		return err
	}

	_, err = statement.Exec()

	return err
}
