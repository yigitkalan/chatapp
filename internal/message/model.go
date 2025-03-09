package message

import (
	"database/sql"
	"github.com/yigitkalan/chatapp/internal/bases"
)

type Message struct {
	bases.Model
	Text   string `json:"text" binding:"required"`
	FromId uint   `json:"from_id" binding:"gte=0"`
	ToId   uint   `json:"to_id" binding:"gte=0"`
	Time   sql.NullTime 
}
