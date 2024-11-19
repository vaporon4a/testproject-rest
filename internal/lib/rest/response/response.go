package response

import (
	"strings"

	"github.com/go-playground/validator/v10"
)

type Response struct {
	Status string `json:"status"`
	Error  string `json:"error,omitempty"`
}

const (
	StatusSuccess = "success"
	StatusError   = "error"
)

// Success returns a Response with the status set to "success" and an empty
// error message.
func Success() Response {
	return Response{Status: StatusSuccess}
}

// Error returns a Response with the status set to "error" and the error
// message set to the string representation of the given error.
func Error(err error) Response {
	return Response{Status: StatusError, Error: err.Error()}
}

// ValidationError returns a Response with an error message constructed
// from the provided validation errors. It iterates over each validation
// error and creates a descriptive error message based on the validation
// tag, such as "required", "uuid4", "gte", or "oneof". The messages are
// concatenated and included in the Response's Error field.
func ValidationError(err validator.ValidationErrors) Response {
	var errors []string
	for _, err := range err {
		switch err.ActualTag() {
		case "required":
			errors = append(errors, err.Field()+" is required")
		case "uuid4":
			errors = append(errors, err.Field()+" must be a valid UUID")
		case "gte":
			errors = append(errors, err.Field()+" must be greater than or equal to 0")
		case "oneof":
			errors = append(errors, err.Field()+" must be one of "+err.Param())
		default:
			errors = append(errors, err.Field()+" is invalid")
		}
	}

	return Response{Status: StatusError, Error: strings.Join(errors, ", ")}
}
