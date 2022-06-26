package dto

import (
	"github.com/nakoding-community/goboil-clean/internal/model"

	"github.com/google/uuid"
)

//post
type (
	CreateUserRequest struct {
		Name        string  `json:"name" validate:"required"`
		Email       *string `json:"email,omitempty" validate:"omitempty"`
		PhoneNumber string  `json:"phone_number"`
		Password    string  `json:"password"`
		Token       string  `json:"token"`
	}
)

//put
type (
	UpdateUserRequest struct {
		ID          uuid.UUID `param:"id" validate:"required"`
		Name        string    `json:"name"`
		Email       string    `json:"email" validate:"omitempty,email"`
		PhoneNumber string    `json:"phone_number"`
		Password    string    `json:"password"`
		Token       string    `json:"token"`
	}
)

//get
type ()

// response
type (
	UserResponse struct {
		model.UserEntityModel
	}
)
