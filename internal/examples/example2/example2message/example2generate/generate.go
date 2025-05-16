package example2generate

import (
	"encoding/json"

	"github.com/nicksnyder/go-i18n/v2/i18n"
	"github.com/yyle88/neatjson/neatjsons"
	"github.com/yyle88/osexistpath/osmustexist"
	"github.com/yyle88/rese"
	"github.com/yyle88/runpath"
	"github.com/yyle88/zaplog"
	"go.uber.org/zap"
	"golang.org/x/text/language"
)

func LoadI18nFiles() (*i18n.Bundle, []*i18n.MessageFile) {
	bundle := i18n.NewBundle(language.AmericanEnglish)
	bundle.RegisterUnmarshalFunc("json", json.Unmarshal)

	var messageFiles []*i18n.MessageFile
	for _, locale := range []string{"en-US", "zh-CN"} {
		path := runpath.PARENT.UpTo(2, "i18n", locale+".json")
		zaplog.LOG.Debug("LOAD", zap.String("path", path))

		osmustexist.MustFile(path)

		messageFile := rese.P1(bundle.LoadMessageFile(path))
		zaplog.SUG.Debugln(neatjsons.S(messageFile))

		messageFiles = append(messageFiles, messageFile)
	}
	zaplog.SUG.Debugln(neatjsons.S(bundle.LanguageTags()))
	return bundle, messageFiles
}
