package message

import (
	"fmt"
	"github.com/yigitkalan/chatapp/database"
	"gorm.io/gorm"
)

type MessageRepository interface {
	Add(message Message) error
	Update(message Message) error
	Delete(id uint) error
	GetById(id uint) (Message, error)
	GetAll() ([]Message, error)
}

type messageRepositoryImpl struct {
	connection *gorm.DB
}

func GetRepo() MessageRepository {
	db := database.Connect()
	return &messageRepositoryImpl{connection: db}
}

func (mr *messageRepositoryImpl) Update(message Message) error {

	var msg Message
	mr.connection.First(&msg, message.ID)
	fmt.Println(msg)
	msg = message
	if msg.ID == 0 {
		panic("Message Not Found")
	}
	result := mr.connection.Save(&msg)
	return result.Error
}

func (mr *messageRepositoryImpl) Delete(id uint) error {

	msg := Message{}
	mr.connection.First(&msg, id)
	fmt.Println(msg)
	msg.IsDeleted = true
	result := mr.connection.Save(&msg)
	return result.Error
}

func (mr *messageRepositoryImpl) Add(message Message) error {
    message.SetDefaults()
	result := mr.connection.Create(&message)
	return result.Error
}

func (mr *messageRepositoryImpl) GetAll() ([]Message, error) {
	var msgs []Message
	result := mr.connection.Find(&msgs)
	return msgs, result.Error
}

func (mr *messageRepositoryImpl) GetById(id uint) (Message, error) {
	msg := Message{}
	result := mr.connection.First(&msg, id)
	return msg, result.Error
}
