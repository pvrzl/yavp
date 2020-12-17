package yavp

import (
	"encoding/json"
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestErrorsMethod(t *testing.T) {
	errorMessage := errors.New("error message")
	err := make(Errors)
	assert.Error(t, err, "errors type should be considered as error type")
	assert.Equal(t, "", err.Error(), "empty error should return empty string")
	err["name"] = errorMessage
	assert.Equal(t, "name: error message.", err.Error(), "format of the error is invalid for one type error")
	err["second"] = errorMessage
	assert.Equal(t, "name: error message; second: error message.", err.Error(), "format of the error is invalid for multiple type error")
	err["second"] = Errors{"inside": errorMessage}
	assert.Equal(t, "name: error message; second: (inside: error message.).", err.Error(), "format of the error is invalid for Errors inside of the error")
}

func TestMarshalJSON(t *testing.T) {
	errorMessage := errors.New("error message")
	err := make(Errors)
	err["name"] = errorMessage
	b, errorMarshal := json.Marshal(err)
	assert.NoError(t, errorMarshal, "should not return error when json marshal")
	assert.Equal(t, `{"name":"error message"}`, string(b), "format of the json is invalid for single type of error")
	err["second"] = Errors{"inside": errorMessage}
	_, errorMarshal = json.Marshal(err)
	assert.NoError(t, errorMarshal, "should not return error when the error has nested error")
}

func TestMerge(t *testing.T) {
	errorMessage := errors.New("error message")
	err := Errors{"name": errorMessage}
	secondErr := Errors{"second": errorMessage}
	merged := MergeErrors(err, secondErr, errorMessage)
	assert.Equal(t, `name: error message; second: error message; unknown: error message.`, merged.Error(), "format of merged error is invalid")
	assert.Nil(t, MergeErrors(nil, nil), "nil error should return with nil too")
}
