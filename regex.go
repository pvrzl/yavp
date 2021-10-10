package yavp

import "regexp"

func regexValidation(pattern string) func(string, error) error {
	return func(s string, e error) error {
		isMatch := regexp.MustCompile(pattern).MatchString
		if isMatch(s) {
			return nil
		}
		return e
	}
}

// RegexValidation is a function that detect wether the string match patter or not
func RegexValidation(pattern string) StringValidator {
	return StringValidator{
		validator: regexValidation(pattern),
		message:   ErrInvalidValue,
	}
}

var (
	// IsAlpha is a function that check wether the string is alpha or not
	IsAlpha = RegexValidation(`^[a-zA-Z]+$`)
	// IsAlphaNumeric is a function that check wether the string is alpha or not
	IsAlphaNumeric = RegexValidation(`^[a-zA-Z0-9]+$`)

	IsHexaDecimal = RegexValidation(`^(0[xX])?[0-9a-fA-F]+$`)

	IsHexColor = RegexValidation(`^#(?:[0-9a-fA-F]{3}|[0-9a-fA-F]{6})$`)

	IsRGB = RegexValidation(`^(rgb)?\(?([01]?\d\d?|2[0-4]\d|25[0-5])(\W+)([01]?\d\d?|2[0-4]\d|25[0-5])\W+(([01]?\d\d?|2[0-4]\d|25[0-5])\)?)$`)

	IsRGBA = RegexValidation(`^rgba?\((\d+),\s*(\d+),\s*(\d+)(?:,\s*(\d+(?:\.\d+)?))?\)$`)

	IsEmail                = RegexValidation("^[a-zA-Z0-9.!#$%&'*+/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$")
	IsIndonesiaMobilePhone = RegexValidation(`^(\+?62|0)8(1[123456789]|2[1238]|3[1238]|5[12356789]|7[78]|9[56789]|8[123456789])([\s?|\d]{5,11})$`)

	IsCreditCard = RegexValidation(`^(?:4[0-9]{12}(?:[0-9]{3,6})?|5[1-5][0-9]{14}|(222[1-9]|22[3-9][0-9]|2[3-6][0-9]{2}|27[01][0-9]|2720)[0-9]{12}|6(?:011|5[0-9][0-9])[0-9]{12,15}|3[47][0-9]{13}|3(?:0[0-5]|[68][0-9])[0-9]{11}|(?:2131|1800|35\d{3})\d{11}|6[27][0-9]{14}|^(81[0-9]{14,17}))$`)
)
