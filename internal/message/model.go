package message

import (
	"database/sql"
	"github.com/yigitkalan/chatapp/internal/bases"
)

type Message struct {
	bases.Model
	Text   string
	FromId uint
	ToId   uint
	Time   sql.NullTime
}
