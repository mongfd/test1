package router

import (
	"github.com/gin-gonic/gin"
	"test1/router/index"
)


func init() {

	router := gin.New()
	i := index.Index{}
	i.Add(router)
}


type Router interface {
	Add(*gin.Engine)
}