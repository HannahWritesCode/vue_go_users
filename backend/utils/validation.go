package utils

import (
	"fmt"
	"regexp"
	"strings"

	"github.com/go-playground/validator/v10"
)

// ValidationMessages takes ValidationErrors and returns a slice of strings
// with the output error messages
func ValidationMessages(errs validator.ValidationErrors) []string {
	output := []string{}

	for _, err := range errs {
		output = append(output, buildMessage(err))
	}

	return output
}

// convertStructField takes the standard struct field in camel case and
// converts it to a more readable format
func convertStructField(s string) string {
	reg := regexp.MustCompile("([A-Z])")
	res := reg.ReplaceAllString(s, " $1")

	trimmed := strings.TrimSpace(res)
	return trimmed
}

// buildMessage takes the error and does any substitutions if needed
func buildMessage(err validator.FieldError) string {
	fieldName := err.StructField()

	readableFieldName := convertStructField(fieldName)

	param := err.Param()

	errorTag := err.Tag()

	errorText := "The %s field is %s."

	switch errorTag {
	case "required":
		errorText = "The %s field is %s."
	case "url", "email":
		errorText = "The %s field must be a valid %s."
	case "eqfield":
		return fmt.Sprintf("%s must match %s field.", readableFieldName, convertStructField(param))
	}

	return fmt.Sprintf(errorText, readableFieldName, errorTag)
}
