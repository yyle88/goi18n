package example1message_test

import (
	"testing"

	"github.com/yyle88/goi18n/internal/examples/example1/example1message"
	"github.com/yyle88/neatjson/neatjsons"
)

func TestLoadI18nFiles(t *testing.T) {
	bundle, messageFiles := example1message.LoadI18nFiles()
	t.Log(len(messageFiles))
	t.Log(len(bundle.LanguageTags()))
}

func TestNewErrorAlreadyExist(t *testing.T) {
	messageID, templateValues := example1message.NewErrorAlreadyExist(&example1message.ErrorAlreadyExistParam{
		What: "abc",
		Code: "123",
	})
	t.Log(messageID)
	t.Log(neatjsons.S(templateValues))
}
