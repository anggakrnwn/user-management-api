package helpers

import (
	"fmt"
	"strings"

	"github.com/go-playground/validator/v10"
	"gorm.io/gorm"
)

func TranslateErrorMessage(err error) map[string]string {
	errorMap := make(map[string]string)

	fmt.Println("DEBUG error isi:", err.Error())

	if validationErros, ok := err.(validator.ValidationErrors); ok {
		for _, fieldError := range validationErros {
			field := fieldError.Field()
			switch fieldError.Tag() {
			case "required":
				errorMap[field] = fmt.Sprintf("%s is required", field)
			case "email":
				errorMap[field] = "invalid email format"
			case "unique":
				errorMap[field] = fmt.Sprintf("%s already exists", field)
			case "min":
				errorMap[field] = fmt.Sprintf("%s must be at least %s characters", field, fieldError.Param())
			case "max":
				errorMap[field] = fmt.Sprintf("%s mustbe at most %s characters", field, fieldError.Param())
			case "numerc":
				errorMap[field] = fmt.Sprintf("%s must be a number", field)
			default:
				errorMap[field] = "invalid value"
			}
		}
	}

	if err != nil {
		if strings.Contains(err.Error(), "duplicate entry") {
			if strings.Contains(err.Error(), "username") {
				errorMap["username"] = "username already exists"

			}
			if strings.Contains(err.Error(), "email") {
				errorMap["email"] = "email already exists"
			}
		} else if err == gorm.ErrRecordNotFound {
			errorMap["error"] = "record not found"
		}
	}
	return errorMap
}

func IsDuplicateEntryError(err error) bool {
	return err != nil && strings.Contains(err.Error(), "duplicate entry")
}
