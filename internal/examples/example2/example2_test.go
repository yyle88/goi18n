package example2_test

import (
	"testing"

	"github.com/nicksnyder/go-i18n/v2/i18n"
	"github.com/stretchr/testify/require"
	"github.com/yyle88/goi18n"
	"github.com/yyle88/goi18n/internal/examples/example2/example2message"
)

var caseBundle *i18n.Bundle

func TestMain(m *testing.M) {
	caseBundle, _ = example2message.LoadI18nFiles()
	m.Run()
}

func TestI18nActiveUsers(t *testing.T) {
	t.Run("one-1-en", func(t *testing.T) {
		localizer := i18n.NewLocalizer(caseBundle, "en-US")
		const count = 1
		msg, err := localizer.Localize(example2message.I18nActiveUsers(
			&example2message.ActiveUsersParam{
				Count: count,
			},
			goi18n.NewPluralConfig(count),
		))
		require.NoError(t, err)
		t.Log(msg)
		require.Equal(t, "One user is active in the project.", msg)
	})
	t.Run("other-en", func(t *testing.T) {
		localizer := i18n.NewLocalizer(caseBundle, "en-US")
		const count = 9999
		msg, err := localizer.Localize(example2message.I18nActiveUsers(
			&example2message.ActiveUsersParam{
				Count: count,
			},
			goi18n.NewPluralConfig(count),
		))
		require.NoError(t, err)
		t.Log(msg)
		require.Equal(t, "9999 users are active in the project.", msg)
	})
	t.Run("other-zh", func(t *testing.T) {
		localizer := i18n.NewLocalizer(caseBundle, "zh-CN")
		const count = 9999
		msg, err := localizer.Localize(example2message.I18nActiveUsers(
			&example2message.ActiveUsersParam{
				Count: count,
			},
			goi18n.NewPluralConfig(count),
		))
		require.NoError(t, err)
		t.Log(msg)
		require.Equal(t, "项目中有 9999 个活跃用户。", msg)
	})
}

func TestI18nCompletedTasks(t *testing.T) {
	t.Run("one-1-en", func(t *testing.T) {
		localizer := i18n.NewLocalizer(caseBundle, "en-US")
		const count = 1
		msg, err := localizer.Localize(example2message.I18nCompletedTasks(
			&example2message.CompletedTasksParam{
				Count: count,
			},
			goi18n.NewPluralConfig(count),
		))
		require.NoError(t, err)
		t.Log(msg)
		require.Equal(t, "One task is completed in the project.", msg)
	})
	t.Run("other-en", func(t *testing.T) {
		localizer := i18n.NewLocalizer(caseBundle, "en-US")
		const count = 9999
		msg, err := localizer.Localize(example2message.I18nCompletedTasks(
			&example2message.CompletedTasksParam{
				Count: count,
			},
			goi18n.NewPluralConfig(count),
		))
		require.NoError(t, err)
		t.Log(msg)
		require.Equal(t, "9999 tasks are completed in the project.", msg)
	})
	t.Run("other-zh", func(t *testing.T) {
		localizer := i18n.NewLocalizer(caseBundle, "zh-CN")
		const count = 9999
		msg, err := localizer.Localize(example2message.I18nCompletedTasks(
			&example2message.CompletedTasksParam{
				Count: count,
			},
			goi18n.NewPluralConfig(count),
		))
		require.NoError(t, err)
		t.Log(msg)
		require.Equal(t, "项目中完成了 9999 个任务。", msg)
	})
}

func TestI18nOpenIssues(t *testing.T) {
	t.Run("one-1-en", func(t *testing.T) {
		localizer := i18n.NewLocalizer(caseBundle, "en-US")
		const count = 1
		msg, err := localizer.Localize(example2message.I18nOpenIssues(
			&example2message.OpenIssuesParam{
				Count: count,
			},
			goi18n.NewPluralConfig(count),
		))
		require.NoError(t, err)
		t.Log(msg)
		require.Equal(t, "There is one open issue in the project.", msg)
	})
	t.Run("other-en", func(t *testing.T) {
		localizer := i18n.NewLocalizer(caseBundle, "en-US")
		const count = 3
		msg, err := localizer.Localize(example2message.I18nOpenIssues(
			&example2message.OpenIssuesParam{
				Count: count,
			},
			goi18n.NewPluralConfig(count),
		))
		require.NoError(t, err)
		t.Log(msg)
		require.Equal(t, "There are 3 open issues in the project.", msg)
	})
	t.Run("other-zh", func(t *testing.T) {
		localizer := i18n.NewLocalizer(caseBundle, "zh-CN")
		const count = 3
		msg, err := localizer.Localize(example2message.I18nOpenIssues(
			&example2message.OpenIssuesParam{
				Count: count,
			},
			goi18n.NewPluralConfig(count),
		))
		require.NoError(t, err)
		t.Log(msg)
		require.Equal(t, "项目中有 3 个未解决问题。", msg)
	})
}

func TestI18nPendingReviews(t *testing.T) {
	t.Run("one-1-en", func(t *testing.T) {
		localizer := i18n.NewLocalizer(caseBundle, "en-US")
		const count = 1
		msg, err := localizer.Localize(example2message.I18nPendingReviews(
			count,
			goi18n.NewPluralConfig(count),
		))
		require.NoError(t, err)
		t.Log(msg)
		require.Equal(t, "There is one pending review in the project.", msg)
	})
	t.Run("other-en", func(t *testing.T) {
		localizer := i18n.NewLocalizer(caseBundle, "en-US")
		const count = 3
		msg, err := localizer.Localize(example2message.I18nPendingReviews(
			count,
			goi18n.NewPluralConfig(count),
		))
		require.NoError(t, err)
		t.Log(msg)
		require.Equal(t, "There are 3 pending reviews in the project.", msg)
	})
	t.Run("other-zh", func(t *testing.T) {
		localizer := i18n.NewLocalizer(caseBundle, "zh-CN")
		const count = 3
		msg, err := localizer.Localize(example2message.I18nPendingReviews(
			count,
			goi18n.NewPluralConfig(count),
		))
		require.NoError(t, err)
		t.Log(msg)
		require.Equal(t, "项目中有 3 个待审。", msg)
	})
}
