package message

import (
	"github.com/yigitkalan/chatapp/database"
	"gorm.io/gorm"
)

type MessageRepository interface {
	Save(message Message) error
	Add(message Message) error
}

type messageRepositoryImpl struct {
	connection *gorm.DB
}

func GetRepo() MessageRepository {
	db := database.Connect()
	return &messageRepositoryImpl{connection: db}
}

func (mr *messageRepositoryImpl) Add(message Message) error {

	panic("unimplemented")
}

func (mr *messageRepositoryImpl) Save(message Message) error {
	panic("unimplemented")
}
