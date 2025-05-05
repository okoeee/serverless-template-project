package models

import "github.com/google/uuid"

type UserId uuid.UUID

func newUserId() UserId {
	return UserId(uuid.New())
}

func (id UserId) String() string {
	return uuid.UUID(id).String()
}

func ParseUserId(id string) (UserId, error) {
	userId, err := uuid.Parse(id)
	if err != nil {
		return UserId{}, err
	}
	return UserId(userId), nil
}

type User struct {
	UserId UserId
	Name   string
	Email  string
}

func NewUser(name, email string) *User {
	user := User{
		UserId: newUserId(),
		Name:   name,
		Email:  email,
	}
	return &user
}
