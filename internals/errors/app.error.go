package errors

import "github.com/gofiber/fiber/v2"

type FiberError = fiber.Error

type AppError struct {
	FiberError
	Fields []FieldError
}

func (e *AppError) FmtFields() []FieldError {
	return e.Fields
}

func (e *AppError) Error() string {
	return e.Message
}
