package yavp

import "errors"

type stringValidator struct {
	validator func(string, error) error
	message   error
}

func (sv stringValidator) WithError(message error) stringValidator {
	sv.message = message
	return sv
}

func (sv stringValidator) WithErrorMessage(message string) stringValidator {
	sv.message = errors.New(message)
	return sv
}

func (sv stringValidator) Validate(s string) error {
	return sv.validator(s, sv.message)
}

func requiredString(s string, e error) error {
	if s == "" {
		return e
	}

	return nil
}

// RequiredString is a require validator(string)
var RequiredString = stringValidator{
	validator: requiredString,
	message:   ErrRequired,
}

func minLengthString(l int) func(string, error) error {
	return func(s string, e error) error {
		if len(s) < l {
			return e
		}

		return nil
	}
}

// MinLengthString is a minimum length validator(string)
func MinLengthString(l int) stringValidator {
	return stringValidator{
		validator: minLengthString(l),
		message:   ErrInvalidValue,
	}
}
