package handlers

import (
	"context"
	"net/http"

	"users-api/internal/core/domain"
	"users-api/internal/core/domain/customer"
	"users-api/internal/core/domain/users"

	"github.com/labstack/echo/v4"
)

type SaveUserHandler struct {
	validator *domain.Validator
	processor users.SaveUserProcessor
}

// @Param RequestBody body users.UserRequest true "Request Body""
// @Success 201
// @Router /users [post]
func NewSaveUserHandler(options *domain.Options, processor users.SaveUserProcessor) *SaveUserHandler {
	return &SaveUserHandler{
		validator: options.Validator,
		processor: processor,
	}
}

func (h SaveUserHandler) SaveUser(c echo.Context) error {
	ctx := context.Background()
	var userRequest users.UserRequest

	if err := h.validateRequest(c, &userRequest); err != nil {
		return h.response(c, nil, err)
	}

	e := h.processor.ProcessRequest(ctx, userRequest)
	err, ok := e.(*customer.Error)
	if !ok {
		return h.response(c, nil, nil)
	}
	if err != nil {
		return h.response(c, nil, err)
	}
	return h.response(c, nil, nil)
}

func (h SaveUserHandler) response(c echo.Context, response interface{}, error *customer.Error) error {
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

func (h SaveUserHandler) validateRequest(c echo.Context, userRequest *users.UserRequest) *customer.Error {
	if err := c.Bind(&userRequest); err != nil {
		return &customer.InvalidRequestValidationError
	}

	if err := h.validator.Struct(userRequest); err != nil {
		return &customer.InvalidRequestValidationError
	}
	return nil
}
