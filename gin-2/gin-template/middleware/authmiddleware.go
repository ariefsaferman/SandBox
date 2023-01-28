package middleware

import (
	"log"
	"net/http"

	"git.garena.com/sea-labs-id/batch-05/gin-template/util"
	"github.com/gin-gonic/gin"
)

func AuthTokenMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		_, err := util.ParseToken(ctx)
		if err != nil {
			log.Println("INI ERROR MIDDLEWARE")
			log.Println(err)
			ctx.JSON(http.StatusUnauthorized, gin.H{
				"code":    "UNAUTHORIZED",
				"message": err.Error(),
			})
			ctx.Abort()
			return
		}
		ctx.Next()
	}
}
