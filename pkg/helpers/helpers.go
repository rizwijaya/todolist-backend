package helpers

import (
	"github.com/go-playground/validator/v10"
)

func FormatValidationError(err error) []string {
	var errors []string
	//Ubah tipe data eror ke validator error
	for _, e := range err.(validator.ValidationErrors) {
		errors = append(errors, e.Error())
	}

	return errors
}
