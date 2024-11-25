package rest

import (
	"CoreBaseGo/internal/interfaces/rest/messages"
	"github.com/gin-gonic/gin"
)

// Response represents the desired response format
type Response struct {
	Status  int            `json:"status"`
	Data    interface{}    `json:"data"`
	Message *messages.Data `json:"message"`
}

// JSONOutput generates a Response struct and sends it as JSON
func JSONOutput(context *gin.Context, status int, data interface{}, messageCode int, messageText string) {
	// Create the message data
	messageData := messages.Text(messageCode, messageText)
	// Create the result
	result := &Response{
		Status:  status,
		Data:    data,
		Message: messageData,
	}
	// Send the response as JSON
	context.JSON(status, result)
	context.Abort()
}
