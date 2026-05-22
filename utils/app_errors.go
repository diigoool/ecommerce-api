package utils

type AppError struct {
	Code    string
	Message string
	Status  int
}

func (e *AppError) Error() string {
	return e.Message
}

func NewBadRequestError(message string) *AppError {
	return &AppError{
		Code:    "BAD_REQUEST",
		Message: message,
		Status:  400,
	}
}

func NewNotFoundError(message string) *AppError {
	return &AppError{
		Code:    "NOT_FOUND",
		Message: message,
		Status:  404,
	}
}

func NewUnauthorizedError(message string) *AppError {
	return &AppError{
		Code:    "UNAUTHORIZED",
		Message: message,
		Status:  401,
	}
}

func NewInternalServerError(message string) *AppError {
	return &AppError{
		Code:    "INTERNAL_SERVER_ERROR",
		Message: message,
		Status:  500,
	}
}
