package handlers

import (
	"context"
	"net/http"

	"users-api/internal/core/domain/users"

	"github.com/labstack/echo/v4"
)

type FindUserHandler struct {
	processor users.FindUserProcessor
}

func NewFindUserHandler(processor users.FindUserProcessor) *FindUserHandler {
	return &FindUserHandler{
		processor: processor,
	}
}

func (h FindUserHandler) FindUser(c echo.Context) error {
	ctx := context.Background()
	users, err := h.processor.ProcessRequest(ctx)
	if err != nil {
		return h.response(c, nil, err)
	}
	return c.JSON(http.StatusOK, users)
}

func (h FindUserHandler) response(c echo.Context, response interface{}, error error) error {
	if err := c.JSON(http.StatusOK, response); err != nil {
		return err
	}
	return nil
}
