package goi18n_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/yyle88/goi18n"
)

func TestNewPluralConfig(t *testing.T) {
	param := goi18n.NewPluralConfig(1)
	t.Log(param.PluralCount)
	require.Equal(t, 1, param.PluralCount)
}
