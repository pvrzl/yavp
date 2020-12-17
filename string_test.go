package yavp

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRequiredString(t *testing.T) {
	errorMessage := errors.New("error message")
	assert.Error(t, RequiredString.Validate(""), "shoul return error when required called with empty value")
	assert.NoError(t, RequiredString.Validate("valid"), "shoul not return error when required called with valid value")
	assert.Equal(t, errorMessage, RequiredString.WithError(errorMessage).Validate(""), "should return with custom error")
	assert.Equal(t, errorMessage, RequiredString.WithErrorMessage(errorMessage.Error()).Validate(""), "should return with custom error")
}

func TestMinLengthString(t *testing.T) {
	assert.Equal(t, ErrInvalidValue, MinLengthString(4).Validate("hol"), "should return error when MinLengthString called with invalid value")
	assert.NoError(t, MinLengthString(4).Validate("hola"), "should not return error when MinLengthString called with valid value")
}

func TestMaxLengthString(t *testing.T) {
	assert.Equal(t, ErrInvalidValue, MaxLengthString(4).Validate("holaa"), "should return error when MinLengthString called with invalid value")
	assert.NoError(t, MaxLengthString(4).Validate("hola"), "should not return error when MinLengthString called with valid value")
}

func TestIsInString(t *testing.T) {
	assert.Equal(t, ErrInvalidValue, IsInString([]string{}).Validate("hol"), "should return error when IsInString called with invalid value")
	assert.NoError(t, IsInString([]string{"hol"}).Validate("hol"), "should not return error when IsInString called with valid value")
}

func TestIsNotInString(t *testing.T) {
	assert.Equal(t, ErrInvalidValue, IsNotInString([]string{"hol"}).Validate("hol"), "should return error when IsInString called with invalid value")
	assert.NoError(t, IsNotInString([]string{}).Validate("hol"), "should not return error when IsInString called with valid value")
}
