package main

import (
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

func main() {
	fmt.Println("--- Password Generate ---")

	password := "constraseña123"
	hash, _ := HashPassword(password)
	fmt.Println("Contraseña   : ", password)
	fmt.Println("Hash         : ", hash)
	match := CheckHashPassword([]byte(password), []byte(hash))
	fmt.Println("Coincidencia : ", match)

}

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

func CheckHashPassword(password, hash []byte) bool {
	err := bcrypt.CompareHashAndPassword(hash, password)
	return err == nil
}
