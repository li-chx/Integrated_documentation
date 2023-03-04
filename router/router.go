package router

import "github.com/gin-gonic/gin"

func Start() {
	e := gin.Default()

	e.Run(":8080")
}
