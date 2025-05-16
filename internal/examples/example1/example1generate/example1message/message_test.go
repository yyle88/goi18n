package example1message_test

import (
	"testing"

	"github.com/yyle88/goi18n/internal/examples/example1/example1generate/example1message"
	"github.com/yyle88/neatjson/neatjsons"
)

func TestNewErrorAlreadyExist(t *testing.T) {
	messageID, templateValues := example1message.NewErrorAlreadyExist(&example1message.ErrorAlreadyExistParam{
		What: "abc",
		Code: "123",
	})
	t.Log(messageID)
	t.Log(neatjsons.S(templateValues))
}
