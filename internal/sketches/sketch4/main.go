package main

import (
	"fmt"

	"github.com/BurntSushi/toml"
	"github.com/nicksnyder/go-i18n/v2/i18n"
	"github.com/yyle88/neatjson/neatjsons"
	"github.com/yyle88/rese"
	"github.com/yyle88/runpath"
	"github.com/yyle88/zaplog"
	"golang.org/x/text/language"
)

func main() {
	// 创建语言包
	bundle := i18n.NewBundle(language.English)
	bundle.RegisterUnmarshalFunc("toml", toml.Unmarshal)
	{
		messageFile := rese.P1(bundle.LoadMessageFile(runpath.PARENT.Join("active.en.toml"))) // 加载英文翻译文件
		zaplog.SUG.Debugln(neatjsons.S(messageFile))
	}
	{
		messageFile := rese.P1(bundle.LoadMessageFile(runpath.PARENT.Join("active.zh.toml"))) // 加载中文翻译文件
		zaplog.SUG.Debugln(neatjsons.S(messageFile))
	}

	zaplog.SUG.Debugln("---")

	{
		localizers := []*i18n.Localizer{
			i18n.NewLocalizer(bundle, "en"),
			i18n.NewLocalizer(bundle, "zh"),
		}
		for num := 0; num <= 3; num++ {
			for _, localizer := range localizers {
				fmt.Println(localizer.MustLocalize(&i18n.LocalizeConfig{
					MessageID:   "Cats",
					PluralCount: num,
					TemplateData: map[string]interface{}{
						"count": num,
					},
				}))
				fmt.Println(localizer.MustLocalize(&i18n.LocalizeConfig{
					MessageID:   "我有几只猫",
					PluralCount: num,
					TemplateData: map[string]interface{}{
						"猫的数量": num,
					},
				}))
			}
		}
	}

	zaplog.SUG.Debugln("---")
}
