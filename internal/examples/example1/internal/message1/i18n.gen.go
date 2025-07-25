package message1

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

func NewErrorAlreadyExist(data *ErrorAlreadyExistParam) (string, map[string]any) {
	return "ERROR_ALREADY_EXIST", data.GetTemplateValues()
}

func I18nErrorAlreadyExist(data *ErrorAlreadyExistParam) *i18n.LocalizeConfig {
	messageID, valuesMap := NewErrorAlreadyExist(data)
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

func NewErrorBadParam(data *ErrorBadParamParam) (string, map[string]any) {
	return "ERROR_BAD_PARAM", data.GetTemplateValues()
}

func I18nErrorBadParam(data *ErrorBadParamParam) *i18n.LocalizeConfig {
	messageID, valuesMap := NewErrorBadParam(data)
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

func NewErrorNotExist(data *ErrorNotExistParam) (string, map[string]any) {
	return "ERROR_NOT_EXIST", data.GetTemplateValues()
}

func I18nErrorNotExist(data *ErrorNotExistParam) *i18n.LocalizeConfig {
	messageID, valuesMap := NewErrorNotExist(data)
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

func NewNewMessages(data *NewMessagesParam) (string, map[string]any) {
	return "NEW_MESSAGES", data.GetTemplateValues()
}

func I18nNewMessages(data *NewMessagesParam) *i18n.LocalizeConfig {
	messageID, valuesMap := NewNewMessages(data)
	return &i18n.LocalizeConfig{
		MessageID:    messageID,
		TemplateData: valuesMap,
	}
}

func NewPleaseConfirm[Value comparable](value Value) (string, Value) {
	return "PLEASE_CONFIRM", value
}

func I18nPleaseConfirm[Value comparable](value Value) *i18n.LocalizeConfig {
	messageID, tempValue := NewPleaseConfirm(value)
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

func NewSayHello(data *SayHelloParam) (string, map[string]any) {
	return "SAY_HELLO", data.GetTemplateValues()
}

func I18nSayHello(data *SayHelloParam) *i18n.LocalizeConfig {
	messageID, valuesMap := NewSayHello(data)
	return &i18n.LocalizeConfig{
		MessageID:    messageID,
		TemplateData: valuesMap,
	}
}

func NewSuccess() string {
	return "SUCCESS"
}

func I18nSuccess() *i18n.LocalizeConfig {
	messageID := NewSuccess()
	return &i18n.LocalizeConfig{
		MessageID: messageID,
	}
}

func NewWelcome() string {
	return "WELCOME"
}

func I18nWelcome() *i18n.LocalizeConfig {
	messageID := NewWelcome()
	return &i18n.LocalizeConfig{
		MessageID: messageID,
	}
}
