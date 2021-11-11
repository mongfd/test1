package index

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type Router interface {
	Add(*gin.Engine)
}

type Index struct {
}

func (i *Index) Add(router *gin.Engine) {
	router.Handle(http.MethodGet, "")

}
