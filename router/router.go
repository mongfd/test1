package router

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"test1/router/index"
)

func init() {

	router := gin.New()
	index.Add(router)
}
