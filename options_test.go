package goi18n_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/yyle88/goi18n"
	"github.com/yyle88/neatjson/neatjsons"
	"github.com/yyle88/runpath/runtestpath"
)

func TestNewOptions(t *testing.T) {
	options := goi18n.NewOptions().WithOutputPathWithPkgName(runtestpath.SrcPath(t))
	t.Log(neatjsons.S(options))

	require.Equal(t, "goi18n", options.PkgName)
}
