package handlers

import (
	"fmt"

	"github.com/go-playground/validator/v10"
)

var requestValidator = validator.New()

func validateStruct(payload interface{}) []string {
	if err := requestValidator.Struct(payload); err != nil {
		validationErrors, ok := err.(validator.ValidationErrors)
		if !ok {
			return []string{err.Error()}
		}
		errs := make([]string, 0, len(validationErrors))
		for _, item := range validationErrors {
			errs = append(errs, fmt.Sprintf("%s failed on %s", item.Field(), item.Tag()))
		}
		return errs
	}
	return nil
}
