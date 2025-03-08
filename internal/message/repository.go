package message

import (
	"database/sql"
	"time"
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
	result := mr.connection.Save(&message)
	return result.Error
}

func (mr *messageRepositoryImpl) Delete(id uint) error {

	msg := Message{}
	result := mr.connection.Find(&msg, id)
	msg.DeletedAt = sql.NullTime{Time: time.Now()}
	return result.Error
}

func (mr *messageRepositoryImpl) Add(message Message) error {
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
