package handlers

import (
	"context"
	"net/http"

	"users-api/internal/core/domain/address"

	"github.com/labstack/echo/v4"
)

type FindAddressHandler struct {
	processor address.FindAddressProcessor
}

func NewFindAddressHandler(processor address.FindAddressProcessor) *FindAddressHandler {
	return &FindAddressHandler{
		processor: processor,
	}
}

func (h FindAddressHandler) FindAddress(c echo.Context) error {
	ctx := context.Background()
	userId := c.Param("userId")
	users, err := h.processor.ProcessRequest(ctx, userId)
	if err != nil {
		return h.response(c, nil, err)
	}
	return c.JSON(http.StatusOK, users)
}

func (h FindAddressHandler) response(c echo.Context, response interface{}, error error) error {
	if err := c.JSON(http.StatusOK, response); err != nil {
		return err
	}
	return nil
}
