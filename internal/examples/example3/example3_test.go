package example3_test

import (
	"testing"

	"github.com/nicksnyder/go-i18n/v2/i18n"
	"github.com/stretchr/testify/require"
	"github.com/yyle88/goi18n/internal/examples/example3/internal/message3"
)

var caseBundle *i18n.Bundle

func TestMain(m *testing.M) {
	caseBundle, _ = message3.LoadI18nFiles()
	m.Run()
}

func TestI18nI早上好呀(t *testing.T) {
	t.Run("other-en", func(t *testing.T) {
		localizer := i18n.NewLocalizer(caseBundle, "en-US")
		msg, err := localizer.Localize(message3.I18nI早上好呀(&message3.P早上好呀{
			V某某某: localizer.MustLocalize(message3.I18nI老师()),
		}))
		require.NoError(t, err)
		t.Log(msg)
		require.Equal(t, "Good morning, Teacher", msg)
	})
	t.Run("other-zh", func(t *testing.T) {
		localizer := i18n.NewLocalizer(caseBundle, "zh-CN")
		msg, err := localizer.Localize(message3.I18nI早上好呀(&message3.P早上好呀{
			V某某某: localizer.MustLocalize(message3.I18nI老师()),
		}))
		require.NoError(t, err)
		t.Log(msg)
		require.Equal(t, "早上好呀老师", msg)
	})
}

func TestI18nI吃饭了没(t *testing.T) {
	t.Run("other-en", func(t *testing.T) {
		localizer := i18n.NewLocalizer(caseBundle, "en-US")
		msg, err := localizer.Localize(message3.I18nI吃饭了没())
		require.NoError(t, err)
		t.Log(msg)
		require.Equal(t, "Have you eaten?", msg)
	})
	t.Run("other-zh", func(t *testing.T) {
		localizer := i18n.NewLocalizer(caseBundle, "zh-CN")
		msg, err := localizer.Localize(message3.I18nI吃饭了没())
		require.NoError(t, err)
		t.Log(msg)
		require.Equal(t, "吃饭了没", msg)
	})
}

func TestI18nI我这里有个X你吃吧(t *testing.T) {
	t.Run("other-en", func(t *testing.T) {
		localizer := i18n.NewLocalizer(caseBundle, "en-US")
		msg, err := localizer.Localize(message3.I18nI我这里有个X你吃吧(
			&message3.P我这里有个X你吃吧{
				V叉叉叉: localizer.MustLocalize(message3.I18nI蛋糕()),
			},
		))
		require.NoError(t, err)
		t.Log(msg)
		require.Equal(t, "I have a Cake for you to eat.", msg)
	})
	t.Run("other-zh", func(t *testing.T) {
		localizer := i18n.NewLocalizer(caseBundle, "zh-CN")
		msg, err := localizer.Localize(message3.I18nI我这里有个X你吃吧(
			&message3.P我这里有个X你吃吧{
				V叉叉叉: localizer.MustLocalize(message3.I18nI蛋糕()),
			},
		))
		require.NoError(t, err)
		t.Log(msg)
		require.Equal(t, "我这里有个蛋糕你吃吧", msg)
	})
}

func TestI18nI祝X节日快乐(t *testing.T) {
	t.Run("other-en", func(t *testing.T) {
		localizer := i18n.NewLocalizer(caseBundle, "en-US")
		msg, err := localizer.Localize(message3.I18nI祝X节日快乐(
			localizer.MustLocalize(message3.I18nI老师()),
		))
		require.NoError(t, err)
		t.Log(msg)
		require.Equal(t, "Happy Teacher festival!", msg)
	})
	t.Run("other-zh", func(t *testing.T) {
		localizer := i18n.NewLocalizer(caseBundle, "zh-CN")
		msg, err := localizer.Localize(message3.I18nI祝X节日快乐(
			localizer.MustLocalize(message3.I18nI老师()),
		))
		require.NoError(t, err)
		t.Log(msg)
		require.Equal(t, "祝老师们节日快乐", msg)
	})
}

func TestI18nI祝X某X节快乐(t *testing.T) {
	t.Run("other-en", func(t *testing.T) {
		localizer := i18n.NewLocalizer(caseBundle, "en-US")
		msg, err := localizer.Localize(message3.I18nI祝X某X节快乐(
			&message3.P祝X某X节快乐{
				V某某人: localizer.MustLocalize(message3.I18nI同学()),
				V某某节: localizer.MustLocalize(message3.I18nI春节()),
			},
		))
		require.NoError(t, err)
		t.Log(msg)
		require.Equal(t, "Happy Classmate Spring Festival!", msg)
	})
	t.Run("other-zh", func(t *testing.T) {
		localizer := i18n.NewLocalizer(caseBundle, "zh-CN")
		msg, err := localizer.Localize(message3.I18nI祝X某X节快乐(
			&message3.P祝X某X节快乐{
				V某某人: localizer.MustLocalize(message3.I18nI同学()),
				V某某节: localizer.MustLocalize(message3.I18nI春节()),
			},
		))
		require.NoError(t, err)
		t.Log(msg)
		require.Equal(t, "祝同学们春节快乐", msg)
	})
}
