package domain

import "errors"

var (
	ErrRegistrationNotFound = errors.New("registration not found")
)

type ErrorResponse struct {
	ErrorMessage string `json:"errorMessage"`
}
