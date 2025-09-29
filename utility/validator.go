package utility

import (
	"fmt"

	"github.com/go-playground/validator/v10"
)

var validate = validator.New()

func Validate(entity any) error {
	err := validate.Struct(entity)
	if err != nil {
		if _, ok := err.(*validator.InvalidValidationError); ok {
			return err
		}

		for _, err := range err.(validator.ValidationErrors) {
			fmt.Printf("Field '%s' failed validation tag '%s'\n", err.Field(), err.Tag())
		}
		return err
	}
	return nil
}
