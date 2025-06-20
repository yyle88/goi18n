package message1v2

import (
	"github.com/nicksnyder/go-i18n/v2/i18n"
)

type ErrorAlreadyExistParam struct {
	What any
	Code any
}

func (p *ErrorAlreadyExistParam) GetTemplateValues() map[string]any {
	res := make(map[string]any)
	if p.What != nil {
		res["what"] = p.What
	}
	if p.Code != nil {
		res["code"] = p.Code
	}
	return res
}

func I18nErrorAlreadyExist(data *ErrorAlreadyExistParam) *i18n.LocalizeConfig {
	const messageID = "ERROR_ALREADY_EXIST"
	var valuesMap = data.GetTemplateValues()
	return &i18n.LocalizeConfig{
		MessageID:    messageID,
		TemplateData: valuesMap,
	}
}

type ErrorBadParamParam struct {
	Name any
}

func (p *ErrorBadParamParam) GetTemplateValues() map[string]any {
	res := make(map[string]any)
	if p.Name != nil {
		res["name"] = p.Name
	}
	return res
}

func I18nErrorBadParam(data *ErrorBadParamParam) *i18n.LocalizeConfig {
	const messageID = "ERROR_BAD_PARAM"
	var valuesMap = data.GetTemplateValues()
	return &i18n.LocalizeConfig{
		MessageID:    messageID,
		TemplateData: valuesMap,
	}
}

type ErrorNotExistParam struct {
	What any
	Code any
}

func (p *ErrorNotExistParam) GetTemplateValues() map[string]any {
	res := make(map[string]any)
	if p.What != nil {
		res["what"] = p.What
	}
	if p.Code != nil {
		res["code"] = p.Code
	}
	return res
}

func I18nErrorNotExist(data *ErrorNotExistParam) *i18n.LocalizeConfig {
	const messageID = "ERROR_NOT_EXIST"
	var valuesMap = data.GetTemplateValues()
	return &i18n.LocalizeConfig{
		MessageID:    messageID,
		TemplateData: valuesMap,
	}
}

type NewMessagesParam struct {
	Name  any
	Count any
}

func (p *NewMessagesParam) GetTemplateValues() map[string]any {
	res := make(map[string]any)
	if p.Name != nil {
		res["name"] = p.Name
	}
	if p.Count != nil {
		res["count"] = p.Count
	}
	return res
}

func I18nNewMessages(data *NewMessagesParam) *i18n.LocalizeConfig {
	const messageID = "NEW_MESSAGES"
	var valuesMap = data.GetTemplateValues()
	return &i18n.LocalizeConfig{
		MessageID:    messageID,
		TemplateData: valuesMap,
	}
}

func I18nPleaseConfirm[Value comparable](value Value) *i18n.LocalizeConfig {
	const messageID = "PLEASE_CONFIRM"
	var tempValue = value
	return &i18n.LocalizeConfig{
		MessageID:    messageID,
		TemplateData: tempValue,
	}
}

type SayHelloParam struct {
	Name any
}

func (p *SayHelloParam) GetTemplateValues() map[string]any {
	res := make(map[string]any)
	if p.Name != nil {
		res["name"] = p.Name
	}
	return res
}

func I18nSayHello(data *SayHelloParam) *i18n.LocalizeConfig {
	const messageID = "SAY_HELLO"
	var valuesMap = data.GetTemplateValues()
	return &i18n.LocalizeConfig{
		MessageID:    messageID,
		TemplateData: valuesMap,
	}
}

func I18nSuccess() *i18n.LocalizeConfig {
	const messageID = "SUCCESS"
	return &i18n.LocalizeConfig{
		MessageID: messageID,
	}
}

func I18nWelcome() *i18n.LocalizeConfig {
	const messageID = "WELCOME"
	return &i18n.LocalizeConfig{
		MessageID: messageID,
	}
}
