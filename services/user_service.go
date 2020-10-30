package services

import (
	"github.com/jmlgh/bookstore_users-api/domain/users"
	"github.com/jmlgh/bookstore_users-api/utils/errors"
)

func CreateUser(user users.User) (*users.User, *errors.RestErr) {
	return &user, nil
}
