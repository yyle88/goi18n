package example2message

import "github.com/nicksnyder/go-i18n/v2/i18n"

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

func I18nActiveUsers(data *ActiveUsersParam) *i18n.LocalizeConfig {
	return &i18n.LocalizeConfig{
		MessageID:    "ACTIVE_USERS",
		TemplateData: data.GetTemplateValues(),
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

func I18nCompletedTasks(data *CompletedTasksParam) *i18n.LocalizeConfig {
	return &i18n.LocalizeConfig{
		MessageID:    "COMPLETED_TASKS",
		TemplateData: data.GetTemplateValues(),
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

func I18nOpenIssues(data *OpenIssuesParam) *i18n.LocalizeConfig {
	return &i18n.LocalizeConfig{
		MessageID:    "OPEN_ISSUES",
		TemplateData: data.GetTemplateValues(),
	}
}

type PendingReviewsParam struct {
	Count any
}

func (p *PendingReviewsParam) GetTemplateValues() map[string]any {
	res := make(map[string]any)
	if p.Count != nil {
		res["count"] = p.Count
	}
	return res
}

func NewPendingReviews(data *PendingReviewsParam) (string, map[string]any) {
	return "PENDING_REVIEWS", data.GetTemplateValues()
}

func I18nPendingReviews(data *PendingReviewsParam) *i18n.LocalizeConfig {
	return &i18n.LocalizeConfig{
		MessageID:    "PENDING_REVIEWS",
		TemplateData: data.GetTemplateValues(),
	}
}
