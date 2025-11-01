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

// Options configures code generation settings
// Controls output path, package name, and naming functions
// Supports custom naming for non-ASCII message IDs
//
// Options 配置代码生成设置
// 控制输出路径、包名和命名函数
// 支持非 ASCII 消息 ID 的自定义命名
type Options struct {
	outputPath         string
	pkgName            string
	allowNonAsciiRune  bool                          // 是否允许非 ASCII 字符，默认 false 值表示只允许 ASCII 字符
	unicodeMessageName func(messageID string) string // 基础命名
	unicodeStructName  func(messageID string) string // 类名命名
	unicodeFieldName   func(paramName string) string // 字段命名
	generateNewMessage bool
}

// NewOptions creates Options with default settings
// Sets up default naming functions for Unicode message IDs
//
// NewOptions 创建带默认设置的 Options
// 为 Unicode 消息 ID 设置默认命名函数
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

// WithOutputPath sets output file path
// Path must have .go extension
//
// WithOutputPath 设置输出文件路径
// 路径必须有 .go 扩展名
func (o *Options) WithOutputPath(outputPath string) *Options {
	must.Same(filepath.Ext(outputPath), ".go")
	o.outputPath = outputPath
	return o
}

// WithPkgName sets package name for generated code
//
// WithPkgName 设置生成代码的包名
func (o *Options) WithPkgName(pkgName string) *Options {
	o.pkgName = pkgName
	return o
}

// WithOutputPathWithPkgName sets output path and infers package name
// If file exists: reads package name from existing file
// If file not exists: derives package name from parent DIR name
//
// WithOutputPathWithPkgName 设置输出路径并推断包名
// 如果文件存在：从现有文件读取包名
// 如果文件不存在：从父目录名派生包名
func (o *Options) WithOutputPathWithPkgName(outputPath string) *Options {
	must.Same(filepath.Ext(outputPath), ".go")

	var pkgName string
	if osmustexist.IsFile(outputPath) {
		// Read package name from existing file
		// 从现有文件读取包名
		pkgName = syntaxgo.GetPkgName(outputPath)
	} else {
		// Derive package name from parent DIR name
		// 从父目录名派生包名
		pkgName = filepath.Base(filepath.Dir(outputPath))
		pkgName = strcase.ToSnake(pkgName)
		pkgName = strings.ReplaceAll(pkgName, "_", "")
		pkgName = strings.ToLower(pkgName)
	}

	o.outputPath = must.Nice(outputPath)
	o.pkgName = must.Nice(pkgName)
	return o
}

// GetOutputPath returns output file path
//
// GetOutputPath 返回输出文件路径
func (o *Options) GetOutputPath() string {
	return o.outputPath
}

// GetPkgName returns package name
//
// GetPkgName 返回包名
func (o *Options) GetPkgName() string {
	return o.pkgName
}

// WithAllowNonAsciiRune enables non-ASCII rune support in message IDs
// When enabled: uses custom naming functions for Unicode message IDs
// When disabled: uses strcase.ToCamel for ASCII-based naming
//
// WithAllowNonAsciiRune 启用消息 ID 中的非 ASCII 字符支持
// 启用时：对 Unicode 消息 ID 使用自定义命名函数
// 禁用时：使用 strcase.ToCamel 进行 ASCII 命名
func (o *Options) WithAllowNonAsciiRune(allowNonAsciiRune bool) *Options {
	o.allowNonAsciiRune = allowNonAsciiRune
	return o
}

// WithUnicodeMessageName sets custom naming function for Unicode message IDs
//
// WithUnicodeMessageName 设置 Unicode 消息 ID 的自定义命名函数
func (o *Options) WithUnicodeMessageName(unicodeMessageName func(string) string) *Options {
	o.unicodeMessageName = unicodeMessageName
	return o
}

// WithUnicodeStructName sets custom naming function for Unicode struct names
//
// WithUnicodeStructName 设置 Unicode 结构体名的自定义命名函数
func (o *Options) WithUnicodeStructName(unicodeStructName func(string) string) *Options {
	o.unicodeStructName = unicodeStructName
	return o
}

// WithUnicodeFieldName sets custom naming function for Unicode field names
//
// WithUnicodeFieldName 设置 Unicode 字段名的自定义命名函数
func (o *Options) WithUnicodeFieldName(unicodeFieldName func(string) string) *Options {
	o.unicodeFieldName = unicodeFieldName
	return o
}

// WithGenerateNewMessage enables generation of New* functions
// New* functions return (messageID, templateData) tuples
//
// WithGenerateNewMessage 启用生成 New* 函数
// New* 函数返回 (messageID, templateData) 元组
func (o *Options) WithGenerateNewMessage(generateNewMessage bool) *Options {
	o.generateNewMessage = generateNewMessage
	return o
}
