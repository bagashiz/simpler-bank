package helpers

import (
	"github.com/go-playground/validator/v10"
)

// ValidCurrency is a function to check if the given currency is valid.
var ValidCurrency validator.Func = func(fieldLevel validator.FieldLevel) bool {
	if currency, ok := fieldLevel.Field().Interface().(string); ok {
		switch currency {
		// constants for currency, add more if needed
		case "USD", "EUR", "IDR":
			return true
		}
	}

	return false
}
