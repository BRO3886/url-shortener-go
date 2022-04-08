package api

import (
	"net/http"

	"github.com/BRO3886/url-shortener/shortener"
	"github.com/gin-gonic/gin"
)

type errView struct {
	code int
	err  string
}

var errMap = map[error]int{
	shortener.ErrRedirectNotFound: http.StatusNotFound,
	shortener.ErrRedirectInvalid:  http.StatusBadRequest,
}

func ErrView(c *gin.Context, err error) {
	code, ok := errMap[err]
	if !ok {
		code = http.StatusInternalServerError
	}

	c.AbortWithStatusJSON(code, errView{
		code: code,
		err:  err.Error(),
	})
}
