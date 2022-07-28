package dto

import (
	
	"bookApp/errs"
	"net/mail"
)

func CheckID(id int) *errs.AppError {

	if id == 0 {

		return errs.BadRequestError("Lost information!!!")
	}

	return nil
}

func CheckName(name string) *errs.AppError {

	if name == "" {

		return errs.BadRequestError("Lost Name!!!")
	}
	return nil
}

func CheckPassWord(pass string) *errs.AppError {

	if pass == "" {

		return errs.BadRequestError("Lost Password!!!")
	}

	if len(pass) < 8 {

		return errs.BadRequestError("Password is too short!!!")
	}
	return nil
}

func ValidatePhone(phone string) *errs.AppError {

	if (len(phone) > 0 && len(phone) < 10) || len(phone) > 11 {

		return errs.BadRequestError("Phone number is invalid!!!")
	}
	if len(phone) == 10 || len(phone) == 11 {

		for _, val := range phone {

			num := string(val)
			if !(num >= string('0') && num <= string('9')) {

				return errs.BadRequestError("Phone number is invalid!!!")
			}
		}
	}
	return nil
}

func ValidateEmail(email string) *errs.AppError {

	_, err := mail.ParseAddress(email)
	if err != nil {

		return errs.BadRequestError("Invalid Email!!!")
	}
	return nil
}