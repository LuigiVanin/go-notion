package errors

import "github.com/gofiber/fiber/v2"

type AppError struct {
	fiber.Error
	fields []FieldError
}

func (e *AppError) Fields() []FieldError {
	return e.fields
}
