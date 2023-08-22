package api

import (
	"net/http"
	"task-management-app/internal/usecase"

	"github.com/labstack/echo/v4"
)

type TaskController struct {
	taskUseCase usecase.TaskInteractor
}

func NewTaskController(taskUseCase usecase.TaskInteractor) *TaskController {
	return &TaskController{
		taskUseCase: taskUseCase,
	}
}

func (tc *TaskController) CreateTask(c echo.Context) error {
	var taskRequest struct {
		Title   string `json:"title"`
		Content string `json:"content"`
	}

	if err := c.Bind(&taskRequest); err != nil {
		return c.JSON(http.StatusBadRequest, "Invalid request payload")
	}

	taskID, err := tc.taskUseCase.CreateTask(taskRequest.Title, taskRequest.Content)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())

	}

	return c.JSON(http.StatusCreated, map[string]string{"taskID": taskID})
}

func (tc *TaskController) ListTasks(c echo.Context) error {
	tasks, err := tc.taskUseCase.ListTasks()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, "Failed to retrieve tasks")
	}

	return c.JSON(http.StatusOK, tasks)
}

func (tc *TaskController) UpdateTask(c echo.Context) error {
	var taskRequest struct {
		Title   string `json:"title"`
		Content string `json:"content"`
	}

	if err := c.Bind(&taskRequest); err != nil {
		return c.JSON(http.StatusBadRequest, "Invalid request payload")
	}

	taskID := c.Param("id")

	err := tc.taskUseCase.UpdateTask(taskID, taskRequest.Title, taskRequest.Content)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, "Failed to update task")
	}

	return c.JSON(http.StatusOK, "Task updated successfully")
}
