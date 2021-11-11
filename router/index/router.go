package index

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func addRouter(router *gin.Engine)  {
	router.Handle(http.MethodGet,"",test)
}