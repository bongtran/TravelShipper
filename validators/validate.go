package validators

import (
	"TravelShipper/model"
	"github.com/go-ozzo/ozzo-validation"
)

func ValidateRegister(resource model.RegisterResource) error {
	return validation.ValidateStruct(&resource,
		validation.Field(&resource.Email, validation.Required, validation.Length(5, 256)),
		validation.Field(&resource.Password, validation.Required, validation.Length(1, 512)))
}

func ValidateActivate(resource model.ActivateResource) error {
	return validation.ValidateStruct(&resource,
		validation.Field(&resource.Email, validation.Required, validation.Length(5, 256)),
		validation.Field(&resource.ActivateCode, validation.Required, validation.Length(4, 12)))
}

func ValidateResendActivateCode(resource model.RegisterResource) error {
	return validation.ValidateStruct(&resource,
		validation.Field(&resource.Email, validation.Required, validation.Length(5, 256)))
}

func ValidateResetPassword(resource model.ResetPasswordResource) error {
	return validation.ValidateStruct(&resource,
		validation.Field(&resource.Email, validation.Required, validation.Length(5, 256)),
		validation.Field(&resource.Password, validation.Required, validation.Length(1, 512)),
		validation.Field(&resource.ActivateCode, validation.Required, validation.Length(4, 12)))
}

//func ValidateRegister(resource model.RegisterResource) error {
//	return validation.ValidateStruct(&resource,
//		validation.Field(&resource.Email, validation.Required, validation.Length(5, 256)),
//		validation.Field(&resource.Password, validation.Required, validation.Length(1, 512)))
//}