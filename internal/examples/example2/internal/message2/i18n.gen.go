package message2

import (
	"github.com/nicksnyder/go-i18n/v2/i18n"
	"github.com/yyle88/goi18n"
)

type ActiveUsersParam struct {
	Count any
}

func (p *ActiveUsersParam) GetTemplateValues() map[string]any {
	res := make(map[string]any)
	if p.Count != nil {
		res["count"] = p.Count
	}
	return res
}

func NewActiveUsers(data *ActiveUsersParam) (string, map[string]any) {
	return "ACTIVE_USERS", data.GetTemplateValues()
}

func I18nActiveUsers(data *ActiveUsersParam, pluralConfig *goi18n.PluralConfig) *i18n.LocalizeConfig {
	return &i18n.LocalizeConfig{
		MessageID:    "ACTIVE_USERS",
		TemplateData: data.GetTemplateValues(),
		PluralCount:  pluralConfig.PluralCount,
	}
}

type CompletedTasksParam struct {
	Count any
}

func (p *CompletedTasksParam) GetTemplateValues() map[string]any {
	res := make(map[string]any)
	if p.Count != nil {
		res["count"] = p.Count
	}
	return res
}

func NewCompletedTasks(data *CompletedTasksParam) (string, map[string]any) {
	return "COMPLETED_TASKS", data.GetTemplateValues()
}

func I18nCompletedTasks(data *CompletedTasksParam, pluralConfig *goi18n.PluralConfig) *i18n.LocalizeConfig {
	return &i18n.LocalizeConfig{
		MessageID:    "COMPLETED_TASKS",
		TemplateData: data.GetTemplateValues(),
		PluralCount:  pluralConfig.PluralCount,
	}
}

type OpenIssuesParam struct {
	Count any
}

func (p *OpenIssuesParam) GetTemplateValues() map[string]any {
	res := make(map[string]any)
	if p.Count != nil {
		res["count"] = p.Count
	}
	return res
}

func NewOpenIssues(data *OpenIssuesParam) (string, map[string]any) {
	return "OPEN_ISSUES", data.GetTemplateValues()
}

func I18nOpenIssues(data *OpenIssuesParam, pluralConfig *goi18n.PluralConfig) *i18n.LocalizeConfig {
	return &i18n.LocalizeConfig{
		MessageID:    "OPEN_ISSUES",
		TemplateData: data.GetTemplateValues(),
		PluralCount:  pluralConfig.PluralCount,
	}
}

func NewPendingReviews[Value comparable](value Value) (string, Value) {
	return "PENDING_REVIEWS", value
}

func I18nPendingReviews[Value comparable](value Value, pluralConfig *goi18n.PluralConfig) *i18n.LocalizeConfig {
	return &i18n.LocalizeConfig{
		MessageID:    "PENDING_REVIEWS",
		TemplateData: value,
		PluralCount:  pluralConfig.PluralCount,
	}
}
