package router

import (
	"url-shortener/cmd/app/handler"
	"url-shortener/internal/utils/response"

	"github.com/gin-gonic/gin"
)

func RegisterItemRoutes(r *gin.Engine, h *handler.ItemHandler) {
	r.GET("/", func(c *gin.Context) {
		response.ResponseSuccess(c, nil, "This is URL Shortener")
	})

	r.GET("/redirect/:slug", h.RedirectBySlug)
	r.GET("/items/:slug", h.GetOneItemBySlug)
	r.GET("/items/job/:slug", h.GetOneItemBySlugWithJob)
	r.GET("/items/cached/:slug", h.GetOneItemCachedBySlug)
	r.GET("/items/cached-job/:slug", h.GetOneItemCachedBySlugWithJob)
	r.POST("/items", h.CreateOneItem)
}
