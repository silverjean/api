package models

import (
	"errors"
	"strings"
	"time"
)

type User struct {
	ID         uint64 	 `json:"id,omitempty"`
	Name       string 	 `json:"name,omitempty"`
	Nick       string 	 `json:"nick,omitempty"`
	Email      string 	 `json:"email,omitempty"`
	Password   string 	 `json:"password,omitempty"`
	CreateDate time.Time `json:"create_date,omitempty"`
}

func (user *User) Prepare() error {
	if err := user.validate(); err != nil {
		return err
	}

	user.format()
	return nil
}

func (user *User) validate() error {
	if user.Name == "" {
		return errors.New("field name is required and don't would be empty")
	}
	if user.Nick == "" {
		return errors.New("field nick is required and don't would be empty")
	}
	if user.Email == "" {
		return errors.New("field e-mail is required and don't would be empty")
	}
	if user.Password == "" {
		return errors.New("field password is required and don't would be empty")
	}

	return nil;
}

func (user *User) format() {
	user.Name = strings.TrimSpace(user.Name)
	user.Nick = strings.TrimSpace(user.Nick)
	user.Email = strings.TrimSpace(user.Email)
}