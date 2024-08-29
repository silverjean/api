package models

import (
	"api/src/security"
	"errors"
	"strings"
	"time"

	"github.com/badoux/checkmail"
)

type User struct {
	ID         uint64 	 `json:"id,omitempty"`
	Name       string 	 `json:"name,omitempty"`
	Nick       string 	 `json:"nick,omitempty"`
	Email      string 	 `json:"email,omitempty"`
	Password   string 	 `json:"password,omitempty"`
	CreateDate time.Time `json:"create_date,omitempty"`
}

func (user *User) Prepare(step string) error {
	if err := user.validate(step); err != nil {
		return err
	}

	if err := user.format(step); err != nil {
		return err
	}
	return nil
}

func (user *User) validate(step string) error {
	if user.Name == "" {
		return errors.New("field name is required and don't would be empty")
	}
	if user.Nick == "" {
		return errors.New("field nick is required and don't would be empty")
	}
	if user.Email == "" {
		return errors.New("field e-mail is required and don't would be empty")
	}

	if err := checkmail.ValidateFormat(user.Email); err != nil {
		return errors.New("invalid e-mail")
	}

	if step == "register" && user.Password == "" {
		return errors.New("field password is required and don't would be empty")
	}

	return nil;
}

func (user *User) format(step string) error {
	user.Name = strings.TrimSpace(user.Name)
	user.Nick = strings.TrimSpace(user.Nick)
	user.Email = strings.TrimSpace(user.Email)

	if step == "register" {
		passWithHash, err := security.Hash(user.Password)
		if err != nil {
			return err
		}

		user.Password = string(passWithHash)
	}

	return nil
}