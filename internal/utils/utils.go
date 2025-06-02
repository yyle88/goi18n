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
