package issuer

import (
	"strings"
	"unicode"
)

const (
	hexDigits     = "0123456789abcdefABCDEF"
	macDelimiters = ":-.,"
	macLength     = 12
)

func ValidateMACAddress(macAddress string) bool {
	var invalidCharacter rune = -1
	macAddress = strings.Map(
		func(r rune) rune {
			switch {
			case strings.ContainsRune(hexDigits, r):
				return unicode.ToLower(r)
			case strings.ContainsRune(macDelimiters, r):
				return -1
			default:
				invalidCharacter = r
				return -1
			}
		},
		macAddress,
	)

	if invalidCharacter != -1 || len(macAddress) != macLength {
		return false
	}

	return true
}
