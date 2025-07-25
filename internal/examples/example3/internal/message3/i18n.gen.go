package message3

import (
	"github.com/nicksnyder/go-i18n/v2/i18n"
)

func NewI元旦() string {
	return "元旦"
}

func I18nI元旦() *i18n.LocalizeConfig {
	messageID := NewI元旦()
	return &i18n.LocalizeConfig{
		MessageID: messageID,
	}
}

func NewI吃饭了没() string {
	return "吃饭了没"
}

func I18nI吃饭了没() *i18n.LocalizeConfig {
	messageID := NewI吃饭了没()
	return &i18n.LocalizeConfig{
		MessageID: messageID,
	}
}

func NewI同学() string {
	return "同学"
}

func I18nI同学() *i18n.LocalizeConfig {
	messageID := NewI同学()
	return &i18n.LocalizeConfig{
		MessageID: messageID,
	}
}

type P我这里有个X你吃吧 struct {
	V叉叉叉 any
}

func (p *P我这里有个X你吃吧) GetTemplateValues() map[string]any {
	res := make(map[string]any)
	if p.V叉叉叉 != nil {
		res["叉叉叉"] = p.V叉叉叉
	}
	return res
}

func NewI我这里有个X你吃吧(data *P我这里有个X你吃吧) (string, map[string]any) {
	return "我这里有个X你吃吧", data.GetTemplateValues()
}

func I18nI我这里有个X你吃吧(data *P我这里有个X你吃吧) *i18n.LocalizeConfig {
	messageID, valuesMap := NewI我这里有个X你吃吧(data)
	return &i18n.LocalizeConfig{
		MessageID:    messageID,
		TemplateData: valuesMap,
	}
}

type P早上好呀 struct {
	V某某某 any
}

func (p *P早上好呀) GetTemplateValues() map[string]any {
	res := make(map[string]any)
	if p.V某某某 != nil {
		res["某某某"] = p.V某某某
	}
	return res
}

func NewI早上好呀(data *P早上好呀) (string, map[string]any) {
	return "早上好呀", data.GetTemplateValues()
}

func I18nI早上好呀(data *P早上好呀) *i18n.LocalizeConfig {
	messageID, valuesMap := NewI早上好呀(data)
	return &i18n.LocalizeConfig{
		MessageID:    messageID,
		TemplateData: valuesMap,
	}
}

func NewI春节() string {
	return "春节"
}

func I18nI春节() *i18n.LocalizeConfig {
	messageID := NewI春节()
	return &i18n.LocalizeConfig{
		MessageID: messageID,
	}
}

type P祝X某X节快乐 struct {
	V某某人 any
	V某某节 any
}

func (p *P祝X某X节快乐) GetTemplateValues() map[string]any {
	res := make(map[string]any)
	if p.V某某人 != nil {
		res["某某人"] = p.V某某人
	}
	if p.V某某节 != nil {
		res["某某节"] = p.V某某节
	}
	return res
}

func NewI祝X某X节快乐(data *P祝X某X节快乐) (string, map[string]any) {
	return "祝X某X节快乐", data.GetTemplateValues()
}

func I18nI祝X某X节快乐(data *P祝X某X节快乐) *i18n.LocalizeConfig {
	messageID, valuesMap := NewI祝X某X节快乐(data)
	return &i18n.LocalizeConfig{
		MessageID:    messageID,
		TemplateData: valuesMap,
	}
}

func NewI祝X节日快乐[Value comparable](value Value) (string, Value) {
	return "祝X节日快乐", value
}

func I18nI祝X节日快乐[Value comparable](value Value) *i18n.LocalizeConfig {
	messageID, tempValue := NewI祝X节日快乐(value)
	return &i18n.LocalizeConfig{
		MessageID:    messageID,
		TemplateData: tempValue,
	}
}

func NewI老师() string {
	return "老师"
}

func I18nI老师() *i18n.LocalizeConfig {
	messageID := NewI老师()
	return &i18n.LocalizeConfig{
		MessageID: messageID,
	}
}

func NewI蛋糕() string {
	return "蛋糕"
}

func I18nI蛋糕() *i18n.LocalizeConfig {
	messageID := NewI蛋糕()
	return &i18n.LocalizeConfig{
		MessageID: messageID,
	}
}

func NewI面包() string {
	return "面包"
}

func I18nI面包() *i18n.LocalizeConfig {
	messageID := NewI面包()
	return &i18n.LocalizeConfig{
		MessageID: messageID,
	}
}
