package models

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model `json:"gorm_._model"`
	Email      string `gorm:"not null;unique" json:"email,omitempty"`
	Username   string `gorm:"not null;unique" json:"username,omitempty"`
	Firstname  string `json:"firstname,omitempty"`
	Lastname   string `json:"lastname,omitempty"`
	index      string `json:"index,omitempty"`
	Password   string `gorm:"not null" json:"password,omitempty"`
}

//type UserRepository interface {
//	NewOrder(User) (*User, *errs.AppError)
//}

func (u User) IsMailValid() bool {
	return u.Email != ""
}
