package yavp

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

const (
	validTestComment   = "should not return error when called with valid value"
	invalidTestComment = "should return error when called with invalid value"
)

func TestRegex(t *testing.T) {
	assert.Equal(t, ErrInvalidValue, IsAlpha.Validate("123"), invalidTestComment)
	assert.NoError(t, IsAlpha.Validate("abc"), validTestComment)

	assert.Equal(t, ErrInvalidValue, IsAlphaNumeric.Validate("123@@"), invalidTestComment)
	assert.NoError(t, IsAlphaNumeric.Validate("abc09"), validTestComment)

	assert.Equal(t, ErrInvalidValue, IsHexaDecimal.Validate("123GG"), invalidTestComment)
	assert.NoError(t, IsHexaDecimal.Validate("0xFF00"), validTestComment)

	assert.Equal(t, ErrInvalidValue, IsHexColor.Validate("#000G00"), invalidTestComment)
	assert.NoError(t, IsHexColor.Validate("#000000"), validTestComment)

	assert.Equal(t, ErrInvalidValue, IsRGB.Validate("rgb(255,0,a)"), invalidTestComment)
	assert.NoError(t, IsRGB.Validate("rgb(255,255,255)"), validTestComment)

	assert.Equal(t, ErrInvalidValue, IsRGBA.Validate("rgb5,0,a)"), invalidTestComment)
	assert.NoError(t, IsRGBA.Validate("rgb(255,255,255,0)"), validTestComment)

	assert.Equal(t, ErrInvalidValue, IsEmail.Validate("ef@"), invalidTestComment)
	assert.NoError(t, IsEmail.Validate("ef@ef.com"), validTestComment)
}

// func BenchmarkRegex(b *testing.B) {
// 	str := ascii()
// 	b.ResetTimer()
// 	for N := 0; N < b.N; N++ {
// 		err := RegexValidation(`[\s\S]*`).Validate(str)
// 		if err != nil {
// 			b.Log(err)
// 		}
// 	}
// }
