package example2message

import (
	"embed"
	"encoding/json"

	"github.com/nicksnyder/go-i18n/v2/i18n"
	"github.com/yyle88/neatjson/neatjsons"
	"github.com/yyle88/rese"
	"github.com/yyle88/zaplog"
	"go.uber.org/zap"
	"golang.org/x/text/language"
)

// DefaultLanguage 配置默认语言
var DefaultLanguage = language.AmericanEnglish

//go:embed trans.en-US.json trans.zh-CN.json
var files embed.FS

func LoadI18nFiles() (*i18n.Bundle, []*i18n.MessageFile) {
	bundle := i18n.NewBundle(DefaultLanguage)
	bundle.RegisterUnmarshalFunc("json", json.Unmarshal)

	var messageFiles []*i18n.MessageFile
	for _, name := range []string{"trans.en-US.json", "trans.zh-CN.json"} {
		zaplog.LOG.Debug("LOAD", zap.String("path", name))

		content := rese.A1(files.ReadFile(name))
		//这里文件名 file-name 写 "active.en-US.toml" 或者 "en-US.toml" 都行，内部会通过这个解析出语言标签名称
		messageFile := rese.P1(bundle.ParseMessageFileBytes(content, name))
		zaplog.SUG.Debugln(neatjsons.S(messageFile)) //安利下我的俩工具包

		messageFiles = append(messageFiles, messageFile)
	}

	zaplog.SUG.Debugln(neatjsons.S(bundle.LanguageTags()))
	return bundle, messageFiles
}
