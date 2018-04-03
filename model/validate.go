package model

import (
	"github.com/go-ozzo/ozzo-validation"
)

func (resource RegisterResource) Validate() error {
	return validation.ValidateStruct(&resource,
		validation.Field(&resource.Email, validation.Required, validation.Length(5, 256)),
		validation.Field(&resource.Password, validation.Required, validation.Length(1, 512)))
}

func (resource ActivateResource) Validate() error {
	return validation.ValidateStruct(&resource,
		validation.Field(&resource.Email, validation.Required, validation.Length(5, 256)),
		validation.Field(&resource.ActivateCode, validation.Required, validation.Length(4, 12)))
}

func (resource RegisterResource) ValidateEmail() error {
	return validation.ValidateStruct(&resource,
		validation.Field(&resource.Email, validation.Required, validation.Length(5, 256)))
}

func (resource ResetPasswordResource) Validate() error {
	return validation.ValidateStruct(&resource,
		validation.Field(&resource.Email, validation.Required, validation.Length(5, 256)),
		validation.Field(&resource.Password, validation.Required, validation.Length(1, 512)),
		validation.Field(&resource.ActivateCode, validation.Required, validation.Length(4, 12)))
}

func (resource Location) Validate() error {
	return validation.ValidateStruct(&resource,
		validation.Field(&resource.Country, validation.Required, validation.Length(1, 256)),
		validation.Field(&resource.CountryCode, validation.Required, validation.Length(1, 12)),
		validation.Field(&resource.Province, validation.Required, validation.Length(1, 128)),
		validation.Field(&resource.BeginTime, validation.Required, validation.Length(1, 128)),
		validation.Field(&resource.Hometown, validation.Required, validation.Length(1, 128)))
}
