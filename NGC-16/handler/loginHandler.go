package handler

import (
	"H8-P1-NGC/NGC-16/config"
	"H8-P1-NGC/NGC-16/entity"

	"golang.org/x/crypto/bcrypt"
)

func Login(username, password string) (entity.User, bool) {
	var user entity.User
	row := config.DB.QueryRow("SELECT id, username, password FROM users WHERE username=?", username)

	err := row.Scan(&user.ID, &user.Username, &user.Password)
	if err != nil {
		return user, false
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		return user, false
	}

	return user, true
}
