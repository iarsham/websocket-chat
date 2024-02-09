package models

import (
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
	"time"
)

type Users struct {
	ID       uuid.UUID `json:"id" example:"b710c584-2400-3fa2-9ebb-07eb3ed96c7d"`
	Username string    `json:"username" example:"neymar"`
	Password string    `json:"-" example:"neymarJr123"`
	JoinedAt time.Time `json:"joined_at" example:"2024-01-29T03:09:00+03:30"`
	LastSeen time.Time `json:"last_seen" example:"2024-01-29T04:14:08+03:30"`
	Verified bool      `json:"verified" example:"false"`
}

func (u *Users) ValidatePassword(plainPass string) bool {
	return bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(plainPass)) == nil
}
