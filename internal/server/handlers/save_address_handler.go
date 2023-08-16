package handlers

import (
	"context"
	"net/http"

	"users-api/internal/core/domain"
	"users-api/internal/core/domain/address"
	"users-api/internal/core/domain/customer"

	"github.com/labstack/echo/v4"
)

type SaveAddressHandler struct {
	validator *domain.Validator
	processor address.SaveAddressProcessor
}

func NewSaveAddressHandler(options *domain.Options, processor address.SaveAddressProcessor) *SaveAddressHandler {
	return &SaveAddressHandler{
		validator: options.Validator,
		processor: processor,
	}
}

func (h SaveAddressHandler) SaveAddress(c echo.Context) error {
	ctx := context.Background()
	var addressRequest address.AddressRequest
	userId := c.Param("userId")

	if err := h.validateRequest(c, &addressRequest); err != nil {
		return h.response(c, nil, err)
	}

	e := h.processor.ProcessRequest(ctx, userId, addressRequest)
	err, ok := e.(*customer.Error)
	if !ok {
		return h.response(c, nil, nil)
	}
	if err != nil {
		return h.response(c, nil, err)
	}
	return h.response(c, nil, nil)
}

func (h SaveAddressHandler) response(c echo.Context, response interface{}, error *customer.Error) error {
	if error != nil {
		if err := c.JSON(error.HttpStatusCode, error); err != nil {
			return err
		}
		return nil
	}

	if err := c.JSON(http.StatusAccepted, response); err != nil {
		return err
	}
	return nil
}

func (h SaveAddressHandler) validateRequest(c echo.Context, addressRequest *address.AddressRequest) *customer.Error {
	if err := c.Bind(&addressRequest); err != nil {
		return &customer.InvalidRequestValidationError
	}

	if err := h.validator.Struct(addressRequest); err != nil {
		return &customer.InvalidRequestValidationError
	}
	return nil
}
