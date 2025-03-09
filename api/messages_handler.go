package api

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/yigitkalan/chatapp/internal/message"
)

func GetAllMessagesHandler(c *gin.Context) {
	messages := message.GetAllMessages()
	c.JSON(http.StatusOK, gin.H{
		"messages": messages,
	})
}

func GetMessageByIdHandler(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		panic("error parsing parameter")
	}
	message := message.GetMessageById(uint(id))
	c.JSON(http.StatusOK, gin.H{
		"message": message,
	})
}
