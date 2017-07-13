package models

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres" // for gorm
	"time"
)

// Session : for user sessions
type Session struct {
	Token        string    `gorm:"primary_key;text" json:"token"`
	UserID       uint      `gorm:"not null" json:"user_id"`
	ClientIP     string    `gorm:"text" json:"client_ip"`
	LastActivity time.Time `json:"last_activity"`
}

// SessionsAPI : for interacting with session models
type SessionsAPI struct {
	db *gorm.DB
}

// Create : create a user session in the db
func (sesh SessionsAPI) Create(user User, token, client string) {

	sesh.db.Create(&Session{
		UserID:       user.ID,
		Token:        token,
		ClientIP:     client,
		LastActivity: time.Now(),
	})

}

// ByToken : get a session by its token
func (sesh SessionsAPI) ByToken(token uint) Session {

	var session Session
	sesh.db.Find(&session, token)
	return session

}

// Activity : update a sessions last_activity
func (sesh SessionsAPI) Activity(s *Session) {

	sesh.db.Update(&s, "last_activity", time.Now())

}
