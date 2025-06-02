package message3_test

import (
	"testing"

	"github.com/yyle88/goi18n/internal/examples/example3/internal/message3"
	"github.com/yyle88/neatjson/neatjsons"
)

func TestLoadI18nFiles(t *testing.T) {
	bundle, messageFiles := message3.LoadI18nFiles()
	t.Log(len(messageFiles))
	t.Log(len(bundle.LanguageTags()))
}

func TestNewI祝X某X节快乐(t *testing.T) {
	messageID, templateValues := message3.NewI祝X某X节快乐(&message3.P祝X某X节快乐{
		V某某人: "乐乐",
		V某某节: "天天",
	})
	t.Log(messageID)
	t.Log(neatjsons.S(templateValues))
}
