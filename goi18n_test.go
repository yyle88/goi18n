package goi18n

import (
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/yyle88/neatjson/neatjsons"
	"github.com/yyle88/runpath/runtestpath"
)

func TestNewOptions(t *testing.T) {
	options := NewOptions(runtestpath.SrcPath(t))
	t.Log(neatjsons.S(options))

	require.Equal(t, "goi18n", options.PkgName)
}
