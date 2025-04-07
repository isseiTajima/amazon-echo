package adapter

import (
	"amazon-go/usecase"
	"net/http"

	"github.com/labstack/echo/v4"
)

type UserController struct {
	useCase *usecase.UserUseCase
}

func NewUserController(useCase *usecase.UserUseCase) *UserController {
	return &UserController{useCase: useCase}
}

func (c *UserController) CreateUser(ctx echo.Context) error {
	req := new(createUserRequest)
	if err := ctx.Bind(req); err != nil {
		return err
	}

	id, err := c.useCase.CreateUser(req.Name, req.PhoneNumber, req.Gender)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, map[string]string{
			"error": "ユーザーの作成に失敗しました: " + err.Error(),
		})
	}

	return ctx.JSON(http.StatusCreated, createUserResponse{ID: id})
}

type createUserRequest struct {
	Name        string `json:"name"`
	PhoneNumber string `json:"phone_number"`
	Gender      string `json:"gender"`
}

type createUserResponse struct {
	ID string `json:"id"`
}
