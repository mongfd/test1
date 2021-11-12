package index

import (
	"github.com/gin-gonic/gin"
	"net/http"
)



type Index struct {
}

func (i *Index) Add(router *gin.Engine) {
	router.Handle(http.MethodGet, "")
}
