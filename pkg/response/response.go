package response

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type Field map[string]interface{}

func SuccessJSON(c *gin.Context, msg string) {
	c.JSON(http.StatusOK, gin.H{
		"message": msg,
	})
}

func SuccessJSONWithField(c *gin.Context, msg string, field Field) {
	data := make(map[string]interface{})
	data["message"] = msg
	for k, v := range field {
		data[k] = v
	}
	c.JSON(http.StatusOK, data)
}

func ErrorJSON(c *gin.Context, msg string) {
	c.JSON(http.StatusBadRequest, gin.H{
		"message": msg,
	})
}

func ErrorJSONWithField(c *gin.Context, msg string, field Field) {
	data := make(map[string]interface{})
	data["message"] = msg
	for k, v := range field {
		data[k] = v
	}
	c.JSON(http.StatusBadRequest, data)
}
