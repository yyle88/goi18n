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
	"github.com/yyle88/osexistpath/osmustexist"
	"github.com/yyle88/printgo"
	"github.com/yyle88/rese"
	"github.com/yyle88/sortslice"
	"github.com/yyle88/syntaxgo/syntaxgo_ast"
	"github.com/yyle88/tern"
	"github.com/yyle88/zaplog"
	"go.uber.org/zap"
)

func Generate(messageFiles []*i18n.MessageFile, options *Options) {
	mapParams := ParseParamNames(messageFiles, options)
	zaplog.SUG.Debugln(neatjsons.S(mapParams))

	messageParams := SortMessageParams(mapParams)
	zaplog.SUG.Debugln(neatjsons.S(messageParams))

	contentBytes := CreateMessageFunctions(messageParams, options)
	zaplog.SUG.Debugln(string(contentBytes))

	if options.OutputPath != "" && options.PkgName != "" {
		WriteContentToCodeFile(contentBytes, options)
	} else {
		zaplog.SUG.Debugln("NOT write-content-to-code-file RETURN")
	}
}

type Param struct {
	Names           *linkedhashset.Set[string]
	HasOneAnonymous bool
	NeedPluralCount bool
}

func ParseParamNames(messageFiles []*i18n.MessageFile, options *Options) map[string]*Param {
	res := map[string]*Param{}

	regexpRegexp := tern.BFF(!options.allowNonAsciiRune, func() *regexp.Regexp {
		// 匹配 {{ .变量名 }}，变量名由 ASCII 字母数字下划线组成，非贪婪匹配，允许两边空白。
		// Matches {{ .variable }}, where variable contains ASCII letters, digits, or underscores; uses non-greedy matching and allows optional surrounding whitespace.
		regexpRegexp := regexp.MustCompile(`\{\{\s*\.(\w*?)\s*}}`)
		return regexpRegexp
	}, func() *regexp.Regexp {
		// 匹配 {{ .变量名 }}，变量名由任意非空白字符组成，非贪婪匹配，允许两边空白。
		// Matches {{ .variable }}, where variable contains any non-whitespace characters; uses non-greedy matching and allows optional surrounding whitespace.
		regexpRegexp := regexp.MustCompile(`\{\{\s*\.(\S*?)\s*}}`)
		return regexpRegexp
	})

	for _, messageFile := range messageFiles {
		for _, message := range messageFile.Messages {
			param, ok := res[message.ID]
			if !ok {
				param = &Param{
					Names: linkedhashset.New[string](),
				}
				res[message.ID] = param
			}

			messageTemplates := []string{
				message.Other, //优先使用这个字段，因为我比较懒惰通常只配置默认的翻译，而不会精细化的使用单复数
				message.One,
				message.Zero,
				message.Two,
				message.Few,
				message.Many,
			}

			for _, msgTemplate := range messageTemplates {
				names := parseParamNames(msgTemplate, regexpRegexp)
				for _, name := range names {
					if name == "" {
						mustslice.Length(names, 1)     //匿名参数只能出现一次，而且不能再有其它参数
						must.True(param.Names.Empty()) //当有匿名参数时就不能有其它参数，在其它翻译里也是这样的
						param.HasOneAnonymous = true
					} else {
						must.FALSE(param.HasOneAnonymous) //当有参数时就不能有匿名参数
						param.Names.Add(name)
					}
				}
			}

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

func parseParamNames(msg string, regexpRegexp *regexp.Regexp) []string {
	var results = make([]string, 0)
	if msg != "" {
		subs := regexpRegexp.FindAllStringSubmatch(msg, -1)
		zaplog.LOG.Debug("parse-param-names", zap.String("msg", msg), zap.Any("subs", subs))
		for _, sub := range subs {
			mustslice.Length(sub, 2)
			results = append(results, sub[1])
		}
	}
	zaplog.LOG.Debug("parse-param-names", zap.String("msg", msg), zap.Any("results", results))
	return results
}

func SortMessageParams(res map[string]*Param) []*MessageParam {
	var messageParams = make([]*MessageParam, 0, len(res))
	for messageID, param := range res {
		messageParams = append(messageParams, &MessageParam{
			MessageID: messageID,
			Param:     param,
		})
	}
	sortslice.SortByValue(messageParams, func(a, b *MessageParam) bool {
		return a.MessageID < b.MessageID
	})
	return messageParams
}

type MessageParam struct {
	MessageID string
	Param     *Param
}

func CreateMessageFunctions(messageParams []*MessageParam, options *Options) []byte {
	ptx := printgo.NewPTX()
	for _, messageParam := range messageParams {
		writeNewMsgFunction(ptx, messageParam, options)
		ptx.Println()
	}
	return ptx.Bytes()
}

func writeNewMsgFunction(ptx *printgo.PTX, messageParam *MessageParam, options *Options) {
	var messageName string
	if options.allowNonAsciiRune && utils.HasNonASCII(messageParam.MessageID) {
		messageName = options.unicodeMessageName(messageParam.MessageID)
	} else {
		messageName = strcase.ToCamel(messageParam.MessageID)
	}
	must.Nice(messageName)

	if messageParam.Param.HasOneAnonymous {
		ptx.Println("func New"+messageName+"[Value comparable](value Value)", "(string, Value) {")
		ptx.Println("\treturn", `"`+messageParam.MessageID+`"`, ",", "value")
		ptx.Println("}")
		ptx.Println()

		ptx.Print("func I18n" + messageName + "[Value comparable](value Value,")
		if messageParam.Param.NeedPluralCount {
			ptx.Print("pluralConfig *goi18n.PluralConfig,")
		}
		ptx.Println(")", "*i18n.LocalizeConfig {")
		ptx.Println("\treturn &i18n.LocalizeConfig{")
		ptx.Println("\t\tMessageID:", `"`+messageParam.MessageID+`"`, ",")
		ptx.Println("\t\tTemplateData: value", ",")
		if messageParam.Param.NeedPluralCount {
			ptx.Println("\t\tPluralCount: pluralConfig.PluralCount,")
		}
		ptx.Println("\t}")
		ptx.Println("}")
	} else if !messageParam.Param.Names.Empty() {
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

		ptx.Println("func New"+messageName+"(data *", structName, ")", "(string, map[string]any) {")
		ptx.Println("\treturn", `"`+messageParam.MessageID+`"`, ",", "data.", methodName, "()")
		ptx.Println("}")
		ptx.Println()

		ptx.Print("func I18n"+messageName+"(data *", structName, ",")
		if messageParam.Param.NeedPluralCount {
			ptx.Print("pluralConfig *goi18n.PluralConfig,")
		}
		ptx.Println(")", "*i18n.LocalizeConfig {")
		ptx.Println("\treturn &i18n.LocalizeConfig{")
		ptx.Println("\t\tMessageID:", `"`+messageParam.MessageID+`"`, ",")
		ptx.Println("\t\tTemplateData:", "data.", methodName, "()", ",")
		if messageParam.Param.NeedPluralCount {
			ptx.Println("\t\tPluralCount: pluralConfig.PluralCount,")
		}
		ptx.Println("\t}")
		ptx.Println("}")
	} else {
		ptx.Println("func New"+messageName+"()", "string {")
		ptx.Println("\treturn", `"`+messageParam.MessageID+`"`)
		ptx.Println("}")
		ptx.Println()

		ptx.Print("func I18n" + messageName + "(")
		if messageParam.Param.NeedPluralCount {
			ptx.Print("pluralConfig *goi18n.PluralConfig,")
		}
		ptx.Println(")", "*i18n.LocalizeConfig {")
		ptx.Println("\treturn &i18n.LocalizeConfig{")
		ptx.Println("\t\tMessageID:", `"`+messageParam.MessageID+`"`, ",")
		if messageParam.Param.NeedPluralCount {
			ptx.Println("\t\tPluralCount: pluralConfig.PluralCount,")
		}
		ptx.Println("\t}")
		ptx.Println("}")
	}
}

func WriteContentToCodeFile(contentBytes []byte, options *Options) {
	ptx := printgo.NewPTX()
	ptx.Println("package", must.Nice(options.PkgName))
	ptx.Println()
	ptx.Write(contentBytes)
	ptx.Println()

	contentBytes = ptx.Bytes()
	zaplog.SUG.Debugln(string(contentBytes))

	//把要引用的包写到代码的 import 里面（这样能提高format的速度，否则 format 还得找包，就会很慢，而且也未必能找到引用，因此这里主动设置引用）
	importOptions := syntaxgo_ast.NewPackageImportOptions()
	importOptions.SetInferredObject(&i18n.Message{})
	importOptions.SetInferredObject(&MessageParam{})
	contentBytes = importOptions.InjectImports(contentBytes)

	//调整代码格式和风格。注意：这里即使出错也能返回原来的代码，而且内部已有 warn 级别的日志，因此直接忽略错误
	contentBytes, _ = formatgo.FormatBytes(contentBytes)
	must.Have(contentBytes)

	path := must.Nice(options.OutputPath)

	var perm os.FileMode = 0666 // default file-mode-perm
	if osmustexist.IsFile(path) {
		perm = rese.V1(os.Stat(path)).Mode()
	}
	must.Done(os.WriteFile(path, contentBytes, perm))
}
