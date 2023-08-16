package handlers

import (
	"context"
	"net/http"

	"users-api/internal/core/domain/users"

	"github.com/labstack/echo/v4"
)

type User struct {
	Name     string `json:"name"`
	Age      int    `json:"age"`
	Email    string `json:"email"`
	Password string `json:"password,omitempty"`
	Address  string `json:"address"`
}

type GetUsersHandler struct {
	processor users.GetUsersProcessor
}

func NewGetUsersHandler(processor users.GetUsersProcessor) *GetUsersHandler {
	return &GetUsersHandler{
		processor: processor,
	}
}

// @Param id path string true "User Identifier"
// @Success 200 {object} users.UserResponse
// @Router /users/{id} [get]
func (h GetUsersHandler) GetUsers(c echo.Context) error {
	ctx := context.Background()
	userId := c.Param("id")
	users, err := h.processor.ProcessRequest(ctx, userId)
	if err != nil {
		return h.response(c, nil, err)
	}
	return h.response(c, users, nil)
}

func (h GetUsersHandler) response(c echo.Context, response interface{}, err error) error {
	if err != nil {
		if err.Error() == "sql: no rows in result set" {
			if err := c.JSON(http.StatusOK, "{}"); err != nil {
				return err
			}
		} else {
			if err := c.JSON(http.StatusInternalServerError, err.Error()); err != nil {
				return err
			}
		}
		return nil
	}

	if err := c.JSON(http.StatusOK, response); err != nil {
		return err
	}
	return nil
}
