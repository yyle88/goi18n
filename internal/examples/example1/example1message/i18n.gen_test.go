package example1message_test

import (
	"testing"

	"github.com/yyle88/goi18n"
	"github.com/yyle88/goi18n/internal/examples/example1/example1message"
	"github.com/yyle88/neatjson/neatjsons"
	"github.com/yyle88/osexistpath/osmustexist"
	"github.com/yyle88/runpath/runtestpath"
	"github.com/yyle88/zaplog"
)

func TestGenerate(t *testing.T) {
	bundle, messageFiles := example1message.LoadI18nFiles()
	zaplog.SUG.Debugln(neatjsons.S(bundle.LanguageTags()))

	outputPath := osmustexist.FILE(runtestpath.SrcPath(t))
	options := goi18n.NewOptions().WithOutputPathWithPkgName(outputPath)
	t.Log(neatjsons.S(options))
	goi18n.Generate(messageFiles, options)
}
