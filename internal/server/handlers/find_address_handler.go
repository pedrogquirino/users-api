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

// @Param userId path string true "User Identifier"
// @Success 200 {object} []address.AddressResponse
// @Router /users/{userId}/address [get]
func (h FindAddressHandler) FindAddress(c echo.Context) error {
	ctx := context.Background()
	userId := c.Param("userId")
	address, err := h.processor.ProcessRequest(ctx, userId)
	if err != nil {
		return h.response(c, nil, err)
	}
	return c.JSON(http.StatusOK, address)
}

func (h FindAddressHandler) response(c echo.Context, response interface{}, error error) error {
	if err := c.JSON(http.StatusOK, response); err != nil {
		return err
	}
	return nil
}
