package helper

import "github.com/go-playground/validator/v10"

// struct for Response data
type Response struct {
	Meta Meta        `json:"meta"`
	Data interface{} `json:"data"`
}

// struct for Meta data
type Meta struct {
	Message string `json:"message"`
	Code    int    `json:"code"`
	Status  string `json:"status"`
}

// function to validate api response
func APIResponse(message string, code int, status string, data interface{}) Response {
	meta := Meta{
		Message: message,
		Code:    code,
		Status:  status,
	}

	response := Response{
		Meta: meta,
		Data: data,
	}
	return response
}

// function to format validation error
func FormatValidationError(err error) []string {
	var errors []string
	for _, err := range err.(validator.ValidationErrors) {
		errors = append(errors, err.Error())
	}
	return errors
}