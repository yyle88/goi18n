package goi18n

import (
	"path/filepath"
	"strings"

	"github.com/iancoleman/strcase"
	"github.com/yyle88/goi18n/internal/utils"
	"github.com/yyle88/must"
	"github.com/yyle88/osexistpath/osmustexist"
	"github.com/yyle88/syntaxgo"
)

type Options struct {
	outputPath         string
	pkgName            string
	allowNonAsciiRune  bool                          // 是否允许非 ASCII 字符，默认 false 值表示只允许 ASCII 字符
	unicodeMessageName func(messageID string) string // 基础命名
	unicodeStructName  func(messageID string) string // 类名命名
	unicodeFieldName   func(paramName string) string // 字段命名
	generateNewMessage bool
}

func NewOptions() *Options {
	return &Options{
		outputPath:        "",
		pkgName:           "",
		allowNonAsciiRune: false,
		unicodeMessageName: func(messageID string) string {
			return utils.DefaultUnicodeMessageName(messageID)
		},
		unicodeStructName: func(messageID string) string {
			return utils.DefaultUnicodeStructName(messageID)
		},
		unicodeFieldName: func(paramName string) string {
			return utils.DefaultUnicodeFieldName(paramName)
		},
		generateNewMessage: false,
	}
}

func (o *Options) WithOutputPath(outputPath string) *Options {
	must.Same(filepath.Ext(outputPath), ".go")
	o.outputPath = outputPath
	return o
}

func (o *Options) WithPkgName(pkgName string) *Options {
	o.pkgName = pkgName
	return o
}

func (o *Options) WithOutputPathWithPkgName(outputPath string) *Options {
	must.Same(filepath.Ext(outputPath), ".go")

	var pkgName string
	if osmustexist.IsFile(outputPath) { // use package name in source file code
		pkgName = syntaxgo.GetPkgName(outputPath)
	} else {
		pkgName = filepath.Base(filepath.Dir(outputPath))
		pkgName = strcase.ToSnake(pkgName)
		pkgName = strings.ReplaceAll(pkgName, "_", "")
		pkgName = strings.ToLower(pkgName)
	}

	o.outputPath = must.Nice(outputPath)
	o.pkgName = must.Nice(pkgName)
	return o
}

func (o *Options) GetOutputPath() string {
	return o.outputPath
}

func (o *Options) GetPkgName() string {
	return o.pkgName
}

func (o *Options) WithAllowNonAsciiRune(allowNonAsciiRune bool) *Options {
	o.allowNonAsciiRune = allowNonAsciiRune
	return o
}

func (o *Options) WithUnicodeMessageName(unicodeMessageName func(string) string) *Options {
	o.unicodeMessageName = unicodeMessageName
	return o
}

func (o *Options) WithUnicodeStructName(unicodeStructName func(string) string) *Options {
	o.unicodeStructName = unicodeStructName
	return o
}

func (o *Options) WithUnicodeFieldName(unicodeFieldName func(string) string) *Options {
	o.unicodeFieldName = unicodeFieldName
	return o
}

func (o *Options) WithGenerateNewMessage(generateNewMessage bool) *Options {
	o.generateNewMessage = generateNewMessage
	return o
}
