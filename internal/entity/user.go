package entity

import (
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	ID       int64
	Username string
	Email    string
	Password string
}

func (u *User) HashPassword() error {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(u.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	u.Password = string(hashedPassword)
	return nil
}
