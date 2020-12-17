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
	assert.Equal(t, ErrInvalidValue, MinLengthString(4).Validate("hol"), "shoul return error when TestMingLengthString called with invalid value")
	assert.NoError(t, MinLengthString(4).Validate("hola"), "shoul not return error when TestMingLengthString called with valid value")
}
