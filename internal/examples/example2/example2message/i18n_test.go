package example2message_test

import (
	"testing"

	"github.com/yyle88/goi18n/internal/examples/example2/example2message"
	"github.com/yyle88/neatjson/neatjsons"
)

func TestLoadI18nFiles(t *testing.T) {
	bundle, messageFiles := example2message.LoadI18nFiles()
	t.Log(len(messageFiles))
	t.Log(len(bundle.LanguageTags()))
}

func TestNewActiveUsers(t *testing.T) {
	messageID, templateValues := example2message.NewActiveUsers(&example2message.ActiveUsersParam{
		Count: 8888,
	})
	t.Log(messageID)
	t.Log(neatjsons.S(templateValues))
}
