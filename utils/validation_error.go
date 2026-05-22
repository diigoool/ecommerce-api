package utils

import "github.com/go-playground/validator/v10"

func FormatValidationError(err error) []string {
	var errors []string

	for _, e := range err.(validator.ValidationErrors) {

		switch e.Tag() {
		case "required":
			errors = append(errors, e.Field()+" is required")
		case "gt":
			errors = append(errors, e.Field()+" must be greater than 0")
		case "gte":
			errors = append(errors, e.Field()+" must be greather equal than 0")
		case "notblank":
			errors = append(errors, e.Field()+" cannot be blank")
		}

	}
	return errors
}
