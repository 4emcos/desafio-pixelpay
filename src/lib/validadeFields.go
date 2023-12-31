package lib

import (
	"fmt"
	"strings"

	"github.com/go-playground/validator/v10"
)

type FieldError struct {
	Err validator.FieldError
}

func (field FieldError) String() string {
	var sb strings.Builder

	sb.WriteString("validation failed on field '" + field.Err.Field() + "'")
	sb.WriteString(", condition: " + field.Err.ActualTag())

	if field.Err.Param() != "" {
		sb.WriteString(" { " + field.Err.Param() + " }")
	}

	if field.Err.Value() != nil && field.Err.Value() != "" {
		sb.WriteString(fmt.Sprintf(", actual: %v", field.Err.Value()))
	}

	return sb.String()
}
