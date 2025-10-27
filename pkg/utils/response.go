package utils

import "github.com/gin-gonic/gin"

type successResponse struct {
	StatusCode int         `json:"status_code"`
	Message    string      `json:"message"`
	Data       interface{} `json:"data,omitempty"`
}

type errorResponse struct {
	StatusCode int         `json:"status_code"`
	Message    string      `json:"message"`
	Error      interface{} `json:"error,omitempty"`
}

func ResponseSuccess(c *gin.Context, code int, message string, data interface{}) {
	c.JSON(code, successResponse{
		StatusCode: code,
		Message:    message,
		Data:       data,
	})
}

func ResponseError(c *gin.Context, code int, message string, err error) {
	var errorMessage interface{}

	if err != nil {
		errorMessage = err.Error()
	}

	c.JSON(code, errorResponse{
		StatusCode: code,
		Message:    message,
		Error:      errorMessage,
	})
}
