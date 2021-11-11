package main

import (
	"github.com/gin-gonic/gin"
	_ "net/http/pprof"
)

func main() {
	router := gin.Default()
	router.GET("/login",test)
	err := router.Run(":8081")
	if err != nil {
		return 
	}
}

func test(ctx *gin.Context)  {
	userName := ctx.Param("userName")
	ctx.String(200,"test"+userName)
}