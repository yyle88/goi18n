package utils_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/yyle88/goi18n/internal/utils"
)

func TestHasNonASCII(t *testing.T) {
	require.False(t, utils.HasNonASCII("abc"))
	require.True(t, utils.HasNonASCII("嘿嘿"))
}
