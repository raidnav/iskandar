package helper

import "github.com/gin-gonic/gin"

type ResponseWriter struct {
	Context *gin.Context
}

func (w ResponseWriter) Writer(status int, data interface{}, message string) *gin.Context {
	w.Context.JSON(status, gin.H{
		"code":    status,
		"data":    data,
		"message": message,
	})
	return w.Context
}

func (w ResponseWriter) Message(status int, message string) *gin.Context {
	w.Context.JSON(status, gin.H{
		"code":    status,
		"message": message,
	})
	return w.Context
}

func (w ResponseWriter) Data(status int, data interface{}) *gin.Context {
	w.Context.JSON(status, gin.H{
		"code": status,
		"data": data,
	})
	return w.Context
}
