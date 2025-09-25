package util

import (
	"fmt"
	"strings"

	"github.com/go-playground/validator/v10"
)

var Validate = validator.New()

func ValidationErrorsToMap(err error) map[string]string {
	tagMessages := map[string]string{
		"required": "is required",
		"email":    "must be valid email",
	}

	validationErrors := err.(validator.ValidationErrors)

	errorMaps := make(map[string]string)

	for _, e := range validationErrors {
		msg, exist := tagMessages[e.Tag()]
		if !exist {
			msg = fmt.Sprintf("failed on the %s rule", e.Tag())
		}
		errorMaps[strings.ToLower(e.Field())] = e.Field() + " " + msg
	}
	return errorMaps
}
