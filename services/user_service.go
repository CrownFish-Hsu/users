package services

import (
	"blog/domain/users"
	"blog/utils/errors"
	"golang.org/x/crypto/bcrypt"
)

func CreateUser(user users.User) (*users.User, *errors.RestErr) {
	if err := user.Validate(); err != nil {
		return nil, err
	}

	pwd, err := bcrypt.GenerateFromPassword([]byte(user.Password), 14)
	if err != nil {
		return nil, errors.NewBadRequestError("failed to bcrypt password.")
	}

	user.Password = string(pwd[:])
	if err := user.Save(); err != nil {
		return nil, err
	}

	return &user, nil
}
