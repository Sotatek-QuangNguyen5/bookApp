package utils

import (
	"bookApp/errs"

	"golang.org/x/crypto/bcrypt"
)

func HashPassword(pass string) (string, *errs.AppError) {

	hashpass, err := bcrypt.GenerateFromPassword([]byte(pass), bcrypt.DefaultCost)
	if err != nil {

		return "", errs.InternalServerError("Cannot hash password!!!")
	}
	return string(hashpass), nil
}

func CheckPasswordHash(hashpass, pass string) bool {

	err := bcrypt.CompareHashAndPassword([]byte(hashpass), []byte(pass))
	return err == nil
}