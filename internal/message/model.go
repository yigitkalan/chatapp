package message

import (
	"database/sql"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Message struct {
	gorm.Model
	Text   string
	FromId uuid.UUID
	ToId   uuid.UUID
	Time   sql.NullTime
}
