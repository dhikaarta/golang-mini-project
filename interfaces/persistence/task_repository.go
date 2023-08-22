package persistence

import (
	"database/sql"
	"fmt"
	"task-management-app/internal/entity"
)

type taskRepository struct {
	db *sql.DB
}

func NewTaskRepository(db *sql.DB) *taskRepository {
	return &taskRepository{
		db: db,
	}
}

func (t *taskRepository) CreateTask(task entity.Task) (string, error) {
	result, err := t.db.Exec("INSERT INTO tasks (title, content) VALUES (?, ?)", task.Title, task.Content)
	if err != nil {
		return "", err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return "", nil
	}

	return fmt.Sprintf("%d", id), nil
}

func (t *taskRepository) ListTasks() ([]entity.Task, error) {
	rows, err := t.db.Query("SELECT * FROM tasks")
	if err != nil {
		return nil, err
	}

	var tasks []entity.Task
	for rows.Next() {
		var task entity.Task
		if err := rows.Scan(&task.ID, &task.Title, &task.Content); err != nil {
			return nil, err
		}

		tasks = append(tasks, task)
	}

	return tasks, nil
}

func (t *taskRepository) UpdateTask(task entity.Task) error {
	_, err := t.db.Exec("UPDATE tasks SET title = ?, content = ? WHERE id = ?", task.Title, task.Content, task.ID)
	if err != nil {
		return err
	}

	return nil
}
