package message3v2

import (
	"github.com/nicksnyder/go-i18n/v2/i18n"
)

func I18nI元旦() *i18n.LocalizeConfig {
	const messageID = "元旦"
	return &i18n.LocalizeConfig{
		MessageID: messageID,
	}
}

func I18nI吃饭了没() *i18n.LocalizeConfig {
	const messageID = "吃饭了没"
	return &i18n.LocalizeConfig{
		MessageID: messageID,
	}
}

func I18nI同学() *i18n.LocalizeConfig {
	const messageID = "同学"
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

func I18nI我这里有个X你吃吧(data *P我这里有个X你吃吧) *i18n.LocalizeConfig {
	const messageID = "我这里有个X你吃吧"
	var valuesMap = data.GetTemplateValues()
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

func I18nI早上好呀(data *P早上好呀) *i18n.LocalizeConfig {
	const messageID = "早上好呀"
	var valuesMap = data.GetTemplateValues()
	return &i18n.LocalizeConfig{
		MessageID:    messageID,
		TemplateData: valuesMap,
	}
}

func I18nI春节() *i18n.LocalizeConfig {
	const messageID = "春节"
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

func I18nI祝X某X节快乐(data *P祝X某X节快乐) *i18n.LocalizeConfig {
	const messageID = "祝X某X节快乐"
	var valuesMap = data.GetTemplateValues()
	return &i18n.LocalizeConfig{
		MessageID:    messageID,
		TemplateData: valuesMap,
	}
}

func I18nI祝X节日快乐[Value comparable](value Value) *i18n.LocalizeConfig {
	const messageID = "祝X节日快乐"
	var tempValue = value
	return &i18n.LocalizeConfig{
		MessageID:    messageID,
		TemplateData: tempValue,
	}
}

func I18nI老师() *i18n.LocalizeConfig {
	const messageID = "老师"
	return &i18n.LocalizeConfig{
		MessageID: messageID,
	}
}

func I18nI蛋糕() *i18n.LocalizeConfig {
	const messageID = "蛋糕"
	return &i18n.LocalizeConfig{
		MessageID: messageID,
	}
}

func I18nI面包() *i18n.LocalizeConfig {
	const messageID = "面包"
	return &i18n.LocalizeConfig{
		MessageID: messageID,
	}
}
