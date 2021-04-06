package yavp

import (
	"errors"
	"time"
)

type timeValidator struct {
	validator func(time.Time, error) error
	message   error
}

func (tv timeValidator) WithError(message error) timeValidator {
	tv.message = message
	return tv
}

func (tv timeValidator) WithErrorMessage(message string) timeValidator {
	tv.message = errors.New(message)
	return tv
}

func (tv timeValidator) Validate(t time.Time) error {
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

func IsAfter(t time.Time) timeValidator {
	return timeValidator{
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

func IsBefore(t time.Time) timeValidator {
	return timeValidator{
		validator: isBefore(t),
		message:   ErrInvalidValue,
	}
}

func is(a time.Time) func(time.Time, error) error {
	return func(t time.Time, e error) error {
		if t.Equal(a) {
			return nil
		}

		return e
	}
}
