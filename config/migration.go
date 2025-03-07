package config

import (
	"database/sql"
	"time"
	"github.com/yigitkalan/chatapp/database"
	m "github.com/yigitkalan/chatapp/internal/message"
)

func Migrate() {

	db, err := database.Connect()
	if err != nil {
		panic("failed to connect db")
	}

	db.AutoMigrate(&m.Message{})

    //Seed Data
	db.Create(&m.Message{Text: "deneme", Time: sql.NullTime{Time: time.Now()}})

}
