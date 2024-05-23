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
		return errors.New("invalid name")
	}
	return nil
}

func NewUser(id UserId, name string) *User {
	return &User{
		Id:   id,
		Name: name,
	}
}
