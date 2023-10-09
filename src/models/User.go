package models

import (
	"errors"
	"strings"
	"time"
)

type User struct {
	Id        uint64    `json:"id,omitempty"`
	Name      string    `json:"name,omitempty"`
	Nickname  string    `json:"nickname,omitempty"`
	Email     string    `json:"email,omitempty"`
	Password  string    `json:"password,omitempty"`
	CreatedAt time.Time `json:"createdAt,omitempty"`
}

func (user *User) Prepare() error {
	if err := user.validator(); err != nil {
		return err
	}

	user.Format()
	return nil
}

func (user *User) validator() error {
	if user.Name == "" {
		return errors.New("the name field is required and cannot be empty")
	}

	if user.Nickname == "" {
		return errors.New("the nickname field is required and cannot be empty")
	}

	if user.Email == "" {
		return errors.New("the email field is required and cannot be empty")
	}

	if user.Password == "" {
		return errors.New("the password field is required and cannot be empty")
	}

	return nil
}

func (user *User) Format() {
	user.Name = strings.TrimSpace(user.Name)
	user.Nickname = strings.TrimSpace(user.Nickname)
	user.Email = strings.TrimSpace(user.Email)
}
