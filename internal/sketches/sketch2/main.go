package main

import (
	"fmt"

	"github.com/nicksnyder/go-i18n/v2/i18n"
	"golang.org/x/text/language"
)

func main() {
	bundle := i18n.NewBundle(language.English)
	localizer := i18n.NewLocalizer(bundle, "en")
	catsMessage := &i18n.Message{
		ID:    "Cats",
		Zero:  "I have no cats", // useless: English only supports "one" (count=1) and "other" (all other counts, including 0).
		One:   "I have one cat.",
		Two:   "I have two cats.", // useless: English only supports "one" (count=1) and "other" (all other counts, including 0).
		Other: "I have {{.count}} cats.",
	}
	// According to Unicode CLDR plural rules for English, the "zero" category is not defined.
	// When PluralCount is 0, go - i18n selects the "other" message ("I have 0 cats.") instead
	// of the "zero" message ("I have no cats"), as English only supports "one" (count=1) and
	// "other" (all other counts, including 0).
	for num := 0; num <= 3; num++ {
		fmt.Println(localizer.MustLocalize(&i18n.LocalizeConfig{
			DefaultMessage: catsMessage,
			PluralCount:    num, // this value is used to choose one/many template
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
