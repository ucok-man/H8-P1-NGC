package handler

import (
	"H8-P1-NGC/NGC-16/config"
	"H8-P1-NGC/NGC-16/entity"

	"golang.org/x/crypto/bcrypt"
)

func Register(user entity.User) error {
	hash, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	_, err = config.DB.Exec("INSERT INTO users(username, password) VALUES (?, ?)", user.Username, hash)
	if err != nil {
		return err
	}
	return nil
}
