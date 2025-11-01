// Package goi18n provides code generation for go-i18n with type-safe message handling
// Auto generates Go structs and functions from i18n message files
// Supports named params, anonymous params, and message without params
// Integrates with go-i18n package to provide type-safe translation
//
// Package goi18n 为 go-i18n 提供类型安全的代码生成
// 从国际化消息文件自动生成 Go 结构体和函数
// 支持命名参数、匿名参数和无参数消息
// 集成 go-i18n 包以提供类型安全的翻译
package goi18n

import (
	"os"
	"regexp"

	"github.com/emirpasic/gods/v2/maps/linkedhashmap"
	"github.com/emirpasic/gods/v2/sets/linkedhashset"
	"github.com/iancoleman/strcase"
	"github.com/nicksnyder/go-i18n/v2/i18n"
	"github.com/yyle88/formatgo"
	"github.com/yyle88/goi18n/internal/utils"
	"github.com/yyle88/must"
	"github.com/yyle88/must/mustslice"
	"github.com/yyle88/neatjson/neatjsons"
	"github.com/yyle88/printgo"
	"github.com/yyle88/sortx"
	"github.com/yyle88/syntaxgo/syntaxgo_ast"
	"github.com/yyle88/tern"
	"github.com/yyle88/zaplog"
	"go.uber.org/zap"
)

// Generate creates type-safe Go code from i18n message files
// Parses message files, extracts params, and generates structs and functions
// Output code provides type-safe wrappers for go-i18n LocalizeConfig
//
// Generate 从国际化消息文件生成类型安全的 Go 代码
// 解析消息文件，提取参数，生成结构体和函数
// 输出代码为 go-i18n LocalizeConfig 提供类型安全的包装
func Generate(messageFiles []*i18n.MessageFile, options *Options) {
	mapParams := ParseParamNames(messageFiles, options)
	zaplog.SUG.Debugln(neatjsons.S(mapParams))

	messageParams := SortMessageParams(mapParams)
	zaplog.SUG.Debugln(neatjsons.S(messageParams))

	srcCode := CreateMessageFunctions(messageParams, options)
	zaplog.SUG.Debugln(string(srcCode))

	if options.outputPath != "" && options.pkgName != "" {
		WriteContentToCodeFile(srcCode, options)
	} else {
		zaplog.SUG.Debugln("NOT write-content-to-code-file RETURN")
	}
}

// Param contains message param metadata for code generation
// Names: named params extracted from message template
// HasOneAnonymous: whether message uses anonymous param {{ . }}
// NeedPluralCount: whether message needs count for one/other forms
//
// Param 包含用于代码生成的消息参数元数据
// Names: 从消息模板提取的命名参数
// HasOneAnonymous: 消息是否使用匿名参数 {{ . }}
// NeedPluralCount: 消息是否需要 one/other 形式的计数
type Param struct {
	Names           *linkedhashset.Set[string]
	HasOneAnonymous bool
	NeedPluralCount bool
}

// ParseParamNames extracts params from message files and builds metadata map
// Scans message templates using regex to find named and anonymous params
// Detects when messages need count for one/other forms
// Returns map of message ID to Param metadata
//
// ParseParamNames 从消息文件提取参数并构建元数据映射
// 使用正则扫描消息模板查找命名和匿名参数
// 检测消息是否需要 one/other 形式的计数
// 返回消息 ID 到 Param 元数据的映射
func ParseParamNames(messageFiles []*i18n.MessageFile, options *Options) map[string]*Param {
	res := map[string]*Param{}

	// Select regex pattern based on ASCII vs Unicode support
	// ASCII mode: \w matches [a-zA-Z0-9_]
	// Unicode mode: \S matches any non-whitespace character
	//
	// 根据 ASCII 或 Unicode 支持选择正则模式
	// ASCII 模式：\w 匹配 [a-zA-Z0-9_]
	// Unicode 模式：\S 匹配任意非空白字符
	pattern := tern.BFF(!options.allowNonAsciiRune, func() *regexp.Regexp {
		// Match {{ .variable }}, variable contains ASCII letters/digits/underscores
		// 匹配 {{ .变量名 }}，变量名由 ASCII 字母数字下划线组成
		return regexp.MustCompile(`\{\{\s*\.(\w*?)\s*}}`)
	}, func() *regexp.Regexp {
		// Match {{ .variable }}, variable contains any non-whitespace chars
		// 匹配 {{ .变量名 }}，变量名由任意非空白字符组成
		return regexp.MustCompile(`\{\{\s*\.(\S*?)\s*}}`)
	})

	for _, messageFile := range messageFiles {
		for _, message := range messageFile.Messages {
			// Get or create Param metadata for this message ID
			// 获取或创建此消息 ID 的 Param 元数据
			param, ok := res[message.ID]
			if !ok {
				param = &Param{
					Names: linkedhashset.New[string](),
				}
				res[message.ID] = param
			}

			// Process all forms: Other/One/Zero/Two/Few/Many
			// 处理所有形式：Other/One/Zero/Two/Few/Many
			messageTemplates := []string{
				message.Other, // Process default translation first // 优先处理默认翻译
				message.One,
				message.Zero,
				message.Two,
				message.Few,
				message.Many,
			}

			// Extract params from each template form
			// 从每个模板形式提取参数
			for _, msgTemplate := range messageTemplates {
				names := parseParamNames(msgTemplate, pattern)
				for _, name := range names {
					if name == "" {
						// Anonymous param {{ . }} - must be the one param, no others allowed
						// 匿名参数 {{ . }} - 必须是唯一参数，不允许其他参数
						mustslice.Length(names, 1)
						must.True(param.Names.Empty())
						param.HasOneAnonymous = true
					} else {
						// Named param {{ .name }} - cannot mix with anonymous param
						// 命名参数 {{ .name }} - 不能与匿名参数混合
						must.FALSE(param.HasOneAnonymous)
						param.Names.Add(name)
					}
				}
			}

			// Detect if message needs count for one/other forms
			// Count >= 2 means at least two forms defined (e.g., Other + One)
			//
			// 检测消息是否需要 one/other 形式的计数
			// 计数 >= 2 表示至少定义了两种形式（例如 Other + One）
			var count = 0
			for _, msgTemplate := range messageTemplates {
				if msgTemplate != "" {
					count++
				}
			}
			if count >= 2 {
				param.NeedPluralCount = true
			}
		}
	}
	return res
}

// parseParamNames extracts param names from message template using regex
// Matches {{ .paramName }} patterns and returns param names
// Returns empty slice for blank messages
//
// parseParamNames 使用正则从消息模板提取参数名
// 匹配 {{ .参数名 }} 模式并返回参数名
// 空消息返回空切片
func parseParamNames(msg string, pattern *regexp.Regexp) []string {
	var results = make([]string, 0)
	if msg != "" {
		// Extract all {{ .param }} matches from template
		// 从模板提取所有 {{ .param }} 匹配
		subs := pattern.FindAllStringSubmatch(msg, -1)
		zaplog.LOG.Debug("parse-param-names", zap.String("msg", msg), zap.Any("subs", subs))
		for _, sub := range subs {
			mustslice.Length(sub, 2) // sub[0] is full match, sub[1] is param name // sub[0] 是完整匹配，sub[1] 是参数名
			results = append(results, sub[1])
		}
	}
	zaplog.LOG.Debug("parse-param-names", zap.String("msg", msg), zap.Any("results", results))
	return results
}

// SortMessageParams converts param map to sorted slice of MessageParam
// Sorts messages alphabetically by ID for stable code generation
// Uses sortx to maintain consistent output across runs
//
// SortMessageParams 将参数映射转换为排序的 MessageParam 切片
// 按 ID 字母顺序排序消息以保证稳定的代码生成
// 使用 sortx 保持多次运行输出一致
func SortMessageParams(res map[string]*Param) []*MessageParam {
	var messageParams = make([]*MessageParam, 0, len(res))
	for messageID, param := range res {
		messageParams = append(messageParams, &MessageParam{
			MessageID: messageID,
			Param:     param,
		})
	}
	sortx.SortByValue(messageParams, func(a, b *MessageParam) bool {
		return a.MessageID < b.MessageID
	})
	return messageParams
}

// MessageParam combines message ID with param metadata
// Used in sorted slice for stable code generation
//
// MessageParam 结合消息 ID 和参数元数据
// 用于排序切片以保证稳定的代码生成
type MessageParam struct {
	MessageID string
	Param     *Param
}

// CreateMessageFunctions generates Go code for message functions
// Creates type-safe structs and functions for each message
// Handles named params, anonymous params, and messages without params
//
// CreateMessageFunctions 为消息函数生成 Go 代码
// 为每个消息创建类型安全的结构体和函数
// 处理命名参数、匿名参数和无参数消息
func CreateMessageFunctions(messageParams []*MessageParam, options *Options) []byte {
	ptx := printgo.NewPTX()
	for _, messageParam := range messageParams {
		writeNewMsgFunction(ptx, messageParam, options)
		ptx.Println()
	}
	return ptx.Bytes()
}

// writeNewMsgFunction generates code for single message
// Handles three cases: anonymous param, named params, or no params
// Generates struct definitions and I18n* functions for each message
//
// writeNewMsgFunction 为单个消息生成代码
// 处理三种情况：匿名参数、命名参数或无参数
// 为每个消息生成结构体定义和 I18n* 函数
func writeNewMsgFunction(ptx *printgo.PTX, messageParam *MessageParam, options *Options) {
	// Convert message ID to function name (PascalCase or custom naming)
	// 将消息 ID 转换为函数名（驼峰命名或自定义命名）
	var messageName string
	if options.allowNonAsciiRune && utils.HasNonASCII(messageParam.MessageID) {
		messageName = options.unicodeMessageName(messageParam.MessageID)
	} else {
		messageName = strcase.ToCamel(messageParam.MessageID)
	}
	must.Nice(messageName)

	// Case 1: Anonymous param {{ . }} - generates generic function
	// 情况1：匿名参数 {{ . }} - 生成泛型函数
	if messageParam.Param.HasOneAnonymous {
		if options.generateNewMessage {
			ptx.Println("func New"+messageName+"[Value comparable](value Value)", "(string, Value) {")
			ptx.Println("\treturn", `"`+messageParam.MessageID+`"`, ",", "value")
			ptx.Println("}")
			ptx.Println()
		}

		ptx.Print("func I18n" + messageName + "[Value comparable](value Value,")
		if messageParam.Param.NeedPluralCount {
			ptx.Print("pluralConfig *goi18n.PluralConfig,")
		}
		ptx.Println(")", "*i18n.LocalizeConfig {")

		if options.generateNewMessage {
			ptx.Println("messageID, tempValue := New" + messageName + "(value)")
		} else {
			ptx.Println("const messageID =", `"`+messageParam.MessageID+`"`)
			ptx.Println("var tempValue = value")
		}

		ptx.Println("\treturn &i18n.LocalizeConfig{")
		ptx.Println("\t\tMessageID: messageID", ",")
		ptx.Println("\t\tTemplateData: tempValue", ",")
		if messageParam.Param.NeedPluralCount {
			ptx.Println("\t\tPluralCount: pluralConfig.PluralCount,")
		}
		ptx.Println("\t}")
		ptx.Println("}")
	} else if !messageParam.Param.Names.Empty() {
		// Case 2: Named params {{ .name }} {{ .code }} - generates struct and methods
		// 情况2：命名参数 {{ .name }} {{ .code }} - 生成结构体和方法
		var structName string
		if options.allowNonAsciiRune && utils.HasNonASCII(messageParam.MessageID) {
			structName = options.unicodeStructName(messageParam.MessageID)
		} else {
			structName = messageName + "Param"
		}
		must.Nice(structName)

		fieldNames := linkedhashmap.New[string, string]()
		for _, paramName := range messageParam.Param.Names.Values() {
			var fieldName string
			if options.allowNonAsciiRune && utils.HasNonASCII(paramName) {
				fieldName = options.unicodeFieldName(paramName)
			} else {
				fieldName = strcase.ToCamel(paramName)
			}
			must.Nice(fieldName)

			fieldNames.Put(paramName, fieldName)
		}
		methodName := "GetTemplateValues"

		ptx.Println("type", structName, "struct {")
		fieldNames.Each(func(key string, camelcase string) {
			ptx.Println("\t", camelcase, "any")
		})
		ptx.Println("}")
		ptx.Println()

		ptx.Println("func (p *", structName, ") ", methodName, "() map[string]any {")
		ptx.Println("\tres := make(map[string]any)")
		fieldNames.Each(func(key string, camelcase string) {
			ptx.Println("\tif p.", camelcase, " != nil {")
			ptx.Println("\t\tres[", `"`+key+`"`, "] = p.", camelcase)
			ptx.Println("\t}")
		})
		ptx.Println("\treturn res")
		ptx.Println("}")
		ptx.Println()

		if options.generateNewMessage {
			ptx.Println("func New"+messageName+"(data *", structName, ")", "(string, map[string]any) {")
			ptx.Println("\treturn", `"`+messageParam.MessageID+`"`, ",", "data.", methodName, "()")
			ptx.Println("}")
			ptx.Println()
		}

		ptx.Print("func I18n"+messageName+"(data *", structName, ",")
		if messageParam.Param.NeedPluralCount {
			ptx.Print("pluralConfig *goi18n.PluralConfig,")
		}
		ptx.Println(")", "*i18n.LocalizeConfig {")

		if options.generateNewMessage {
			ptx.Println("messageID, valuesMap := New" + messageName + "(data)")
		} else {
			ptx.Println("const messageID =", `"`+messageParam.MessageID+`"`)
			ptx.Println("var valuesMap = data.", methodName, "()")
		}

		ptx.Println("\treturn &i18n.LocalizeConfig{")
		ptx.Println("\t\tMessageID: messageID", ",")
		ptx.Println("\t\tTemplateData: valuesMap", ",")
		if messageParam.Param.NeedPluralCount {
			ptx.Println("\t\tPluralCount: pluralConfig.PluralCount,")
		}
		ptx.Println("\t}")
		ptx.Println("}")
	} else {
		// Case 3: No params - generates simple function returning config
		// 情况3：无参数 - 生成简单函数返回配置
		if options.generateNewMessage {
			ptx.Println("func New"+messageName+"()", "string {")
			ptx.Println("\treturn", `"`+messageParam.MessageID+`"`)
			ptx.Println("}")
			ptx.Println()
		}

		ptx.Print("func I18n" + messageName + "(")
		if messageParam.Param.NeedPluralCount {
			ptx.Print("pluralConfig *goi18n.PluralConfig,")
		}
		ptx.Println(")", "*i18n.LocalizeConfig {")

		if options.generateNewMessage {
			ptx.Println("messageID := New" + messageName + "()")
		} else {
			ptx.Println("const messageID =", `"`+messageParam.MessageID+`"`)
		}

		ptx.Println("\treturn &i18n.LocalizeConfig{")
		ptx.Println("\t\tMessageID: messageID", ",")
		if messageParam.Param.NeedPluralCount {
			ptx.Println("\t\tPluralCount: pluralConfig.PluralCount,")
		}
		ptx.Println("\t}")
		ptx.Println("}")
	}
}

// WriteContentToCodeFile writes generated code to output file
// Adds package declaration, injects imports, and formats code
// Uses syntaxgo_ast to auto inject needed imports
// Uses formatgo to format the generated code
//
// WriteContentToCodeFile 将生成的代码写入输出文件
// 添加包声明，注入导入，格式化代码
// 使用 syntaxgo_ast 自动注入所需导入
// 使用 formatgo 格式化生成的代码
func WriteContentToCodeFile(srcCode []byte, options *Options) {
	ptx := printgo.NewPTX()
	ptx.Println("package", must.Nice(options.pkgName))
	ptx.Println()
	ptx.Write(srcCode)
	ptx.Println()

	srcCode = ptx.Bytes()
	zaplog.SUG.Debugln(string(srcCode))

	//把要引用的包写到代码的 import 里面（这样能提高format的速度，否则 format 还得找包，就会很慢，而且也未必能找到引用，因此这里主动设置引用）
	importOptions := syntaxgo_ast.NewPackageImportOptions()
	importOptions.SetInferredObject(&i18n.Message{})
	importOptions.SetInferredObject(&MessageParam{})
	srcCode = importOptions.InjectImports(srcCode)

	//调整代码格式和风格。注意：这里即使出错也能返回原来的代码，而且内部已有 warn 级别的日志，因此直接忽略错误
	srcCode, _ = formatgo.FormatBytes(srcCode)
	//前面不判断是否有错，只需要判断结果有没有内容，格式化出错也不影响结果
	must.Have(srcCode)

	path := must.Nice(options.outputPath)

	// when file exist WriteFile truncates it before writing, without changing permissions.
	must.Done(os.WriteFile(path, srcCode, 0666))
}
