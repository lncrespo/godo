package dbal

import (
	"database/sql"
	"errors"
	"time"
)

type Todo struct {
	Id          int64
	Title       string
	Description string
	Priority    int16
	CreatedAt   time.Time
	Project     Project
}

func GetTodoByTitle(title string, projectName string) (Todo, error) {
	todo := Todo{}

	if db == nil {
		return todo, errors.New("Database connection is not established.")
	}

	query := "SELECT `id`, `title`, `description`, `priority`, `created_at` FROM `todo` WHERE `title` = ?"

	project, err := GetProjectByName(projectName)

	if err != nil && projectName == "" {
		query += " AND `project_id` IS NULL"

		statement, err := db.Prepare(query)

		if err != nil {
			return todo, err
		}

		err = statement.QueryRow(title).Scan(&todo.Id, &todo.Title, &todo.Description, &todo.Priority, &todo.CreatedAt)

		return todo, err
	}

	query += " AND `project_id` = ?"

	statement, err := db.Prepare(query)

	if err != nil {
		return todo, err
	}

	todo.Project = project
	err = statement.QueryRow(title, project.Id).Scan(&todo.Id, &todo.Title, &todo.Description, &todo.Priority, &todo.CreatedAt)

	return todo, err
}

func AddTodo(todo Todo) (int64, error) {
	if db == nil {
		return -1, errors.New("Database connection is not established.")
	}

	_, err := GetTodoByTitle(todo.Title, todo.Project.Name)

	if err == nil {
		return -1, errors.New("Todo already exists in the given project")
	}

	statement, err := db.Prepare(
		"INSERT INTO `todo` (`title`, `description`, `priority`, `project_id`) VALUES (?, ?, ?, ?)")

	if err != nil {
		return -1, err
	}

	var result sql.Result

	// TODO DRY - Is there another way to write this a bit cleaner?
	if todo.Project == (Project{}) {
		result, err = statement.Exec(todo.Title, todo.Description, todo.Priority, nil)
	} else {
		result, err = statement.Exec(todo.Title, todo.Description, todo.Priority, todo.Project.Id)
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
