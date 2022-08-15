package config

import (
	"bookApp/errs"

	"github.com/joho/godotenv"
)

func GetPort() string {

	envMap, err := godotenv.Read("./.env")
	if err != nil {

		panic("Cannot find PORT to run server");
	}
	return envMap["PORT"]
}

func GetSecretKey() (string, *errs.AppError) {

	envMap, err := godotenv.Read("./.env")
	if err != nil {

		return "", errs.InternalServerError("Cannot get secret_key!!!");
	}
	return envMap["SECRET_KEY"], nil;
}
