package helper

import "github.com/gin-gonic/gin"

type ReturnType struct {
	Code int
	Msg  string
	Data interface{}
}

func ApiReturn(code int, msg string, data interface{}) gin.H {
	return gin.H{
		"code": code,
		"msg":  msg,
		"data": data,
	}
}
