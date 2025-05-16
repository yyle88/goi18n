package main

import (
	"fmt"

	"github.com/BurntSushi/toml"
	"github.com/nicksnyder/go-i18n/v2/i18n"
	"github.com/yyle88/must"
	"github.com/yyle88/neatjson/neatjsons"
	"github.com/yyle88/runpath"
	"github.com/yyle88/zaplog"
	"golang.org/x/text/language"
)

func main() {
	// 创建语言包
	bundle := i18n.NewBundle(language.English)
	bundle.RegisterUnmarshalFunc("toml", toml.Unmarshal)
	{
		messageFile, err := bundle.LoadMessageFile(runpath.PARENT.Join("active.en.toml")) // 加载英文翻译文件
		must.Done(err)
		zaplog.SUG.Debugln(neatjsons.S(messageFile))
	}
	{
		messageFile, err := bundle.LoadMessageFile(runpath.PARENT.Join("active.zh.toml")) // 加载中文翻译文件
		must.Done(err)
		zaplog.SUG.Debugln(neatjsons.S(messageFile))
	}

	zaplog.SUG.Debugln("---")

	{
		localizer := i18n.NewLocalizer(bundle, "en")
		// According to Unicode CLDR plural rules for English, the "zero" category is not defined.
		// When PluralCount is 0, go - i18n selects the "other" message ("I have 0 cats.") instead
		// of the "zero" message ("I have no cats"), as English only supports "one" (count=1) and
		// "other" (all other counts, including 0).
		for num := 0; num <= 3; num++ {
			fmt.Println(localizer.MustLocalize(&i18n.LocalizeConfig{
				MessageID:   "Cats",
				PluralCount: num, // this value is used to choose one/many template
				TemplateData: map[string]interface{}{
					"count": num, // this value is used to set value into the template
				},
			}))
		}
		//Output:
		//I have 0 cats.
		//I have one cat.
		//I have 2 cats.
		//I have 3 cats.
	}

	zaplog.SUG.Debugln("---")

	{
		localizer := i18n.NewLocalizer(bundle, "zh")
		// According to Unicode CLDR plural rules for Chinese, only the "other" category is defined.
		// When PluralCount is 0, 1, or any other value, go - i18n selects the "other" message
		// (e.g., "我有 0 只猫。") as Chinese does not distinguish "zero", "one", or other plural forms.
		for num := 0; num <= 3; num++ {
			fmt.Println(localizer.MustLocalize(&i18n.LocalizeConfig{
				MessageID:   "Cats",
				PluralCount: num, // this value is used to choose one/many template
				TemplateData: map[string]interface{}{
					"count": num, // this value is used to set value into the template
				},
			}))
		}
		//Output:
		//我有 0 只猫。
		//我有 1 只猫。
		//我有 2 只猫。
		//我有 3 只猫。
	}

	zaplog.SUG.Debugln("---")
}
