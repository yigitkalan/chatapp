package user

import (
	"github.com/yigitkalan/chatapp/internal/bases"
	"github.com/yigitkalan/chatapp/internal/message"
)

type User struct {
	bases.Model
	Username         string
	SentMessages     []message.Message `gorm:"foreignKey:FromId"`
	ReceivedMessages []message.Message `gorm:"foreignKey:ToId"`
}
