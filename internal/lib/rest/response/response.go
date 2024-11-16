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

func Success() Response {
	return Response{Status: StatusSuccess}
}

func Error(err error) Response {
	return Response{Status: StatusError, Error: err.Error()}
}

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
