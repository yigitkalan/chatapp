package config

import (
	"database/sql"
	"github.com/yigitkalan/chatapp/database"
	m "github.com/yigitkalan/chatapp/internal/message"
	u "github.com/yigitkalan/chatapp/internal/user"
	"time"
)

func Migrate() {

	db := database.Connect()

	db.AutoMigrate(&m.Message{}, &u.User{})

}

func FeedSeed() {

	//Seed Data
	db := database.Connect()
	db.Create(&m.Message{Text: "deneme", Time: sql.NullTime{Time: time.Now()}})
}
