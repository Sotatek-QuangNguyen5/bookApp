package config

import (

	"bookApp/errs"
	"errors"
	"time"
	"github.com/golang-jwt/jwt/v4"
)

func NewAccessJsonWebToken(data map[string]interface{}) (*string, *errs.AppError) {

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{

		"exp" : time.Now().Add(30 * time.Minute).Unix(),
		"data" : data,
	})

	secretKey, e := GetSecretKey()
	if e != nil {

		return nil, e
	}
	tokenString, err := token.SignedString([]byte(secretKey))

	if err != nil {

		return nil, errs.InternalServerError("Cannot generate token");
	}
	return &tokenString, nil
}

func NewRefreshJsonWebToken(data map[string]interface{}) (*string, *errs.AppError) {

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{

		"exp" : time.Now().Add(90 * time.Minute).Unix(),
		"data" : data,
	})
	secretKey, e := GetSecretKey()
	if e != nil {

		return nil, e
	}
	tokenString, err := token.SignedString([]byte(secretKey))

	if err != nil {

		return nil, errs.InternalServerError("Cannot generate token");
	}
	return &tokenString, nil
}

func VerifyToken(tokenString string) (jwt.MapClaims, *errs.AppError) {

	token, err := jwt.Parse(tokenString, func(t *jwt.Token) (interface{}, error) {

		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {

			return nil, errors.New("unregistered token")
		}
		secretKey, e := GetSecretKey()
		if e != nil {

			return nil, errors.New("cannot get secret key")
		}
		return []byte(secretKey), nil
	})

	if err != nil {
		v, _ := err.(*jwt.ValidationError)

		if v.Errors == jwt.ValidationErrorExpired {
			
			return nil, errs.NewUnauthenticatedError(v.Error())
		}

		return nil, errs.InternalServerError("Unexpected error when parse token: " + err.Error())
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok || !token.Valid {

		return nil, errs.NewUnauthenticatedError("Invalid token")
	}
	return claims, nil
}