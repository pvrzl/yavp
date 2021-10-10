package yavp

import (
	"encoding/json"
	"errors"
	"strings"
)

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

func isContains(seed string) func(string, error) error {
	return func(s string, e error) error {
		if !strings.Contains(s, seed) {
			return e
		}
		return nil
	}
}

// IsContains is a method to check wether the string contains specifc substring
func IsContains(seed string) StringValidator {
	return StringValidator{
		validator: isContains(seed),
		message:   ErrInvalidValue,
	}
}

func isEquals(comparison string) func(string, error) error {
	return func(s string, e error) error {
		if !strings.EqualFold(s, comparison) {
			return e
		}
		return nil
	}
}

// IsEquals is a method to check wether the string equals with another string(case insensitive)
func IsEquals(comparison string) StringValidator {
	return StringValidator{
		validator: isEquals(comparison),
		message:   ErrInvalidValue,
	}
}

func isJSON() func(string, error) error {
	return func(s string, e error) error {
		if !json.Valid([]byte(s)) {
			return e
		}
		return nil
	}
}

// IsJSON is a method to check wether the string is a valid json or not
func IsJSON() StringValidator {
	return StringValidator{
		validator: isJSON(),
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
