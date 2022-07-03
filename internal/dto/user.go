package dto

import (
	"github.com/nakoding-community/goboil-clean/internal/model"
	res "github.com/nakoding-community/goboil-clean/pkg/util/response"

	"github.com/google/uuid"
)

// request
type (
	CreateUserRequest struct {
		Name     string  `json:"name" validate:"required"`
		Email    *string `json:"email,omitempty" validate:"omitempty"`
		Password string  `json:"password"`
	}
)

type (
	UpdateUserRequest struct {
		ID       uuid.UUID `param:"id" validate:"required"`
		Name     string    `json:"name"`
		Email    string    `json:"email" validate:"omitempty,email"`
		Password string    `json:"password"`
	}
)

// response
type (
	UserResponse struct {
		model.UserEntityModel
	}
	UserResponseDoc struct {
		Body struct {
			Meta res.Meta     `json:"meta"`
			Data UserResponse `json:"data"`
		} `json:"body"`
	}
)
