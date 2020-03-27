package exceptions

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

// ErrorWrapper declare json structure
type ErrorWrapper struct {
	Err Error `json:"error"`
}

// Error json format of error
type Error struct {
	Code    int    `json:"code"`
	Status  string `json:"status"`
	Message string `json:"message"`
}

// Error method for error
func (e Error) Error() string {
	return fmt.Sprintf("Code: %d Status: %s Message: %s", e.Code, e.Status, e.Message)
}

// Error method for error wrapper
func (e ErrorWrapper) Error() string {
	return fmt.Sprintf("Code: %d Status: %s Message: %s", e.Err.Code, e.Err.Status, e.Err.Message)
}

// Response method
func (e ErrorWrapper) Response(c *gin.Context) {
	c.Header("Error-Message", e.Err.Message)
	c.JSON(e.Err.Code, e)
}

// BadRequest return incorrect request error json
func BadRequest(message string) ErrorWrapper {
	e := Error{
		Code:    http.StatusBadRequest,
		Status:  http.StatusText(http.StatusBadRequest),
		Message: message,
	}

	return ErrorWrapper{e}
}
