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

func maxLengthString(l int) func(string, error) error {
	return func(s string, e error) error {
		if len(s) > l {
			return e
		}

		return nil
	}
}

// MaxLengthString is a maximum length validator(string)
func MaxLengthString(l int) stringValidator {
	return stringValidator{
		validator: maxLengthString(l),
		message:   ErrInvalidValue,
	}
}

func isInString(list []string) func(string, error) error {
	return func(s string, e error) error {
		if stringInSlice(s, list) {
			return nil
		}

		return e
	}
}

// IsInString is a method to check wether the string is in the list or not
func IsInString(list []string) stringValidator {
	return stringValidator{
		validator: isInString(list),
		message:   ErrInvalidValue,
	}
}

func isNotInString(list []string) func(string, error) error {
	return func(s string, e error) error {
		if stringInSlice(s, list) {
			return e
		}

		return nil
	}
}

// IsNotInString is a method to check wether the string is not in the list or not
func IsNotInString(list []string) stringValidator {
	return stringValidator{
		validator: isNotInString(list),
		message:   ErrInvalidValue,
	}
}
