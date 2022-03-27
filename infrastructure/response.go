package infrastructure

import (
	"fmt"

	"github.com/go-playground/validator/v10"
)

// Application default response.
// Allows to standardize the response from the server.
type Response struct {
	Status  bool              `json:"status"`
	Message string            `json:"message"`
	Errors  map[string]string `json:"errors"`
	Data    interface{}       `json:"data"`
}

// Allows to initialize response fields.
func (r *Response) Construct(status bool, message string, err error, data interface{}) {
	r.Status = status
	r.Message = message
	r.Errors = mapValidationErrors(err)
	r.Data = data
}

// Allows to build application response.
func BuildResponse(status bool, message string, err error, data interface{}) Response {
	response := new(Response)
	response.Construct(status, message, err, data)
	return *response
}

// Allows to build empty response with only string message.
func EmptyResponse(message string) Response {
	return BuildResponse(true, message, nil, nil)
}

// Alows to build success response.
func OkResponse(message string, data interface{}) Response {
	return BuildResponse(true, message, nil, data)
}

// Allows to build bad response.
func BadResponse(message string, err error) Response {
	return BuildResponse(false, message, err, nil)
}

// Mapping error string to readable dictionary.
func mapValidationErrors(err error) map[string]string {
	errors := make(map[string]string)
	if err == nil {
		return errors
	}

	return mapErrors(&errors, err)
}

func mapErrors(errors *map[string]string, err error) map[string]string {
	for _, e := range err.(validator.ValidationErrors) {
		field := e.StructField()
		reason := constructReason(e)
		(*errors)[field] = reason
	}
	return *errors
}

func constructReason(e validator.FieldError) string {
	reason := e.ActualTag()
	if param := e.Param(); len(param) > 0 {
		reason += fmt.Sprintf(" %s", param)
	}
	return reason
}
