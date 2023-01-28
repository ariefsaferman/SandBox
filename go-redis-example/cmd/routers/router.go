package routers

import (
	"git.garena.com/aldino.rahman/go-redis-example/config"
	"github.com/gin-gonic/gin"
)


func New() *gin.Engine {

	router:= gin.Default()
	rdc:= config.Get_Redis()

	router.GET("/ping",func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"message":"PONG",
		})
	})

	InitStudentRouter(router,rdc)

	

	return router
	
}