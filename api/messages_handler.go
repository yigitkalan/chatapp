package api

import (
	"net/http"
	"strconv"
	"github.com/gin-gonic/gin"
	m "github.com/yigitkalan/chatapp/internal/message"
)

func MapMessageRoutes(router *gin.Engine){
	router.GET("/api/messages", GetAllMessagesHandler)
	router.POST("/api/messages/create", CreateMessageHandler)
	router.GET("/api/messages/:id", GetMessageByIdHandler)
}



func GetAllMessagesHandler(c *gin.Context) {
	messages := m.GetAllMessages()
	c.JSON(http.StatusOK, gin.H{
		"messages": messages,
	})
}

func CreateMessageHandler(c *gin.Context) {
	var newMsg m.Message
	if err := c.ShouldBind(&newMsg); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"Result": err.Error(),
		})
		return
	}
    m.CreateMessage(newMsg)
	c.JSON(http.StatusOK, gin.H{
		"Result": "Success",
	})
}

func GetMessageByIdHandler(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		panic("error parsing parameter")
	}
	message := m.GetMessageById(uint(id))
	c.JSON(http.StatusOK, gin.H{
		"message": message,
	})
}
