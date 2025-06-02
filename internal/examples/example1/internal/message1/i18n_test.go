package message1_test

import (
	"testing"

	"github.com/yyle88/goi18n/internal/examples/example1/internal/message1"
	"github.com/yyle88/neatjson/neatjsons"
)

func TestLoadI18nFiles(t *testing.T) {
	bundle, messageFiles := message1.LoadI18nFiles()
	t.Log(len(messageFiles))
	t.Log(len(bundle.LanguageTags()))
}

func TestNewErrorAlreadyExist(t *testing.T) {
	messageID, templateValues := message1.NewErrorAlreadyExist(&message1.ErrorAlreadyExistParam{
		What: "abc",
		Code: "123",
	})
	t.Log(messageID)
	t.Log(neatjsons.S(templateValues))
}
