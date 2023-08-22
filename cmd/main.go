package main

import (
	"task-management-app/framework/database"
	"task-management-app/interfaces/api"
	"task-management-app/interfaces/persistence"
	"task-management-app/internal/usecase"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	db := database.Connect()

	taskRepository := persistence.NewTaskRepository(db)
	userRepository := persistence.NewUserRepository(db)

	taskUseCase := usecase.NewTaskInteractor(taskRepository)
	userUseCase := usecase.NewUserUseCase(userRepository)

	taskController := api.NewTaskController(*taskUseCase)
	userController := api.NewUserController(*userUseCase)

	e.POST("/tasks", taskController.CreateTask)
	e.GET("/tasks", taskController.ListTasks)
	e.PUT("/tasks/:id", taskController.UpdateTask)

	e.POST("/register", userController.RegisterUser)
	e.POST("/login", userController.LoginUser)
	e.POST("/logout", userController.LogoutUser)
	e.GET("/users", userController.ListUser)

	e.Start(":8080")
}
