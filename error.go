package yavp

import (
	"encoding/json"
	"errors"
	"fmt"
	"sort"
	"strings"
)

type (
	// Errors is a key value error type
	Errors map[string]error
)

var (
	// ErrRequired is a error message for required value
	ErrRequired = errors.New("is required")
	// ErrLatLong is a error message for invalid lat long
	ErrLatLong = errors.New("invalid lat long")
	// ErrIPError is a error message for invalid IP address
	ErrIPError = errors.New("invalid IP address")
	// ErrInvalidValue is a error message for invalid generic value
	ErrInvalidValue = errors.New("invalid value")
)

// MergeErrors is a function that merge all of the Errors that occured during validation
func MergeErrors(errors ...error) error {
	errorsMap := make(Errors)

	for _, err := range errors {
		if err != nil {
			errorMap, ok := err.(Errors)
			if ok {
				for k, v := range errorMap {
					errorsMap[k] = v
				}
			} else {
				errorsMap["unknown"] = err
			}
		}
	}

	if len(errorsMap) == 0 {
		return nil
	}

	return errorsMap
}

// Error is a method to convert the error into a string
func (errorsMap Errors) Error() string {
	if len(errorsMap) == 0 {
		return ""
	}

	keys := make([]string, 0, len(errorsMap))
	for key := range errorsMap {
		keys = append(keys, key)
	}

	sort.Strings(keys)

	var s strings.Builder
	for i, key := range keys {
		if i > 0 {
			s.WriteString("; ")
		}
		if errs, ok := errorsMap[key].(Errors); ok {
			fmt.Fprintf(&s, "%v: (%v)", key, errs)
		} else {
			fmt.Fprintf(&s, "%v: %v", key, errorsMap[key].Error())
		}
	}

	s.WriteString(".")
	return s.String()
}

// MarshalJSON is a method that convert Errors into json
func (errorsMap Errors) MarshalJSON() ([]byte, error) {
	errors := map[string]interface{}{}
	for key, err := range errorsMap {
		if ms, ok := err.(json.Marshaler); ok {
			errors[key] = ms
		} else {
			errors[key] = err.Error()
		}
	}
	return json.Marshal(errors)
}
