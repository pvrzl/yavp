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

	assert.Equal(t, ErrInvalidValue, IsCreditCard.Validate("foo"), invalidTestComment)
	assert.NoError(t, IsCreditCard.Validate("375556917985515"), validTestComment)
	assert.NoError(t, IsCreditCard.Validate("4716461583322103"), validTestComment)

	assert.Equal(t, ErrInvalidValue, IsLocale.Validate("12"), invalidTestComment)
	assert.NoError(t, IsLocale.Validate("es_ES"), validTestComment)
	assert.NoError(t, IsLocale.Validate("en"), validTestComment)
	assert.NoError(t, IsLocale.Validate("uz_Latn_UZ"), validTestComment)

	assert.Equal(t, ErrInvalidValue, IsBase64.Validate("Vml2YW11cyBmZXJtZtesting123"), invalidTestComment)
	assert.NoError(t, IsBase64.Validate("TG9yZW0gaXBzdW0gZG9sb3Igc2l0IGFtZXQsIGNvbnNlY3RldHVyIGFkaXBpc2NpbmcgZWxpdC4="), validTestComment)
	assert.NoError(t, IsBase64.Validate("Zg=="), validTestComment)

	assert.Equal(t, ErrInvalidValue, IsHash("md5").Validate("q94375dj93458w34"), invalidTestComment)
	assert.NoError(t, IsHash("md5").Validate("d94f3f016ae679c3008de268209132f2"), validTestComment)

	assert.Equal(t, ErrInvalidValue, IsIMEI.Validate("q94375dj93458w34"), invalidTestComment)
	assert.NoError(t, IsIMEI.Validate("352099001761481"), validTestComment)

	assert.Equal(t, ErrInvalidValue, IsBIC.Validate("SBIC23NXXX"), invalidTestComment)
	assert.NoError(t, IsBIC.Validate("SBICKEN1345"), validTestComment)

	assert.Equal(t, ErrInvalidValue, IsBtcAddress().Validate("4J98t1WpEZ73CNmQviecrnyiWrnqRhWNLy"), invalidTestComment)
	assert.NoError(t, IsBtcAddress().Validate("1MUz4VMYui5qY1mxUiG8BQ1Luv6tqkvaiL"), validTestComment)

	assert.Equal(t, ErrInvalidValue, IsEthereumAddress.Validate("0xGHIJK05pwm37asdf5555QWERZCXV2345AoEuIdHt"), invalidTestComment)
	assert.NoError(t, IsEthereumAddress.Validate("0x0000000000000000000000000000000000000001"), validTestComment)
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
