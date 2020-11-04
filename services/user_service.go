package services

import (
	"github.com/jmlgh/bookstore_users-api/domain/users"
	"github.com/jmlgh/bookstore_users-api/utils/errors"
)

func GetUser(userId int64) (*users.User, *errors.RestErr) {
	u := &users.User{Id: userId}
	if err := u.Get(); err != nil {
		return nil, err
	}
	return u, nil
}

func CreateUser(user users.User) (*users.User, *errors.RestErr) {
	if err := user.Validate(); err != nil {
		return nil, err
	}

	if err := user.Save(); err != nil {
		return nil, err
	}

	return &user, nil
}
