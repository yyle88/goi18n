package example1message

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
	return &i18n.LocalizeConfig{
		MessageID:    "ERROR_ALREADY_EXIST",
		TemplateData: data.GetTemplateValues(),
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
	return &i18n.LocalizeConfig{
		MessageID:    "ERROR_BAD_PARAM",
		TemplateData: data.GetTemplateValues(),
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
	return &i18n.LocalizeConfig{
		MessageID:    "ERROR_NOT_EXIST",
		TemplateData: data.GetTemplateValues(),
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
	return &i18n.LocalizeConfig{
		MessageID:    "NEW_MESSAGES",
		TemplateData: data.GetTemplateValues(),
	}
}

func NewPleaseConfirm[Value comparable](value Value) (string, Value) {
	return "PLEASE_CONFIRM", value
}

func I18nPleaseConfirm[Value comparable](value Value) *i18n.LocalizeConfig {
	return &i18n.LocalizeConfig{
		MessageID:    "PLEASE_CONFIRM",
		TemplateData: value,
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
	return &i18n.LocalizeConfig{
		MessageID:    "SAY_HELLO",
		TemplateData: data.GetTemplateValues(),
	}
}

func NewSuccess() string {
	return "SUCCESS"
}

func I18nSuccess() *i18n.LocalizeConfig {
	return &i18n.LocalizeConfig{
		MessageID: "SUCCESS",
	}
}

func NewWelcome() string {
	return "WELCOME"
}

func I18nWelcome() *i18n.LocalizeConfig {
	return &i18n.LocalizeConfig{
		MessageID: "WELCOME",
	}
}
