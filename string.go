package yavp

import "errors"

type StringValidator struct {
	validator func(string, error) error
	message   error
}

func (sv StringValidator) WithError(message error) StringValidator {
	sv.message = message
	return sv
}

func (sv StringValidator) WithErrorMessage(message string) StringValidator {
	sv.message = errors.New(message)
	return sv
}

func (sv StringValidator) Validate(s string) error {
	return sv.validator(s, sv.message)
}

func requiredString(s string, e error) error {
	if s == "" {
		return e
	}

	return nil
}

// RequiredString is a require validator(string)
var RequiredString = StringValidator{
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
func MinLengthString(l int) StringValidator {
	return StringValidator{
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
func MaxLengthString(l int) StringValidator {
	return StringValidator{
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
func IsInString(list []string) StringValidator {
	return StringValidator{
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
func IsNotInString(list []string) StringValidator {
	return StringValidator{
		validator: isNotInString(list),
		message:   ErrInvalidValue,
	}
}

func ValidateString(key string, value string, validators ...StringValidator) error {
	errors := make(Errors)
	for _, validator := range validators {
		if err := validator.Validate(value); err != nil {
			errors[key] = err
			return errors
		}
	}

	return nil
}
