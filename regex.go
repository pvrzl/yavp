package yavp

import (
	"fmt"
	"regexp"
	"strings"
)

var (
	hashAlogrithmLength = map[string]int{
		"md5":       32,
		"md4":       32,
		"sha1":      40,
		"sha256":    64,
		"sha384":    96,
		"sha512":    128,
		"ripemd128": 32,
		"ripemd160": 40,
		"tiger128":  32,
		"tiger160":  40,
		"tiger192":  48,
		"crc32":     8,
		"crc32b":    8,
	}
	hashAlogrithmList = []string{
		"md5",
		"md4",
		"sha1",
		"sha256",
		"sha384",
		"sha512",
		"ripemd128",
		"ripemd160",
		"tiger128",
		"tiger160",
		"tiger192",
		"crc32",
		"crc32b",
	}
)

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

// IsHash check if the string is a hash of type algorithm. Algorithm is one of ['md4', 'md5', 'sha1', 'sha256', 'sha384', 'sha512', 'ripemd128', 'ripemd160', 'tiger128', 'tiger160', 'tiger192', 'crc32', 'crc32b']
func IsHash(algorithm string) StringValidator {
	if IsInString(hashAlogrithmList).Validate(algorithm) != nil {
		panic("unknown algorithm")
	}
	pattern := fmt.Sprintf(`^[a-fA-F0-9]{%d}$`, hashAlogrithmLength[algorithm])
	return StringValidator{
		validator: regexValidation(pattern),
		message:   ErrInvalidValue,
	}
}

func isBtcAddress(s string, e error) error {
	bech32 := `^(bc1)[a-z0-9]{25,39}$`
	base58 := `^(1|3)[A-HJ-NP-Za-km-z1-9]{25,39}$`
	if strings.HasPrefix(s, "bc1") {
		return regexValidation(bech32)(s, e)
	}
	return regexValidation(base58)(s, e)
}

func IsBtcAddress() StringValidator {
	return StringValidator{
		validator: isBtcAddress,
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

	IsEmail = RegexValidation("^[a-zA-Z0-9.!#$%&'*+/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$")

	IsIndonesiaMobilePhone = RegexValidation(`^(\+?62|0)8(1[123456789]|2[1238]|3[1238]|5[12356789]|7[78]|9[56789]|8[123456789])([\s?|\d]{5,11})$`)

	IsCreditCard = RegexValidation(`^(?:4[0-9]{12}(?:[0-9]{3,6})?|5[1-5][0-9]{14}|(222[1-9]|22[3-9][0-9]|2[3-6][0-9]{2}|27[01][0-9]|2720)[0-9]{12}|6(?:011|5[0-9][0-9])[0-9]{12,15}|3[47][0-9]{13}|3(?:0[0-5]|[68][0-9])[0-9]{11}|(?:2131|1800|35\d{3})\d{11}|6[27][0-9]{14}|^(81[0-9]{14,17}))$`)

	IsLocale = RegexValidation(`^[A-Za-z]{2,4}([_-]([A-Za-z]{4}|[\d]{3}))?([_-]([A-Za-z]{2}|[\d]{3}))?$`)

	IsBase64 = RegexValidation(`^(?:[A-Za-z0-9+/]{4})*(?:[A-Za-z0-9+/]{2}==|[A-Za-z0-9+/]{3}=)?$`)

	IsIMEI = RegexValidation(`^[0-9]{15}$`)

	IsBIC = RegexValidation(`^[A-Za-z]{6}[A-Za-z0-9]{2}([A-Za-z0-9]{3})?$`)

	IsEthereumAddress = RegexValidation(`^(0x)[0-9a-fA-F]{40}$`)
)
