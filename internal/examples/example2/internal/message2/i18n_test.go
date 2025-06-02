package message2_test

import (
	"testing"

	"github.com/yyle88/goi18n/internal/examples/example2/internal/message2"
	"github.com/yyle88/neatjson/neatjsons"
)

func TestLoadI18nFiles(t *testing.T) {
	bundle, messageFiles := message2.LoadI18nFiles()
	t.Log(len(messageFiles))
	t.Log(len(bundle.LanguageTags()))
}

func TestNewActiveUsers(t *testing.T) {
	messageID, templateValues := message2.NewActiveUsers(&message2.ActiveUsersParam{
		Count: 8888,
	})
	t.Log(messageID)
	t.Log(neatjsons.S(templateValues))
}
