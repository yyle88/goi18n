package example1generate

import (
	"testing"

	"github.com/nicksnyder/go-i18n/v2/i18n"
	"github.com/yyle88/goi18n"
	"github.com/yyle88/neatjson/neatjsons"
	"github.com/yyle88/osexistpath/osmustexist"
	"github.com/yyle88/rese"
	"github.com/yyle88/runpath"
	"github.com/yyle88/zaplog"
	"go.uber.org/zap"
	"golang.org/x/text/language"
	"gopkg.in/yaml.v3"
)

func TestGenerate(t *testing.T) {
	bundle := i18n.NewBundle(language.AmericanEnglish)
	bundle.RegisterUnmarshalFunc("yaml", yaml.Unmarshal)

	var messageFiles []*i18n.MessageFile
	for _, locale := range []string{"en-US", "zh-CN", "km-KH"} {
		path := runpath.PARENT.UpTo(1, "i18n", locale+".yaml")
		zaplog.LOG.Debug("LOAD", zap.String("path", path))

		osmustexist.MustFile(path)

		messageFile := rese.P1(bundle.LoadMessageFile(path))
		zaplog.SUG.Debugln(neatjsons.S(messageFile))

		messageFiles = append(messageFiles, messageFile)
	}
	zaplog.SUG.Debugln(neatjsons.S(bundle.LanguageTags()))

	outputPath := osmustexist.FILE(runpath.PARENT.Join("example1message/message.go"))
	goi18n.Generate(messageFiles, goi18n.NewOptions(outputPath))
}
