package utils

import (
	"os"
	"unicode"

	"github.com/yyle88/must"
	"github.com/yyle88/osexistpath/osmustexist"
	"github.com/yyle88/rese"
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

func RewriteFileKeepMode(path string, contentBytes []byte) {
	var perm os.FileMode = 0666 // default file-mode-perm
	if osmustexist.IsFile(path) {
		perm = rese.V1(os.Stat(path)).Mode()
	}
	must.Done(os.WriteFile(path, contentBytes, perm))
}
