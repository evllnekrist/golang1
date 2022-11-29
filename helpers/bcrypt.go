package helpers

import (
	"fmt"

	"golang.org/x/crypto/bcrypt"

	"github.com/novalagung/gubrak"
)

func HashPass(p string) string {
	salt := 8
	password := []byte(p)
	hash, _ := bcrypt.GenerateFromPassword(password, salt)

	return string(hash)
}

func ComparePass(h, p []byte) bool {
	fmt.Println(gubrak.RandomInt(10,20))
	hash, pass := []byte(h), []byte(p)

	err := bcrypt.CompareHashAndPassword(hash, pass)

	return err == nil
}
