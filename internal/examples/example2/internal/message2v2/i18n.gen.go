package message2v2

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

func I18nActiveUsers(data *ActiveUsersParam, pluralConfig *goi18n.PluralConfig) *i18n.LocalizeConfig {
	const messageID = "ACTIVE_USERS"
	var valuesMap = data.GetTemplateValues()
	return &i18n.LocalizeConfig{
		MessageID:    messageID,
		TemplateData: valuesMap,
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

func I18nCompletedTasks(data *CompletedTasksParam, pluralConfig *goi18n.PluralConfig) *i18n.LocalizeConfig {
	const messageID = "COMPLETED_TASKS"
	var valuesMap = data.GetTemplateValues()
	return &i18n.LocalizeConfig{
		MessageID:    messageID,
		TemplateData: valuesMap,
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

func I18nOpenIssues(data *OpenIssuesParam, pluralConfig *goi18n.PluralConfig) *i18n.LocalizeConfig {
	const messageID = "OPEN_ISSUES"
	var valuesMap = data.GetTemplateValues()
	return &i18n.LocalizeConfig{
		MessageID:    messageID,
		TemplateData: valuesMap,
		PluralCount:  pluralConfig.PluralCount,
	}
}

func I18nPendingReviews[Value comparable](value Value, pluralConfig *goi18n.PluralConfig) *i18n.LocalizeConfig {
	const messageID = "PENDING_REVIEWS"
	var tempValue = value
	return &i18n.LocalizeConfig{
		MessageID:    messageID,
		TemplateData: tempValue,
		PluralCount:  pluralConfig.PluralCount,
	}
}
