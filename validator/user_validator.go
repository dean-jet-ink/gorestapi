package validator

import (
	"gorestapi/model"

	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/go-ozzo/ozzo-validation/v4/is"
)

type IUserValidator interface {
	UserValidate(user *model.User) error
}

type UserValidator struct {
}

func NewUserValidator() IUserValidator {
	return &UserValidator{}
}

func (uv *UserValidator) UserValidate(user *model.User) error {
	return validation.ValidateStruct(
		user,
		validation.Field(
			&user.Email,
			validation.Required.Error("email is required"),
			is.Email.Error("invalid email"),
		),
		validation.Field(
			&user.Password,
			validation.Required.Error("password is required"),
			validation.RuneLength(6, 30).Error("password must be 6-30 characters"),
		),
	)
}
