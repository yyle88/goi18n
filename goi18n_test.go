package goi18n_test

import (
	"encoding/json"
	"testing"

	"github.com/nicksnyder/go-i18n/v2/i18n"
	"github.com/yyle88/goi18n"
	"github.com/yyle88/neatjson/neatjsons"
	"github.com/yyle88/rese"
	"github.com/yyle88/zaplog"
	"golang.org/x/text/language"
	"gopkg.in/yaml.v3"
)

// TestGenerate tests basic code generation with mixed message patterns
// Includes messages with named params, without params, and anonymous params
// Uses both English and Chinese message files with YAML format
//
// TestGenerate 测试包含多种消息模式的基本代码生成
// 包括带命名参数、无参数和匿名参数的消息
// 使用 YAML 格式的英文和中文消息文件
func TestGenerate(t *testing.T) {
	// Create i18n bundle with American English as base language
	// 创建以美式英语为基础语言的 i18n 包
	bundle := i18n.NewBundle(language.AmericanEnglish)
	bundle.RegisterUnmarshalFunc("yaml", yaml.Unmarshal)

	var messageFiles []*i18n.MessageFile
	{
		// Parse English message file with various message patterns
		// 解析包含各种消息模式的英文消息文件
		messageFile := rese.P1(bundle.ParseMessageFileBytes([]byte(`
SAY_HELLO: "Hello, {{ .name }}!"
WELCOME: "Welcome to this app!"
NEW_MESSAGES: "Hi, {{ .name }}! You have {{ .count }} new messages."
SUCCESS: "Success"
ERROR_BAD_PARAM: "Invalid parameter: {{ .name }}"
ERROR_ALREADY_EXIST: "{{ .what }} {{ .code }} already exists"
ERROR_NOT_EXIST: "{{ .what }} {{ .code }} does not exist"
PLEASE_CONFIRM: "Please confirm to {{ . }}"
`), "en-US.yaml"))
		zaplog.SUG.Debugln(neatjsons.S(messageFile))
		messageFiles = append(messageFiles, messageFile)
	}
	{
		// Parse Chinese message file with same message IDs as English version
		// 解析与英文版本具有相同消息 ID 的中文消息文件
		messageFile := rese.P1(bundle.ParseMessageFileBytes([]byte(`
SAY_HELLO: "你好，{{.name}}！"
WELCOME: "欢迎使用此应用！"
NEW_MESSAGES: "嗨，{{.name}}！您有 {{.count}} 条新消息。"
SUCCESS: "成功"
ERROR_BAD_PARAM: "无效参数：{{.name}}"
ERROR_ALREADY_EXIST: "{{.what}} {{.code}} 已存在"
ERROR_NOT_EXIST: "{{.what}} {{.code}} 不存在"
PLEASE_CONFIRM: "请确认{{.}}"
`), "zh-CN.yaml"))
		zaplog.SUG.Debugln(neatjsons.S(messageFile))
		messageFiles = append(messageFiles, messageFile)
	}
	zaplog.SUG.Debugln(neatjsons.S(bundle.LanguageTags()))

	// Generate type-safe Go code from message files
	// 从消息文件生成类型安全的 Go 代码
	goi18n.Generate(messageFiles, goi18n.NewOptions())
}

// TestGenerate_SingleValueYaml tests code generation from YAML files with simple named params
// Messages use single template variables instead of complex multi-param patterns
// Validates generation of param structs and support functions
//
// TestGenerate_SingleValueYaml 测试从 YAML 文件生成简单命名参数的代码
// 消息使用单个模板变量而非复杂的多参数模式
// 验证参数结构体和支持函数的生成
func TestGenerate_SingleValueYaml(t *testing.T) {
	bundle := i18n.NewBundle(language.AmericanEnglish)
	bundle.RegisterUnmarshalFunc("yaml", yaml.Unmarshal)

	var messageFiles []*i18n.MessageFile
	{
		messageFile := rese.P1(bundle.ParseMessageFileBytes([]byte(`
ORDER_PLACED: "Your order for {{ .item }} has been placed."
THANK_YOU: "Thank you for shopping with us, {{ .name }}!"
TOTAL_AMOUNT: "Your total is {{ .total }}."
STOCK_AVAILABLE: "We have {{ .quantity }} of {{ .item }} in stock."
OUT_OF_STOCK: "{{ .item }} is currently out of stock."
SHIPPING_INFO: "Your order will be shipped to {{ .address }}."
PAYMENT_SUCCESS: "Payment of {{ .amount }} has been successfully processed."
INVALID_PRODUCT: "Product {{ .productId }} does not exist."
LOW_BALANCE: "Your balance is low. Please add funds to continue shopping."
`), "en-US.yaml"))
		zaplog.SUG.Debugln(neatjsons.S(messageFile))
		messageFiles = append(messageFiles, messageFile)
	}
	{
		messageFile := rese.P1(bundle.ParseMessageFileBytes([]byte(`
ORDER_PLACED: "您的{{ .item }}订单已提交。"
THANK_YOU: "感谢您在我们这里购物，{{ .name }}！"
TOTAL_AMOUNT: "您的总计是{{ .total }}。"
STOCK_AVAILABLE: "我们有{{ .quantity }}件{{ .item }}库存。"
OUT_OF_STOCK: "{{ .item }}目前缺货。"
SHIPPING_INFO: "您的订单将寄送到{{ .address }}。"
PAYMENT_SUCCESS: "已成功处理{{ .amount }}的付款。"
INVALID_PRODUCT: "产品{{ .productId }}不存在。"
LOW_BALANCE: "您的余额不足。请添加资金以继续购物。"
`), "zh-CN.yaml"))
		zaplog.SUG.Debugln(neatjsons.S(messageFile))
		messageFiles = append(messageFiles, messageFile)
	}
	zaplog.SUG.Debugln(neatjsons.S(bundle.LanguageTags()))

	// Generate type-safe Go code from message files
	// 从消息文件生成类型安全的 Go 代码
	goi18n.Generate(messageFiles, goi18n.NewOptions())
}

// TestGenerate_PluralYaml tests code generation with CLDR plural forms from YAML files
// Messages define both "one" and "other" forms to handle singular and plural cases
// Demonstrates how plural messages generate functions that accept PluralCount param
//
// TestGenerate_PluralYaml 测试从 YAML 文件生成 CLDR 复数形式的代码
// 消息定义"one"和"other"形式以处理单数和复数情况
// 演示复数消息如何生成接受 PluralCount 参数的函数
func TestGenerate_PluralYaml(t *testing.T) {
	bundle := i18n.NewBundle(language.AmericanEnglish)
	bundle.RegisterUnmarshalFunc("yaml", yaml.Unmarshal)

	var messageFiles []*i18n.MessageFile
	{
		messageFile := rese.P1(bundle.ParseMessageFileBytes([]byte(`
NEW_NOTIFICATIONS:
  one: "You have one new notification."
  other: "You have {{.Count}} new notifications."
NEW_COMMENTS:
  one: "Your post received one new comment."
  other: "Your post received {{.Count}} new comments."
NEW_LIKES:
  one: "Your post has one new like."
  other: "Your post has {{.Count}} new likes."
NEW_MESSAGES:
  one: "You have one new message."
  other: "You have {{.Count}} new messages."
`), "en-US.yaml"))
		zaplog.SUG.Debugln(neatjsons.S(messageFile))
		messageFiles = append(messageFiles, messageFile)
	}
	{
		messageFile := rese.P1(bundle.ParseMessageFileBytes([]byte(`
NEW_NOTIFICATIONS:
  one: "您有一个新通知。"
  other: "您有 {{.Count}} 个新通知。"
NEW_COMMENTS:
  one: "您的帖子收到一个新评论。"
  other: "您的帖子收到 {{.Count}} 个新评论。"
NEW_LIKES:
  one: "您的帖子有一个新点赞。"
  other: "您的帖子有 {{.Count}} 个新点赞。"
NEW_MESSAGES:
  one: "您有一条新消息。"
  other: "您有 {{.Count}} 条新消息。"
`), "zh-CN.yaml"))
		zaplog.SUG.Debugln(neatjsons.S(messageFile))
		messageFiles = append(messageFiles, messageFile)
	}
	zaplog.SUG.Debugln(neatjsons.S(bundle.LanguageTags()))

	// Generate type-safe Go code with plural support from message files
	// 从消息文件生成支持复数的类型安全 Go 代码
	goi18n.Generate(messageFiles, goi18n.NewOptions())
}

// TestGenerate_SingleValueJson tests code generation from JSON format message files
// Uses JSON instead of YAML format to demonstrate format-agnostic support
// Validates param struct generation works consistently across formats
//
// TestGenerate_SingleValueJson 测试从 JSON 格式消息文件生成代码
// 使用 JSON 而非 YAML 格式以演示格式无关的支持
// 验证参数结构体生成在不同格式下都能正常工作
func TestGenerate_SingleValueJson(t *testing.T) {
	bundle := i18n.NewBundle(language.AmericanEnglish)
	bundle.RegisterUnmarshalFunc("json", json.Unmarshal)

	var messageFiles []*i18n.MessageFile
	{
		messageFile := rese.P1(bundle.ParseMessageFileBytes([]byte(`{
			"TASK_ASSIGNED": "Task {{.task}} has been assigned to {{.user}}.",
			"TASK_COMPLETED": "Task {{.task}} is completed.",
			"DEADLINE_REMINDER": "Reminder: Task {{.task}} is due on {{.date}}.",
			"PRIORITY_HIGH": "Task {{.task}} is marked as high priority."
		}`), "en-US.json"))
		zaplog.SUG.Debugln(neatjsons.S(messageFile))
		messageFiles = append(messageFiles, messageFile)
	}
	{
		messageFile := rese.P1(bundle.ParseMessageFileBytes([]byte(`{
			"TASK_ASSIGNED": "任务 {{.task}} 已分配给 {{.user}}。",
			"TASK_COMPLETED": "任务 {{.task}} 已完成。",
			"DEADLINE_REMINDER": "提醒：任务 {{.task}} 将于 {{.date}} 到期。",
			"PRIORITY_HIGH": "任务 {{.task}} 被标记为高优先级。"
		}`), "zh-CN.json"))
		zaplog.SUG.Debugln(neatjsons.S(messageFile))
		messageFiles = append(messageFiles, messageFile)
	}
	zaplog.SUG.Debugln(neatjsons.S(bundle.LanguageTags()))

	// Generate type-safe Go code from JSON format message files
	// 从 JSON 格式消息文件生成类型安全的 Go 代码
	goi18n.Generate(messageFiles, goi18n.NewOptions())
}

// TestGenerate_PluralJson tests code generation with plural forms from JSON format files
// Combines JSON format with CLDR plural rules to handle one and other cases
// Ensures plural support works consistently in JSON as it does in YAML
//
// TestGenerate_PluralJson 测试从 JSON 格式文件生成复数形式的代码
// 将 JSON 格式与 CLDR 复数规则结合以处理 one 和 other 情况
// 确保 JSON 中的复数支持与 YAML 中的工作方式一致
func TestGenerate_PluralJson(t *testing.T) {
	bundle := i18n.NewBundle(language.AmericanEnglish)
	bundle.RegisterUnmarshalFunc("json", json.Unmarshal)

	var messageFiles []*i18n.MessageFile
	{
		messageFile := rese.P1(bundle.ParseMessageFileBytes([]byte(`{
			"OPEN_ISSUES": {
				"one": "There is one open issue in project {{.project}}.",
				"other": "There are {{.count}} open issues in project {{.project}}."
			},
			"COMPLETED_TASKS": {
				"one": "One task is completed in project {{.project}}.",
				"other": "{{.count}} tasks are completed in project {{.project}}."
			},
			"PENDING_REVIEWS": {
				"one": "There is one pending review for project {{.project}}.",
				"other": "There are {{.count}} pending reviews for project {{.project}}."
			},
			"ACTIVE_USERS": {
				"one": "One user is active on project {{.project}}.",
				"other": "{{.count}} users are active on project {{.project}}."
			}
		}`), "en-US.json"))
		zaplog.SUG.Debugln(neatjsons.S(messageFile))
		messageFiles = append(messageFiles, messageFile)
	}
	{
		messageFile := rese.P1(bundle.ParseMessageFileBytes([]byte(`{
			"OPEN_ISSUES": {
				"one": "项目 {{.project}} 中有一个未解决问题。",
				"other": "项目 {{.project}} 中有 {{.count}} 个未解决问题。"
			},
			"COMPLETED_TASKS": {
				"one": "项目 {{.project}} 中完成了一个任务。",
				"other": "项目 {{.project}} 中完成了 {{.count}} 个任务。"
			},
			"PENDING_REVIEWS": {
				"one": "项目 {{.project}} 中有一个待审。",
				"other": "项目 {{.project}} 中有 {{.count}} 个待审。"
			},
			"ACTIVE_USERS": {
				"one": "项目 {{.project}} 中有一个活跃用户。",
				"other": "项目 {{.project}} 中有 {{.count}} 个活跃用户。"
			}
		}`), "zh-CN.json"))
		zaplog.SUG.Debugln(neatjsons.S(messageFile))
		messageFiles = append(messageFiles, messageFile)
	}
	zaplog.SUG.Debugln(neatjsons.S(bundle.LanguageTags()))

	// Generate type-safe Go code with plural support from JSON message files
	// 从 JSON 消息文件生成支持复数的类型安全 Go 代码
	goi18n.Generate(messageFiles, goi18n.NewOptions())
}
