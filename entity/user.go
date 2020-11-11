package entity

import (
	uuid "github.com/satori/go.uuid"
)

// User is used to store the user's data
type User struct {
	UUID     uuid.UUID `json:"uuid" gorm:"primary_key"`
	Phone    string    `json:"phone" gorm:"unique;not null"`
	Name     string    `json:"name" gorm:"not null"`
	Role     string    `json:"role" gorm:"not null"`
	Password string    `json:"password" gorm:"unique;not null"`
	Timestamp
}
