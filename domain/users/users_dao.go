package users

import (
	"fmt"

	"github.com/SaravananPitchaimuthu/Bookstore_users-api/utils/errors"
)

var (
	usersDB = make(map[int64]*User)
)

func (user *User) Get() *errors.RestErr {
	result := usersDB[user.Id]

	if result == nil {
		return errors.NewNotFoundError(fmt.Sprintf("user %d not found", user.Id))
	}

	user.Id = result.Id
	user.FirstName = result.FirstName
	user.LastName = result.LastName
	user.Email = result.Email
	user.DateCreated = result.DateCreated

	return nil
}

func (user *User) Save() *errors.RestErr {
	current := usersDB[user.Id]

	if current != nil {
		if user.Email == current.Email {
			return errors.NewBadRequestError(fmt.Sprintf("email id %s already exists", user.Email))
		}
		return errors.NewBadRequestError(fmt.Sprintf("user id %d already exists", user.Id))
	}

	usersDB[user.Id] = user
	return nil
}
