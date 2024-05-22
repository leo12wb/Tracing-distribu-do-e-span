package validator

import (
	"regexp"

	"github.com/go-playground/validator/v10"
)

func Initialize() *validator.Validate {
	validate := validator.New()

	validate.RegisterValidation("cep", validateCEP)

	return validate
}

func validateCEP(fl validator.FieldLevel) bool {
	cepRegex := regexp.MustCompile(`^\d{5}-?\d{3}$`)
	return cepRegex.MatchString(fl.Field().String())
}
