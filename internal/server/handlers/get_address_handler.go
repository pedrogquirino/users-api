package handlers

import (
	"context"
	"net/http"

	"users-api/internal/core/domain/address"

	"github.com/labstack/echo/v4"
)

type GetAddresssHandler struct {
	processor address.GetAddressProcessor
}

func NewGetAddresssHandler(processor address.GetAddressProcessor) *GetAddresssHandler {
	return &GetAddresssHandler{
		processor: processor,
	}
}

func (h GetAddresssHandler) GetAddress(c echo.Context) error {
	ctx := context.Background()
	userId := c.Param("userId")
	addressId := c.Param("addressId")
	users, err := h.processor.ProcessRequest(ctx, addressId, userId)
	if err != nil {
		return h.response(c, nil, err)
	}
	return h.response(c, users, nil)
}

func (h GetAddresssHandler) response(c echo.Context, response interface{}, err error) error {
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
