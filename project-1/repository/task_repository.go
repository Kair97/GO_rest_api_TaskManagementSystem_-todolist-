package repository

import (
	"database/sql"
	"errors"
	"strconv"
	"try/database"
	"try/models"
)

func GetTasks() ([]models.Task, error) {

	query := `
		Select id, title, description, completed, priority 
		from tasks
	`
	rows, err := database.DB.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var tasks []models.Task

	for rows.Next() {
		var task models.Task

		err := rows.Scan(
			&task.ID,
			&task.Title,
			&task.Description,
			&task.Completed,
			&task.Priority,
		)
		if err != nil {
			return nil, err
		}
		tasks = append(tasks, task)
	}
	return tasks, nil

}

func CreateTask(title string, description *string, priority int) (models.Task, error) {

	query := `
		insert into tasks (title, description, completed, priority)
		values ($1, $2, false, $3)
		returning id
	`

	var id int
	err := database.DB.QueryRow(
		query,
		title,
		description,
		priority,
	).Scan(&id)

	if err != nil {
		return models.Task{}, err
	}

	task := models.Task{
		ID:          id,
		Title:       title,
		Description: description,
		Completed:   false,
		Priority:    priority,
	}

	return task, nil

}

func GetTaskByID(id int) (models.Task, error) {

	query := `
	select id, title, description, completed, priority 
	from tasks
	where id = $1
	`

	var task models.Task

	err := database.DB.QueryRow(query, id).Scan(
		&task.ID,
		&task.Title,
		&task.Description,
		&task.Completed,
		&task.Priority,
	)
	if err == sql.ErrNoRows {
		return models.Task{}, nil
	}

	if err != nil {
		return models.Task{}, err
	}

	return task, err

}

func UpdateTask(id int, upd models.UpdateTask) (int64, error) {

	query := `
		update tasks
		set title = $1, description = $2, completed = $3, priority = $4
		where id = $5  
	`

	result, err := database.DB.Exec(query, upd.Title, *upd.Description, upd.Completed, upd.Priority, id)
	if err != nil {
		return 0, err
	}

	affected, err := result.RowsAffected()
	if err != nil {
		return 0, err
	}

	return affected, nil

}

func DeleteTask(id int) (int64, error) {

	query := `
	delete from tasks where id = $1
	`

	result, err := database.DB.Exec(query, id)
	if err != nil {
		return 0, err
	}

	affected, err := result.RowsAffected()
	if err != nil {
		return 0, err
	}

	return affected, nil

}

var ErrNothingToUpdate = errors.New("nothing to update")

var Notfound = errors.New("not found")

func PatchTask(id int, pat models.PatchTask) error {

	if pat.Title == nil && pat.Description == nil && pat.Completed == nil && pat.Priority == nil {
		return ErrNothingToUpdate
	}

	query := "Update tasks set "
	args := []any{}
	argId := 1

	if pat.Title != nil {
		query += "title = $" + strconv.Itoa(argId)
		argId++
		args = append(args, *pat.Title)
	}

	if pat.Description != nil {
		if argId > 1 {
			query += ", "
		}
		query += "description = $" + strconv.Itoa(argId)
		argId++
		args = append(args, *pat.Description)
	}

	if pat.Completed != nil {
		if argId > 1 {
			query += ", "
		}
		query += "completed = $" + strconv.Itoa(argId)
		argId++
		args = append(args, *pat.Completed)
	}

	if pat.Priority != nil {
		if argId > 1 {
			query += ", "
		}
		query += "priority = $" + strconv.Itoa(argId)
		argId++
		args = append(args, *pat.Priority)
	}

	query += " where id = $" + strconv.Itoa(argId)
	args = append(args, id)

	result, err := database.DB.Exec(query, args...)
	if err != nil {
		return err
	}

	affected, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if affected == 0 {
		return Notfound
	}

	return nil

}
