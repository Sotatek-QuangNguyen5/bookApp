package dto

import "bookApp/errs"

func CheckID(id int) *errs.AppError {

	if id == 0 {

		return errs.BadRequestError("Lost information!!!")
	}

	return nil
}