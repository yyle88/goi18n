package utils_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/yyle88/goi18n/internal/utils"
)

func TestHasNonASCII(t *testing.T) {
	require.False(t, utils.HasNonASCII("abc"))
	require.False(t, utils.HasNonASCII("ABC123"))
	require.False(t, utils.HasNonASCII("hello_world"))
	require.True(t, utils.HasNonASCII("嘿嘿"))
	require.True(t, utils.HasNonASCII("你好world"))
	require.True(t, utils.HasNonASCII("café"))
	require.False(t, utils.HasNonASCII(""))
}

func TestHasLetterPrefix(t *testing.T) {
	// Test uppercase letters
	require.True(t, utils.HasLetterPrefix("ABC"))
	require.True(t, utils.HasLetterPrefix("Z123"))

	// Test lowercase letters
	require.True(t, utils.HasLetterPrefix("abc"))
	require.True(t, utils.HasLetterPrefix("hello"))

	// Test non-letter prefix
	require.False(t, utils.HasLetterPrefix("123abc"))
	require.False(t, utils.HasLetterPrefix("_test"))
	require.False(t, utils.HasLetterPrefix("@symbol"))

	// Test Chinese characters
	require.False(t, utils.HasLetterPrefix("你好"))
	require.False(t, utils.HasLetterPrefix("测试ABC"))

	// Test edge cases
	require.False(t, utils.HasLetterPrefix(""))
	require.True(t, utils.HasLetterPrefix("a"))
	require.True(t, utils.HasLetterPrefix("A"))
}

func TestCapitalizeFirst(t *testing.T) {
	// Test lowercase to uppercase
	require.Equal(t, "Hello", utils.CapitalizeFirst("hello"))
	require.Equal(t, "World", utils.CapitalizeFirst("world"))

	// Test already capitalized
	require.Equal(t, "ABC", utils.CapitalizeFirst("ABC"))
	require.Equal(t, "Hello", utils.CapitalizeFirst("Hello"))

	// Test Chinese characters
	require.Equal(t, "你好", utils.CapitalizeFirst("你好"))
	require.Equal(t, "测试", utils.CapitalizeFirst("测试"))

	// Test mixed content
	require.Equal(t, "Hello世界", utils.CapitalizeFirst("hello世界"))

	// Test single character
	require.Equal(t, "A", utils.CapitalizeFirst("a"))
	require.Equal(t, "B", utils.CapitalizeFirst("B"))

	// Test edge cases
	require.Equal(t, "", utils.CapitalizeFirst(""))
	require.Equal(t, "1abc", utils.CapitalizeFirst("1abc"))
	require.Equal(t, "_test", utils.CapitalizeFirst("_test"))
}

func TestDefaultUnicodeMessageName(t *testing.T) {
	// Test ASCII letter prefix
	require.Equal(t, "Hello", utils.DefaultUnicodeMessageName("hello"))
	require.Equal(t, "World", utils.DefaultUnicodeMessageName("World"))

	// Test Chinese characters - add "I" prefix
	require.Equal(t, "I你好", utils.DefaultUnicodeMessageName("你好"))
	require.Equal(t, "I测试消息", utils.DefaultUnicodeMessageName("测试消息"))

	// Test non-letter prefix - add "I" prefix
	require.Equal(t, "I123", utils.DefaultUnicodeMessageName("123"))
	require.Equal(t, "I_test", utils.DefaultUnicodeMessageName("_test"))
}

func TestDefaultUnicodeStructName(t *testing.T) {
	// Test ASCII letter prefix
	require.Equal(t, "Hello", utils.DefaultUnicodeStructName("hello"))
	require.Equal(t, "World", utils.DefaultUnicodeStructName("World"))

	// Test Chinese characters - add "P" prefix
	require.Equal(t, "P你好", utils.DefaultUnicodeStructName("你好"))
	require.Equal(t, "P测试结构", utils.DefaultUnicodeStructName("测试结构"))

	// Test non-letter prefix - add "P" prefix
	require.Equal(t, "P123", utils.DefaultUnicodeStructName("123"))
	require.Equal(t, "P_test", utils.DefaultUnicodeStructName("_test"))
}

func TestDefaultUnicodeFieldName(t *testing.T) {
	// Test ASCII letter prefix
	require.Equal(t, "Name", utils.DefaultUnicodeFieldName("name"))
	require.Equal(t, "Value", utils.DefaultUnicodeFieldName("Value"))

	// Test Chinese characters - add "V" prefix
	require.Equal(t, "V名称", utils.DefaultUnicodeFieldName("名称"))
	require.Equal(t, "V字段", utils.DefaultUnicodeFieldName("字段"))

	// Test non-letter prefix - add "V" prefix
	require.Equal(t, "V123", utils.DefaultUnicodeFieldName("123"))
	require.Equal(t, "V_field", utils.DefaultUnicodeFieldName("_field"))
}
