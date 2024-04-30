package types

import (
	"time"
)

type UserStore interface {
	CreateUser(user User) error
	GetUserByID(id int) (*User, error)
	GetUserByEmail(email string) (*User, error)
}

// refactor json case?
// json naming convention doesnt matter, 
// code still working even if db col is snake_case and here is camel case
// investigate??
type User struct {
	ID        int    `json:"userID"`
	Username  string `json:"username"`
	Email     string `json:"email"`
	Password  string `json:"-"`
	Xp        int    `json:"xp"`
	LevelID   int    `json:"levelID"`
	Health    int    `json:"health"` 
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
	IsArchived int `json:"isArchived"`
}

type RegisterUserPayload struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Email    string `json:"email" `
}

