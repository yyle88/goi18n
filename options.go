package goi18n

import (
	"path/filepath"
	"strings"

	"github.com/iancoleman/strcase"
	"github.com/yyle88/goi18n/internal/utils"
	"github.com/yyle88/must"
)

type Options struct {
	OutputPath         string
	PkgName            string
	allowNonAsciiRune  bool                          // 是否允许非 ASCII 字符，默认 false 值表示只允许 ASCII 字符
	unicodeMessageName func(messageID string) string // 基础命名
	unicodeStructName  func(messageID string) string // 类名命名
	unicodeFieldName   func(paramName string) string // 字段命名
}

func NewOptions() *Options {
	return &Options{
		OutputPath:        "",
		PkgName:           "",
		allowNonAsciiRune: false,
		unicodeMessageName: func(messageID string) string {
			s := messageID
			if utils.HasLetterPrefix(s) {
				return utils.CapitalizeFirst(s)
			}
			return "I" + s
		},
		unicodeStructName: func(messageID string) string {
			s := messageID
			if utils.HasLetterPrefix(s) {
				return utils.CapitalizeFirst(s)
			}
			return "P" + s
		},
		unicodeFieldName: func(paramName string) string {
			s := paramName
			if utils.HasLetterPrefix(s) {
				return utils.CapitalizeFirst(s)
			}
			return "V" + s
		},
	}
}

func (o *Options) WithOutputPath(outputPath string) *Options {
	must.Same(filepath.Ext(outputPath), ".go")
	o.OutputPath = outputPath
	return o
}

func (o *Options) WithPkgName(pkgName string) *Options {
	o.PkgName = pkgName
	return o
}

func (o *Options) WithOutputPathWithPkgName(outputPath string) *Options {
	must.Same(filepath.Ext(outputPath), ".go")

	pkgName := filepath.Base(filepath.Dir(outputPath))
	pkgName = strcase.ToSnake(pkgName)
	pkgName = strings.ReplaceAll(pkgName, "_", "")
	pkgName = strings.ToLower(pkgName)

	o.OutputPath = outputPath
	o.PkgName = pkgName
	return o
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
