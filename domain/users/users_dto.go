package users

import (
	"strings"

	"github.com/SaravananPitchaimuthu/Bookstore_users-api/utils/errors"
)

type User struct {
	Id          int64
	FirstName   string
	LastName    string
	Email       string
	DateCreated string
}

// func Validate(user *User) *errors.RestErr {
// 	user.Email = strings.TrimSpace(strings.ToLower(user.Email))

// 	if user.Email == "" {
// 		return errors.NewBadRequestError("invalid email address")
// 	}

// 	return nil
// }

func (user *User) Validate() *errors.RestErr {
	user.Email = strings.TrimSpace(strings.ToLower(user.Email))

	if user.Email == "" {
		return errors.NewBadRequestError("invalid email address")
	}

	return nil
}
