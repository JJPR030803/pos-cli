package config

import (
	"fmt"
	"strings"
)

// ValidationErrors represents multiple errors
type ValidationErrors struct {
	Errors []FieldErrors
}

type FieldErrors struct {
	Field   string
	Message string
}

func (ve *ValidationErrors) Error() string {
	if len(ve.Errors) == 0 {
		return "no validation errors found"
	}
	var sb strings.Builder
	sb.WriteString("Configuration validation failed: \n")

	for _, err := range ve.Errors {
		sb.WriteString(fmt.Sprintf("   ---- \t%s: %s\n", err.Field, err.Message))
	}
	return sb.String()
}

func (ve *ValidationErrors) Add(field, message string) {
	ve.Errors = append(ve.Errors, FieldErrors{
		Field:   field,
		Message: message,
	})
}

func (ve *ValidationErrors) HasErrors() bool {
	return len(ve.Errors) > 0
}
