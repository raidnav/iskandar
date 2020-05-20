package handlers

import "github.com/gin-gonic/gin"

type ResponseWriter struct {
	ctx *gin.Context
}

func (w *ResponseWriter) Writer(status int, data interface{}, message string) {
	w.ctx.JSON(status, gin.H{
		"code":    status,
		"data":    data,
		"message": message,
	})
}

func (w *ResponseWriter) Message(status int, message string) {
	w.ctx.JSON(status, gin.H{
		"code":    status,
		"message": message,
	})
}

func (w *ResponseWriter) Data(status int, data interface{}) {
	w.ctx.JSON(status, gin.H{
		"code": status,
		"data": data,
	})
}
