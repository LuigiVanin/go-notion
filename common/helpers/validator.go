package helpers

import "github.com/go-playground/validator/v10"

func ValidateJson[T interface{}](data *T) error {
	validate := validator.New()
	if err := validate.Struct(data); err != nil {
		if _, ok := err.(*validator.InvalidValidationError); ok {
			return err
		}
		errors := err.(validator.ValidationErrors)
		return errors
		
	}
	return nil
}