package server

import (
	"fmt"

	"git.garena.com/aldino.rahman/go-redis-example/cmd/routers"
)


func Init()  {
	router:=routers.New()

	err:=router.Run()

	if err != nil{
		fmt.Println("error while running server", err)
		return
	}
}