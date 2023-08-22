package api

import (
	"net/http"
	"task-management-app/internal/entity"
	"task-management-app/internal/usecase"

	"github.com/labstack/echo/v4"
)

type UserController struct {
	userUseCase usecase.UserInteractor
}

func NewUserController(userUseCase usecase.UserInteractor) *UserController {
	return &UserController{
		userUseCase: userUseCase,
	}
}

func (uc *UserController) RegisterUser(c echo.Context) error {
	var user entity.User
	if err := c.Bind(&user); err != nil {
		return c.JSON(http.StatusBadRequest, "Invalid request payload")
	}

	userID, err := uc.userUseCase.CreateUser(user.Username, user.Email, user.Password)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, "Failed to create user")
	}

	return c.JSON(http.StatusCreated, userID)
}

func (uc *UserController) LoginUser(c echo.Context) error {
	var loginData struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}

	if err := c.Bind(&loginData); err != nil {
		return c.JSON(http.StatusBadRequest, "Invalid request payload")
	}

	user, err := uc.userUseCase.LoginUser(loginData.Username, loginData.Password)
	if err != nil {
		return c.JSON(http.StatusUnauthorized, "Invalid credentials")
	}

	//jwt token stuff? belom implemen

	return c.JSON(http.StatusOK, user)
}

func (uc *UserController) LogoutUser(c echo.Context) error {

	//logout stuff here

	return c.JSON(http.StatusOK, "Logged out successfully")
}

func (uc *UserController) ListUser(c echo.Context) error {
	users, err := uc.userUseCase.ListUser()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, "Failed to fetch users")
	}

	return c.JSON(http.StatusOK, users)
}
