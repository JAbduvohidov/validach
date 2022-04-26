package main

import (
	"fmt"
	"github.com/go-playground/validator/v10"
)

type Error struct {
	FieldName      string      `json:"field_name"`
	FieldType      string      `json:"field_type"`
	ExpectedType   string      `json:"expected_type"`
	ParamsBoundary string      `json:"params_boundary"`
	FoundValue     interface{} `json:"found_value"`
}

var validate *validator.Validate

func init() {
	validate = validator.New()
}

func Validate(s interface{}) (errs []*Error) {
	err := validate.Struct(s)
	if err == nil {
		return
	}

	if a, ok := err.(*validator.InvalidValidationError); ok {
		errs = append(errs, &Error{
			FieldType:  a.Type.String(),
			FoundValue: a.Error(),
		})
		return
	}

	for _, err := range err.(validator.ValidationErrors) {
		errs = append(errs, newErrorFromValidationError(err))
	}
	return
}

func newErrorFromValidationError(err validator.FieldError) *Error {
	return &Error{
		FieldName:      err.Field(),
		FieldType:      err.Kind().String(),
		ExpectedType:   err.Tag(),
		ParamsBoundary: err.Param(),
		FoundValue:     err.Value(),
	}
}

func (e Error) String() string {
	return fmt.Sprintf(
		"FieldName: '%s' FieldType: '%s' ExpectedType: '%s', ParamsBoundary: '%s' FoundValue: '%v'",
		e.FieldName,
		e.FieldType,
		e.ExpectedType,
		e.ParamsBoundary,
		e.FoundValue,
	)
}
