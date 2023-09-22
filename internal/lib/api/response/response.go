package response

import (
	"fmt"
	"strings"

	"github.com/go-playground/validator"
)

// Response save data from server
type Response struct {
	Status string `json:"status"`
	Error  string `json:"error,omitempty"`
}

const (
	StatusOK    = "OK"
	StatusError = "Error"
)

// New is function-wrapper for convert regular error to struct
func Error(msg string) Response {
	return Response{
		Status: StatusError,
		Error:  msg,
	}
}

// OK is func for return string "OK" as struct for convert to JSON
func OK() Response {
	return Response{
		Status: StatusOK,
	}
}

// ValidationError is func for return human-readable list of errors from validator errors
func ValidationError(errs validator.ValidationErrors) Response {
	var errMsgs []string

	for _, err := range errs {
		switch err.ActualTag() {
		case "required":
			errMsgs = append(errMsgs, fmt.Sprintf("field %s is a required field", err.Field()))
		case "url":
			errMsgs = append(errMsgs, fmt.Sprintf("field %s is not a vaild URL", err.Field()))
		default:
			errMsgs = append(errMsgs, fmt.Sprintf("field %s is not vaild", err.Field()))
		}
	}

	return Response{
		Status: StatusError,
		Error:  strings.Join(errMsgs, ", "),
	}
}
