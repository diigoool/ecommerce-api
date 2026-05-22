package utils

import (
	"strings"

	"github.com/go-playground/validator/v10"
)

var Validate = validator.New()

func init() {
	Validate.RegisterValidation("notblank", func(fl validator.FieldLevel) bool {

		value := fl.Field().String()

		return strings.TrimSpace(value) != ""
	})

}
