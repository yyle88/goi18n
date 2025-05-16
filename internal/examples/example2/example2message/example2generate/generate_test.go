package example2generate_test

import (
	"testing"

	"github.com/yyle88/goi18n"
	"github.com/yyle88/goi18n/internal/examples/example2/example2message/example2generate"
	"github.com/yyle88/neatjson/neatjsons"
	"github.com/yyle88/osexistpath/osmustexist"
	"github.com/yyle88/runpath"
	"github.com/yyle88/zaplog"
)

func TestGenerate(t *testing.T) {
	bundle, messageFiles := example2generate.LoadI18nFiles()
	zaplog.SUG.Debugln(neatjsons.S(bundle.LanguageTags()))

	outputPath := osmustexist.FILE(runpath.PARENT.UpTo(1, "message.go"))
	options := goi18n.NewOptions().WithOutputPathWithPkgName(outputPath)
	t.Log(neatjsons.S(options))
	goi18n.Generate(messageFiles, options)
}
