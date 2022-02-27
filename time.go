package yavp

import (
	"errors"
	"time"
)

type TimeValidator struct {
	validator func(time.Time, error) error
	message   error
}

func NewTimeValidator(f func(time.Time, error) error, err error) TimeValidator {
	return TimeValidator{
		f,
		err,
	}
}

func (tv TimeValidator) WithError(message error) TimeValidator {
	tv.message = message
	return tv
}

func (tv TimeValidator) WithErrorMessage(message string) TimeValidator {
	tv.message = errors.New(message)
	return tv
}

func (tv TimeValidator) Validate(t time.Time) error {
	return tv.validator(t, tv.message)
}

func isAfter(a time.Time) func(time.Time, error) error {
	return func(t time.Time, e error) error {
		if t.After(a) {
			return nil
		}

		return e
	}
}

func IsAfter(t time.Time) TimeValidator {
	return TimeValidator{
		validator: isAfter(t),
		message:   ErrInvalidValue,
	}
}

func isBefore(a time.Time) func(time.Time, error) error {
	return func(t time.Time, e error) error {
		if t.Before(a) {
			return nil
		}

		return e
	}
}

func IsBefore(t time.Time) TimeValidator {
	return TimeValidator{
		validator: isBefore(t),
		message:   ErrInvalidValue,
	}
}

func isTimeEqual(a time.Time) func(time.Time, error) error {
	return func(t time.Time, e error) error {
		if t.Equal(a) {
			return nil
		}

		return e
	}
}

func IsTimeEqual(t time.Time) TimeValidator {
	return TimeValidator{
		validator: isTimeEqual(t),
		message:   ErrInvalidValue,
	}
}

func ValidateTime(key string, value time.Time, validators ...TimeValidator) error {
	errors := make(Errors)
	for _, validator := range validators {
		if err := validator.Validate(value); err != nil {
			errors[key] = err
			return errors
		}
	}

	return nil
}
