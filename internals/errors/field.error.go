package errors

type FieldError struct {
	Field string `json:"field"`
	Tag   string `json:"tag"`
}

type FieldErrors struct {
	error
	Errors []FieldError
}

func (e *FieldErrors) Error() string {
	errorText := ""
	for _, error := range e.Errors {
		errorText += error.Field + " " + error.Tag + "\n"
	}
	return errorText
}
