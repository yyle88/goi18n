package goi18n

import (
	"os"
	"path/filepath"
	"regexp"
	"strings"

	"github.com/emirpasic/gods/v2/maps/linkedhashmap"
	"github.com/emirpasic/gods/v2/sets/linkedhashset"
	"github.com/iancoleman/strcase"
	"github.com/nicksnyder/go-i18n/v2/i18n"
	"github.com/yyle88/formatgo"
	"github.com/yyle88/must"
	"github.com/yyle88/must/mustslice"
	"github.com/yyle88/neatjson/neatjsons"
	"github.com/yyle88/printgo"
	"github.com/yyle88/sortslice"
	"github.com/yyle88/syntaxgo/syntaxgo_ast"
	"github.com/yyle88/zaplog"
	"go.uber.org/zap"
)

type Options struct {
	OutputPath string
	PkgName    string
}

func NewOptions(outputPath string) *Options {
	must.Same(filepath.Ext(outputPath), ".go")

	pkgName := filepath.Base(filepath.Dir(outputPath))
	pkgName = strcase.ToSnake(pkgName)
	pkgName = strings.ReplaceAll(pkgName, "_", "")
	pkgName = strings.ToLower(pkgName)

	return &Options{
		OutputPath: outputPath,
		PkgName:    pkgName,
	}
}

func Generate(messageFiles []*i18n.MessageFile, options *Options) map[string]*Param {
	res := ParseParamNames(messageFiles)
	zaplog.SUG.Debugln(neatjsons.S(res))

	messageParams := SortMessageParams(res)
	zaplog.SUG.Debugln(neatjsons.S(messageParams))

	contentBytes := CreateMessageFunctions(messageParams)
	zaplog.SUG.Debugln(string(contentBytes))

	ptx := printgo.NewPTX()
	ptx.Println("package", must.Nice(options.PkgName))
	ptx.Println()
	ptx.Write(contentBytes)
	ptx.Println()

	contentBytes = ptx.Bytes()
	zaplog.SUG.Debugln(string(contentBytes))

	//把要引用的包写到代码的 import 里面（这样能提高format的速度，否则 format 还得找包，就会很慢，而且也未必能找到引用，因此这里主动设置引用）
	importOptions := syntaxgo_ast.NewPackageImportOptions().SetInferredObject(i18n.Message{})
	contentBytes = importOptions.InjectImports(contentBytes)

	//调整代码格式和风格。注意：这里即使出错也能返回原来的代码，而且内部已有 warn 级别的日志，因此直接忽略错误
	contentBytes, _ = formatgo.FormatBytes(contentBytes)
	must.Have(contentBytes)

	must.Done(os.WriteFile(must.Nice(options.OutputPath), contentBytes, 0666))
	return res
}

type Param struct {
	Names           *linkedhashset.Set[string]
	HasOneAnonymous bool
}

func ParseParamNames(messageFiles []*i18n.MessageFile) map[string]*Param {
	res := map[string]*Param{}

	regExp := regexp.MustCompile(`\{\{\s*\.(\w*)\s*}}`)
	for _, messageFile := range messageFiles {
		for _, message := range messageFile.Messages {
			param, ok := res[message.ID]
			if !ok {
				param = &Param{
					Names: linkedhashset.New[string](),
				}
				res[message.ID] = param
			}

			for _, msgTemplate := range []string{
				message.Other, //优先使用这个字段，因为我比较懒惰通常只配置默认的翻译，而不会精细化的使用单复数
				message.One,
				message.Zero,
				message.Two,
				message.Few,
				message.Many,
			} {
				names := parseParamNames(msgTemplate, regExp)
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
		}
	}
	return res
}

func parseParamNames(msg string, regExp *regexp.Regexp) []string {
	var results = make([]string, 0)
	if msg != "" {
		subs := regExp.FindAllStringSubmatch(msg, -1)
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

func CreateMessageFunctions(messageParams []*MessageParam) []byte {
	ptx := printgo.NewPTX()
	for _, messageParam := range messageParams {
		writeNewMsgFunction(ptx, messageParam)
		ptx.Println()
	}
	return ptx.Bytes()
}

func writeNewMsgFunction(ptx *printgo.PTX, messageParam *MessageParam) {
	messageName := strcase.ToCamel(messageParam.MessageID)

	if messageParam.Param.HasOneAnonymous {
		ptx.Println("func New"+messageName+"[Value comparable](value Value)", "(string, Value) {")
		ptx.Println("\treturn", `"`+messageParam.MessageID+`"`, ",", "value")
		ptx.Println("}")
		ptx.Println()

		ptx.Println("func I18n"+messageName+"[Value comparable](value Value)", "*i18n.LocalizeConfig {")
		ptx.Println("\treturn &i18n.LocalizeConfig{")
		ptx.Println("\t\tMessageID:", `"`+messageParam.MessageID+`"`, ",")
		ptx.Println("\t\tTemplateData: value", ",")
		ptx.Println("\t}")
		ptx.Println("}")
	} else if !messageParam.Param.Names.Empty() {
		structName := messageName + "Param"
		fieldNames := linkedhashmap.New[string, string]()
		for _, name := range messageParam.Param.Names.Values() {
			fieldNames.Put(name, strcase.ToCamel(name))
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

		ptx.Println("func I18n"+messageName+"(data *", structName, ")", "*i18n.LocalizeConfig {")
		ptx.Println("\treturn &i18n.LocalizeConfig{")
		ptx.Println("\t\tMessageID:", `"`+messageParam.MessageID+`"`, ",")
		ptx.Println("\t\tTemplateData:", "data.", methodName, "()", ",")
		ptx.Println("\t}")
		ptx.Println("}")
	} else {
		ptx.Println("func New"+messageName+"()", "string {")
		ptx.Println("\treturn", `"`+messageParam.MessageID+`"`)
		ptx.Println("}")
		ptx.Println()

		ptx.Println("func I18n"+messageName+"()", "*i18n.LocalizeConfig {")
		ptx.Println("\treturn &i18n.LocalizeConfig{")
		ptx.Println("\t\tMessageID:", `"`+messageParam.MessageID+`"`, ",")
		ptx.Println("\t}")
		ptx.Println("}")
	}
}
