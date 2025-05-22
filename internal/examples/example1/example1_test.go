package example1_test

import (
	"testing"

	"github.com/nicksnyder/go-i18n/v2/i18n"
	"github.com/stretchr/testify/require"
	"github.com/yyle88/goi18n/internal/examples/example1/example1message"
)

var caseBundle *i18n.Bundle

func TestMain(m *testing.M) {
	caseBundle, _ = example1message.LoadI18nFiles()
	m.Run()
}

func TestI18nSayHello(t *testing.T) {
	t.Run("SAY_HELLO-zh", func(t *testing.T) {
		localizer := i18n.NewLocalizer(caseBundle, "zh-CN")

		msg, err := localizer.Localize(example1message.I18nSayHello(&example1message.SayHelloParam{
			Name: "杨亦乐",
		}))
		require.NoError(t, err)
		t.Log(msg)
		require.Equal(t, "你好，杨亦乐！", msg)
	})
	t.Run("SAY_HELLO-en", func(t *testing.T) {
		localizer := i18n.NewLocalizer(caseBundle, "en-US")

		msg, err := localizer.Localize(example1message.I18nSayHello(&example1message.SayHelloParam{
			Name: "yangyile",
		}))
		require.NoError(t, err)
		t.Log(msg)
		require.Equal(t, "Hello, yangyile!", msg)
	})
	t.Run("SAY_HELLO-km", func(t *testing.T) {
		localizer := i18n.NewLocalizer(caseBundle, "km-KH")

		msg, err := localizer.Localize(example1message.I18nSayHello(&example1message.SayHelloParam{
			Name: "yangyile",
		}))
		require.NoError(t, err)
		t.Log(msg)
		require.Equal(t, "សួស្តី yangyile!", msg) //完全看不懂高棉语啊
	})
}

func TestNewSayHello(t *testing.T) {
	t.Run("SAY_HELLO-zh", func(t *testing.T) {
		localizer := i18n.NewLocalizer(caseBundle, "zh-CN")

		messageID, msgValues := example1message.NewSayHello(&example1message.SayHelloParam{
			Name: "杨亦乐",
		})

		msg, err := localizer.Localize(&i18n.LocalizeConfig{
			MessageID:    messageID,
			TemplateData: msgValues,
		})
		require.NoError(t, err)
		t.Log(msg)
		require.Equal(t, "你好，杨亦乐！", msg)
	})
	t.Run("SAY_HELLO-en", func(t *testing.T) {
		localizer := i18n.NewLocalizer(caseBundle, "en-US")

		messageID, msgValues := example1message.NewSayHello(&example1message.SayHelloParam{
			Name: "yangyile",
		})

		msg, err := localizer.Localize(&i18n.LocalizeConfig{
			MessageID:    messageID,
			TemplateData: msgValues,
		})
		require.NoError(t, err)
		t.Log(msg)
		require.Equal(t, "Hello, yangyile!", msg)
	})
	t.Run("SAY_HELLO-km", func(t *testing.T) {
		localizer := i18n.NewLocalizer(caseBundle, "km-KH")

		messageID, msgValues := example1message.NewSayHello(&example1message.SayHelloParam{
			Name: "yangyile",
		})

		msg, err := localizer.Localize(&i18n.LocalizeConfig{
			MessageID:    messageID,
			TemplateData: msgValues,
		})
		require.NoError(t, err)
		t.Log(msg)
		require.Equal(t, "សួស្តី yangyile!", msg)
	})
}

func TestI18nWelcome(t *testing.T) {
	t.Run("WELCOME-zh", func(t *testing.T) {
		localizer := i18n.NewLocalizer(caseBundle, "zh-CN")

		msg, err := localizer.Localize(example1message.I18nWelcome())
		require.NoError(t, err)
		t.Log(msg)
		require.Equal(t, "欢迎使用此应用！", msg)
	})
}

func TestI18nSuccess(t *testing.T) {
	t.Run("SUCCESS-zh", func(t *testing.T) {
		localizer := i18n.NewLocalizer(caseBundle, "zh-CN")

		msg, err := localizer.Localize(example1message.I18nSuccess())
		require.NoError(t, err)
		t.Log(msg)
		require.Equal(t, "成功", msg)
	})
}

func TestI18nPleaseConfirm(t *testing.T) {
	t.Run("PLEASE_CONFIRM-zh", func(t *testing.T) {
		localizer := i18n.NewLocalizer(caseBundle, "zh-CN")

		msg, err := localizer.Localize(example1message.I18nPleaseConfirm("提交材料"))
		require.NoError(t, err)
		t.Log(msg)
		require.Equal(t, "请确认提交材料", msg)
	})
}

func TestI18nErrorNotExist(t *testing.T) {
	t.Run("ERROR_NOT_EXIST-zh", func(t *testing.T) {
		localizer := i18n.NewLocalizer(caseBundle, "zh-CN")

		msg, err := localizer.Localize(example1message.I18nErrorNotExist(&example1message.ErrorNotExistParam{
			What: "数据库里",
			Code: "账号信息",
		}))
		require.NoError(t, err)
		t.Log(msg)
		require.Equal(t, "数据库里 账号信息 不存在", msg)
	})
}

func TestI18nErrorAlreadyExist(t *testing.T) {
	t.Run("ERROR_ALREADY_EXIST-zh", func(t *testing.T) {
		localizer := i18n.NewLocalizer(caseBundle, "zh-CN")

		msg, err := localizer.Localize(example1message.I18nErrorAlreadyExist(&example1message.ErrorAlreadyExistParam{
			What: "系统里",
			Code: "玩家名",
		}))
		require.NoError(t, err)
		t.Log(msg)
		require.Equal(t, "系统里 玩家名 已存在", msg)
	})
}
