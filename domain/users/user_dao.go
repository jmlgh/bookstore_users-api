package users

import (
	"fmt"
	"github.com/jmlgh/bookstore_users-api/utils/date_utils"
	"github.com/jmlgh/bookstore_users-api/utils/errors"
)

var (
	usersDB = make(map[int64]*User)
)

// Get returns a user specified by its ID
func (u *User) Get() *errors.RestErr {
	result := usersDB[u.Id]
	if result == nil {
		return errors.NewNotFoundError(fmt.Sprintf("user %d not found", u.Id))
	}
	u.Id = result.Id
	u.Email = result.Email
	u.DateCreated = result.DateCreated
	u.FirstName = result.FirstName
	u.LastName = result.LastName

	return nil
}

// Save saves a new user to the database
func (u *User) Save() *errors.RestErr {
	current := usersDB[u.Id]
	if current != nil {
		if current.Email == u.Email {
			return errors.NewBadRequestError(fmt.Sprintf("email %s already registerd", u.Email))
		}
		return errors.NewBadRequestError("user already exists")
	}

	u.DateCreated = date_utils.GetNowString()

	usersDB[u.Id] = u
	return nil
}
