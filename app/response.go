package app

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

// Application default response.
// Allows to standardize the response from the server.
type Response struct {
	Status  bool        `json:"status"`
	Message string      `json:"message"`
	Errors  interface{} `json:"errors"`
	Data    interface{} `json:"data"`
}

// Allows to initialize response fields.
func (r *Response) Construct(status bool, message string, errors interface{}, data interface{}) {
	r.Status = status
	r.Message = message
	r.Errors = errors
	r.Data = data
}

// Allows to build application response.
func BuildResponse(status bool, message string, errors interface{}, data interface{}) Response {
	response := new(Response)
	response.Construct(status, message, errors, data)
	return *response
}

// Allows to build empty response with only string message.
func EmptyResponse(status bool, message string) Response {
	return BuildResponse(status, message, nil, nil)
}

// Alows to build success response.
func OkResponse(message string, data interface{}) Response {
	return BuildResponse(true, message, nil, data)
}

// Allows to build bad response.
func BadResponse(message string, errors interface{}) Response {
	return BuildResponse(false, message, errors, nil)
}

// Mapping http context errors to readable dictionary.
func MapContextErrors(ctx *gin.Context) map[string]string {
	errors := make(map[string]string)
	if ctx.Errors == nil {
		return errors
	}

	for _, err := range ctx.Errors {
		mapErrors(&errors, err.Err)
	}

	return errors
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
