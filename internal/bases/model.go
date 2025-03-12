package bases

import (
	"database/sql"
	"time"
)

type Model struct {
	ID        uint `gorm:"primaryKey;autoIncrement"`
	CreatedAt sql.NullTime
	UpdatedAt sql.NullTime
	IsDeleted bool
}

func (m *Model) SetDefaults(){
    m.CreatedAt.Time = time.Now()
    m.IsDeleted = false;
}
