package main

import (
	"github.com/gin-gonic/gin"
	"github.com/yigitkalan/chatapp/api"
)

func main() {
	router := gin.Default()
	router.HandleMethodNotAllowed = true
	api.MapMessageRoutes(router)


	router.Run()

}
