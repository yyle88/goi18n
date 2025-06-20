package goi18n_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/yyle88/goi18n"
	"github.com/yyle88/must/muststrings"
	"github.com/yyle88/neatjson/neatjsons"
	"github.com/yyle88/osexistpath/osmustexist"
	"github.com/yyle88/runpath"
	"github.com/yyle88/runpath/runtestpath"
	"github.com/yyle88/syntaxgo"
)

func TestNewOptions(t *testing.T) {
	options := goi18n.NewOptions().WithOutputPathWithPkgName(runtestpath.SrcPath(t))
	t.Log(neatjsons.S(options))

	require.Equal(t, "goi18n", options.GetPkgName())
}

func TestOptions_WithOutputPath(t *testing.T) {
	options := goi18n.NewOptions().WithOutputPath(runpath.PARENT.Join("/output/message.go"))
	t.Log(neatjsons.S(options))

	muststrings.HasSuffix(options.GetOutputPath(), "goi18n/output/message.go")
}

func TestOptions_WithPkgName(t *testing.T) {
	options := goi18n.NewOptions()

	path := osmustexist.FILE(runtestpath.SrcPath(t))
	pkgName := syntaxgo.GetPkgName(path)
	options.WithOutputPath(path).WithPkgName(pkgName)

	t.Log(neatjsons.S(options))
	require.Equal(t, "goi18n", options.GetPkgName())
}
