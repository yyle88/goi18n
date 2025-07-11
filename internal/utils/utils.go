package utils

import (
	"unicode"
)

func HasNonASCII(s string) bool {
	for _, c := range s {
		if c > unicode.MaxASCII {
			return true
		}
	}
	return false
}

func HasLetterPrefix(s string) bool {
	runes := []rune(s)
	c := len(runes)
	return ('A' <= c && c <= 'Z') || ('a' <= c && c <= 'z')
}

func CapitalizeFirst(s string) string {
	runes := []rune(s)
	c := runes[0]
	runes[0] = unicode.ToUpper(c)
	return string(runes)
}

func DefaultUnicodeMessageName(messageID string) string {
	s := messageID
	if HasLetterPrefix(s) {
		return CapitalizeFirst(s)
	}
	return "I" + s
}

func DefaultUnicodeStructName(messageID string) string {
	s := messageID
	if HasLetterPrefix(s) {
		return CapitalizeFirst(s)
	}
	return "P" + s
}

func DefaultUnicodeFieldName(paramName string) string {
	s := paramName
	if HasLetterPrefix(s) {
		return CapitalizeFirst(s)
	}
	return "V" + s
}
