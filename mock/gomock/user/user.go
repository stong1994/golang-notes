package user

import (
	gomock "mocks"
)

type User struct {
	Person gomock.Male
}

func NewUser(p gomock.Male) *User {
	return &User{Person: p}
}

func (u *User) GetUserInfo(id int64) error {
	return u.Person.Get(id)
}
