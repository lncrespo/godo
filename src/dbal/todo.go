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
	Project     Project
}

func GetTodoById(id int64) (Todo, error) {
	todo := Todo{}

	if db == nil {
		return todo, errors.New("Database connection is not established.")
	}

	query := "SELECT `id`, `title`, `description`, `priority`, `created_at` FROM `todo` WHERE `id` = ?"

	statement, err := db.Prepare(query)

	if err != nil {
		return todo, err
	}

	err = statement.QueryRow(id).
		Scan(&todo.Id, &todo.Title, &todo.Description, &todo.Priority, &todo.CreatedAt)

	return todo, err
}

func GetTodoByTitle(title string, projectName string, checkInactive bool) (Todo, error) {
	todo := Todo{}

	if db == nil {
		return todo, errors.New("Database connection is not established.")
	}

	query := "SELECT `id`, `title`, `description`, `priority`, `created_at` FROM `todo` WHERE `title` = ?"

	if !checkInactive {
		query += " AND `state` = 1"
	}

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

	_, err := GetTodoByTitle(todo.Title, todo.Project.Name, false)

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

func GetTodosByProject(project Project, onlyCheckInactive bool) ([]Todo, error) {
	todos := []Todo{}

	if db == nil {
		return todos, errors.New("Database connection is not established.")
	}

	query := "SELECT `id`, `title`, `description`, `priority`, `created_at` FROM `todo` WHERE `project_id` "

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

	statement, err := db.Prepare(query)

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

		err := rows.Scan(&todo.Id, &todo.Title, &todo.Description, &todo.Priority, &todo.CreatedAt)

		if err != nil {
			continue
		}

		todos = append(todos, todo)
	}

	return todos, nil
}

func ChangeTodoStateById(id int64, state int16) error {
	todo, err := GetTodoById(id)

	if err != nil && todo == (Todo{}) {
		return errors.New("Could not fetch todo from database.")
	}

	statement, err := db.Prepare("UPDATE `todo` SET `state` = ? WHERE `id` = ?")

	if err != nil {
		return err
	}

	_, err = statement.Exec(state, id)

	return err
}

func RemoveTodoById(id int64) error {
	todo, err := GetTodoById(id)

	if err != nil && todo == (Todo{}) {
		return errors.New("Could not fetch todo from database.")
	}

	statement, err := db.Prepare("DELETE FROM `todo` WHERE `id` = ?")

	if err != nil {
		return err
	}

	_, err = statement.Exec(id)

	return err
}
