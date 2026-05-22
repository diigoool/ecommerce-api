package utils

import "log"

func HandleError(err error) *AppError {

	log.Println(err)

	// kalau AppError
	if appErr, ok := err.(*AppError); ok {
		return appErr
	}

	return NewInternalServerError("ERROR")

}
