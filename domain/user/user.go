package user

import (
	"errors"
)

type User struct {
	Id   UserId
	Name string
}

func (u *User) Validate() error {
	if u.Name == "" {
		return errors.New("invalid email format")
	}
	return nil
}

func NewUser(id UserId, name string) *User {
	return &User{
		Id:   id,
		Name: name,
	}
}
