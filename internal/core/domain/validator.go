package domain

import (
	"github.com/go-playground/validator"
)

type Validator struct {
	*validator.Validate
}

func NewValidator() *Validator {
	return &Validator{Validate: validator.New()}
}
