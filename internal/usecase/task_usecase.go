package usecase

import (
	"task-management-app/interfaces/persistence"
	"task-management-app/internal/entity"
)

type TaskInteractor struct {
	taskRepository persistence.TaskRepository
}

func NewTaskInteractor(taskRepository persistence.TaskRepository) *TaskInteractor {
	return &TaskInteractor{
		taskRepository: taskRepository,
	}
}

func (t *TaskInteractor) CreateTask(title, content string) (string, error) {
	task := entity.Task{
		Title:   title,
		Content: content,
	}

	createdTaskID, err := t.taskRepository.CreateTask(task)
	if err != nil {
		return "", err
	}

	return createdTaskID, nil
}

func (t *TaskInteractor) ListTasks() ([]entity.Task, error) {
	tasks, err := t.taskRepository.ListTasks()
	if err != nil {
		return nil, err
	}

	return tasks, nil
}

func (t *TaskInteractor) UpdateTask(id, title, content string) error {
	task := entity.Task{
		ID:      id,
		Title:   title,
		Content: content,
	}

	err := t.taskRepository.UpdateTask(task)
	if err != nil {
		return err
	}

	return nil
}
