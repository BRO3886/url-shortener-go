package api

import (
	"net/http"

	"github.com/BRO3886/url-shortener/shortener"
	"github.com/gin-gonic/gin"
)

type RedirectHandler interface {
	Redirect(c *gin.Context)
	SetRedirect(c *gin.Context)
}

type handler struct {
	redirectService shortener.RedirectService
}

// Redirect implements RedirectHandler
func (h *handler) Redirect(c *gin.Context) {
	code := c.Param("code")
	if code == "" {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "code is required",
		})
		return
	}

	redirect, err := h.redirectService.Find(code)
	if err != nil {
		ErrView(c, err)
		return
	}

	c.Redirect(http.StatusMovedPermanently, redirect.URL)
}

// SetRedirect implements RedirectHandler
func (h *handler) SetRedirect(c *gin.Context) {
	redirect := &shortener.Redirect{}

	if err := c.BindJSON(redirect); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"code":  http.StatusBadRequest,
			"error": err.Error(),
		})
		return
	}

	redirect, err := h.redirectService.Create(redirect)
	if err != nil {
		ErrView(c, err)
		return
	}

	c.JSON(http.StatusCreated, redirect)

}

func NewHandler(redirectSvc shortener.RedirectService) RedirectHandler {
	return &handler{
		redirectService: redirectSvc,
	}
}

// func setupResponse(c *gin.Context, contentType string) {
// 	c.Header("Content-Type", contentType)
// }

// // useless
// func (h *handler) serializer(contentType string) shortener.RedirectSerializer {
// 	if contentType == "application/json" {
// 		return &json.Redirect{}
// 	}
// 	// TODO: add support for other serializers
// 	return &json.Redirect{}
// }
