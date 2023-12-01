package main

import (
	"H8-P1-NGC/NGC-16/config"
	"H8-P1-NGC/NGC-16/entity"
	"H8-P1-NGC/NGC-16/handler"
	"fmt"
)

func main() {
	config.ConnectDB()

	for {
		fmt.Println("Pilih opsi:")
		fmt.Println("1. Registrasi")
		fmt.Println("2. Login")
		fmt.Println("3. Keluar")

		var choice int
		fmt.Scanln(&choice)

		switch choice {
		case 1:
			var username, password string
			fmt.Println("Registrasi")
			fmt.Println("Masukkan username:")
			fmt.Scanln(&username)
			fmt.Println("Masukkan password:")
			fmt.Scanln(&password)
			// def entity
			user := entity.User{
				Username: username,
				Password: password,
			}

			// def handler check
			err := handler.Register(user)
			if err != nil {
				fmt.Println("Kesalahan saat registrasi:", err)
			} else {
				fmt.Print("Registrasi berhasil.")
			}
		case 2:
			var username, password string
			fmt.Println("Login")
			fmt.Println("Masukkan username:")
			fmt.Scanln(&username)
			fmt.Println("Masukkan password:")
			fmt.Scanln(&password)
			// def user auth w/handler
			user, isAuth := handler.Login(username, password)
			if isAuth {
				fmt.Println("Login berhasil. Selamat datang,", user.Username)
			} else {
				fmt.Println("Login gagal.")
			}

		case 3:
			fmt.Println("Keluar dari program.")
			return
		default:
			fmt.Println("Pilihan tidak valid.")
		}
	}
}
