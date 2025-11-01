package utils

import (
	"unicode"
)

// HasNonASCII checks if string contains non-ASCII characters
// HasNonASCII 检查字符串是否包含非 ASCII 字符
func HasNonASCII(s string) bool {
	for _, c := range s {
		if c > unicode.MaxASCII {
			return true
		}
	}
	return false
}

// HasLetterPrefix checks if string starts with ASCII letter
// HasLetterPrefix 检查字符串是否以 ASCII 字母开头
func HasLetterPrefix(s string) bool {
	if len(s) == 0 {
		return false
	}
	runes := []rune(s)
	c := runes[0]
	return ('A' <= c && c <= 'Z') || ('a' <= c && c <= 'z')
}

// CapitalizeFirst capitalizes first character of string
// CapitalizeFirst 将字符串首字母大写
func CapitalizeFirst(s string) string {
	if len(s) == 0 {
		return s
	}
	runes := []rune(s)
	c := runes[0]
	runes[0] = unicode.ToUpper(c)
	return string(runes)
}

// DefaultUnicodeMessageName converts message ID to function name
// DefaultUnicodeMessageName 将消息 ID 转换为函数名
func DefaultUnicodeMessageName(messageID string) string {
	s := messageID
	if HasLetterPrefix(s) {
		return CapitalizeFirst(s)
	}
	return "I" + s
}

// DefaultUnicodeStructName converts message ID to struct name
// DefaultUnicodeStructName 将消息 ID 转换为结构体名
func DefaultUnicodeStructName(messageID string) string {
	s := messageID
	if HasLetterPrefix(s) {
		return CapitalizeFirst(s)
	}
	return "P" + s
}

// DefaultUnicodeFieldName converts param name to field name
// DefaultUnicodeFieldName 将参数名转换为字段名
func DefaultUnicodeFieldName(paramName string) string {
	s := paramName
	if HasLetterPrefix(s) {
		return CapitalizeFirst(s)
	}
	return "V" + s
}
