package main

import (
	"github.com/gin-gonic/gin"
	"github.com/yigitkalan/chatapp/api"
)

func main() {
	router := gin.Default()
    router.HandleMethodNotAllowed = true

	router.GET("/api/messages", api.GetAllMessagesHandler)
    router.GET("/api/messages/:id", api.GetMessageByIdHandler)

	router.Run()

}

