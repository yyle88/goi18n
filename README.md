[![GitHub Workflow Status (branch)](https://img.shields.io/github/actions/workflow/status/yyle88/goi18n/release.yml?branch=main&label=BUILD)](https://github.com/yyle88/goi18n/actions/workflows/release.yml?query=branch%3Amain)
[![GoDoc](https://pkg.go.dev/badge/github.com/yyle88/goi18n)](https://pkg.go.dev/github.com/yyle88/goi18n)
[![Coverage Status](https://img.shields.io/coveralls/github/yyle88/goi18n/main.svg)](https://coveralls.io/github/yyle88/goi18n?branch=main)
[![Supported Go Versions](https://img.shields.io/badge/Go-1.22--1.25-lightgrey.svg)](https://github.com/yyle88/goi18n)
[![GitHub Release](https://img.shields.io/github/release/yyle88/goi18n.svg)](https://github.com/yyle88/goi18n/releases)
[![Go Report Card](https://goreportcard.com/badge/github.com/yyle88/goi18n)](https://goreportcard.com/report/github.com/yyle88/goi18n)

# Goi18n

`goi18n` replace map[string]interface{} with generic parameters to make go-i18n more concise.

## Overview

`goi18n` is a Go package and code generation toolkit that simplifies internationalization (i18n) in Go applications. It processes i18n message files (e.g., YAML) and generates type-safe Go code, including structs and functions to handle messages. The generated code integrates with the `go-i18n` package, enabling efficient and safe translation rendering across multiple languages.

## Why goi18n?

When working with the excellent [go-i18n](https://github.com/nicksnyder/go-i18n) package, you often need to write repetitive code like this:

```go
// Traditional approach - verbose and error-prone
localizer.Localize(&i18n.LocalizeConfig{
    MessageID: "ERROR_NOT_EXIST",
    TemplateData: map[string]interface{}{
        "what": "User",
        "code": "12345",
    },
})
```

This approach has multiple issues:
- **No type safety**: Simple to make typos in message IDs and param names
- **Runtime errors**: Mistakes are caught at runtime, not compile time
- **Limited IDE support**: No auto-completion on message IDs and params
- **Maintenance burden**: Hard to track which params each message needs

**goi18n solves these problems** by generating type-safe code:

```go
// With goi18n - type-safe and clean
message1.I18nErrorNotExist(&message1.ErrorNotExistParam{
    What: "User",
    Code: "12345",
})
```

Benefits:
- ✅ **Compile-time safety**: Catch errors before deployment
- ✅ **IDE auto-completion**: Complete IntelliSense support
- ✅ **Refactoring friendly**: Rename params with confidence
- ✅ **Self-documenting**: Generated structs show required params
- ✅ **Zero runtime overhead**: Same performance as hand-written code

<!-- TEMPLATE (EN) BEGIN: LANGUAGE NAVIGATION -->
## CHINESE README

[中文说明](README.zh.md)
<!-- TEMPLATE (EN) END: LANGUAGE NAVIGATION -->

## Features

- **Code Generation**: Auto-generates Go structs and functions from i18n message files.
- **Type Safety**: Creates structs with named parameters (e.g., `ErrorAlreadyExistParam`) and functions with anonymous parameters (e.g., `NewConfirmAction`).
- **Flexible Output**: Supports custom output paths and package names, derived from the target DIR.
- **Multi-Language Support**: Tested with English (`en-US`), Simplified Chinese (`zh-CN`), and Khmer (`km-KH`).
- **Integration with go-i18n**: Generates `I18n*` functions that return `i18n.LocalizeConfig`, enabling direct use with `go-i18n`.
- **Clean Naming**: Converts message IDs (e.g., `ERROR_ALREADY_EXIST`) to PascalCase (e.g., `ErrorAlreadyExist`) using `strcase`.
- **Format and Imports**: Ensures generated code is well-formatted (`formatgo`) and includes needed imports (`syntaxgo_ast`).
- **Unicode Support**: Handles non-ASCII message IDs with custom naming functions.
- **Plural Forms**: Supports CLDR plural rules (one, other, few, many, zero, two) across multiple languages.

## Generated Code Structure

For each message type, goi18n generates different code patterns:

### 1. Messages with Named Parameters

**Input (YAML):**
```yaml
ERROR_NOT_EXIST: "{{ .what }} {{ .code }} does not exist"
```

**Generated Code:**
```go
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
    messageID, valuesMap := NewErrorNotExist(data)
    return &i18n.LocalizeConfig{
        MessageID:    messageID,
        TemplateData: valuesMap,
    }
}
```

### 2. Messages with Anonymous Parameters

**Input (YAML):**
```yaml
PLEASE_CONFIRM: "Please confirm {{ . }}"
```

**Generated Code:**
```go
func NewPleaseConfirm[Value comparable](value Value) (string, Value) {
    return "PLEASE_CONFIRM", value
}

func I18nPleaseConfirm[Value comparable](value Value) *i18n.LocalizeConfig {
    messageID, tempValue := NewPleaseConfirm(value)
    return &i18n.LocalizeConfig{
        MessageID:    messageID,
        TemplateData: tempValue,
    }
}
```

### 3. Messages without Parameters

**Input (YAML):**
```yaml
SUCCESS: "Success"
```

**Generated Code:**
```go
func NewSuccess() string {
    return "SUCCESS"
}

func I18nSuccess() *i18n.LocalizeConfig {
    messageID := NewSuccess()
    return &i18n.LocalizeConfig{
        MessageID: messageID,
    }
}
```

## Configuration Options

goi18n provides flexible configuration through the `Options` type:

```go
options := goi18n.NewOptions().
    WithOutputPath("internal/i18n/messages.go").
    WithPkgName("i18n").
    WithGenerateNewMessage(true).
    WithAllowNonAsciiRune(false)
```

Available options:

- **`WithOutputPath(path string)`**: Set the output file path (must end with `.go`)
- **`WithPkgName(name string)`**: Set the package name used in generated code
- **`WithOutputPathWithPkgName(path string)`**: Set output path and auto-infer package name from existing file/parent DIR
- **`WithGenerateNewMessage(bool)`**: Enable generation of `New*` functions that return `(messageID, templateData)` tuples
- **`WithAllowNonAsciiRune(bool)`**: Enable support on non-ASCII characters in message IDs
- **`WithUnicodeMessageName(func)`**: Customize naming function on Unicode message IDs
- **`WithUnicodeStructName(func)`**: Customize naming function on Unicode struct names
- **`WithUnicodeFieldName(func)`**: Customize naming function on Unicode field names

## Installation

```bash
go get github.com/yyle88/goi18n
```

## Usage

### Step 1: Prepare i18n Message Files

Create YAML files with each supported locale in a DIR (e.g., `i18n/`):

**`i18n/en-US.yaml`**:
```yaml
SAY_HELLO: "Hello, {{ .name }}!"
WELCOME: "Welcome to this app!"
SUCCESS: "Success"
PLEASE_CONFIRM: "Please confirm {{ . }}"
ERROR_NOT_EXIST: "{{ .what }} {{ .code }} does not exist"
ERROR_ALREADY_EXIST: "{{ .what }} {{ .code }} already exists"
```

**`i18n/zh-CN.yaml`**:
```yaml
SAY_HELLO: "你好，{{ .name }}！"
WELCOME: "欢迎使用此应用！"
SUCCESS: "成功"
PLEASE_CONFIRM: "请确认{{ . }}"
ERROR_NOT_EXIST: "{{ .what }} {{ .code }} 不存在"
ERROR_ALREADY_EXIST: "{{ .what }} {{ .code }} 已存在"
```

### Step 2: Generate Code

Use the `goi18n.Generate` function to process message files and generate Go code:

```go
package main

import (
    "github.com/nicksnyder/go-i18n/v2/i18n"
    "github.com/yyle88/goi18n"
    "github.com/yyle88/rese"
    "golang.org/x/text/language"
    "gopkg.in/yaml.v3"
)

func main() {
    bundle := i18n.NewBundle(language.AmericanEnglish)
    bundle.RegisterUnmarshalFunc("yaml", yaml.Unmarshal)
    messageFile := rese.P1(bundle.LoadMessageFile("i18n/en-US.yaml"))

    options := goi18n.NewOptions().WithOutputPathWithPkgName("output/message.go")
    goi18n.Generate([]*i18n.MessageFile{messageFile}, options)
}
```

This generates a file (`output/message.go`) with package `output`, containing structs (e.g., `ErrorAlreadyExistParam`) and functions (e.g., `NewSayHello`, `I18nSayHello`).

### Step 3: Use Generated Code

Import the generated package and use the functions to translate:

```go
package example1_test

import (
    "testing"
    "github.com/nicksnyder/go-i18n/v2/i18n"
    "github.com/stretchr/testify/require"
    "github.com/yyle88/goi18n/internal/examples/example1/internal/message1"
)

func TestI18nSayHello(t *testing.T) {
    bundle, _ := message1.LoadI18nFiles()
    localizer := i18n.NewLocalizer(bundle, "zh-CN")

    // Using I18nSayHello
    msg, err := localizer.Localize(message1.I18nSayHello(&message1.SayHelloParam{
        Name: "杨亦乐",
    }))
    require.NoError(t, err)
    require.Equal(t, "你好，杨亦乐！", msg)
}

func TestNewSayHello(t *testing.T) {
    bundle, _ := message1.LoadI18nFiles()
    localizer := i18n.NewLocalizer(bundle, "zh-CN")

    // Using NewSayHello
    messageID, msgValues := message1.NewSayHello(&message1.SayHelloParam{
        Name: "杨亦乐",
    })

    msg, err := localizer.Localize(&i18n.LocalizeConfig{
        MessageID:    messageID,
        TemplateData: msgValues,
    })
    require.NoError(t, err)
    require.Equal(t, "你好，杨亦乐！", msg)
}

func TestI18nErrorNotExist(t *testing.T) {
    bundle, _ := message1.LoadI18nFiles()
    localizer := i18n.NewLocalizer(bundle, "zh-CN")

    msg, err := localizer.Localize(message1.I18nErrorNotExist(&message1.ErrorNotExistParam{
        What: "数据库里",
        Code: "账号信息",
    }))
    require.NoError(t, err)
    require.Equal(t, "数据库里 账号信息 不存在", msg)
}
```

## Example

```
goi18n/
├── goi18n.go         # Main logic of the goi18n package
├── internal/
│   └── examples/
│       ├── example1/ # Example and test code with YAML-based internationalization
│       │   ├── example1_test.go
│       │   └── internal/
│       │       └── message1/
│       │           ├── en-US.yaml
│       │           ├── zh-CN.yaml
│       │           └── km-KH.yaml
│       ├── example2/ # Example and test code with JSON-based internationalization
│       │   ├── example2_test.go
│       │   └── internal/
│       │       └── message2/
│       │           ├── trans.en-US.json
│       │           └── trans.zh-CN.json
│       └── example3/ # Example and test code with 【chinese-first】 internationalization
│           ├── example3_test.go
│           └── internal/
│               └── message3/
│                   ├── msg.en-US.yaml
│                   └── msg.zh-CN.yaml
```


- See [generate logic in example1](internal/examples/example1/internal/message1/i18n.gen_test.go) and [usage examples in example1](internal/examples/example1/example1_test.go).
- See [generate logic in example2](internal/examples/example2/internal/message2/i18n.gen_test.go) and [usage examples in example2](internal/examples/example2/example2_test.go).
- See [generate logic in example3](internal/examples/example3/internal/message3/i18n.gen_test.go) and [usage examples in example3](internal/examples/example3/example3_test.go).

## Testing

The project includes tests in `internal/examples/example1/example1_test.go`, covering various use cases:

- Named parameters: `I18nSayHello`, `I18nErrorAlreadyExist`
- Anonymous parameters: `I18nPleaseConfirm`
- No parameters: `I18nWelcome`, `I18nSuccess`

<!-- TEMPLATE (EN) BEGIN: STANDARD PROJECT FOOTER -->
<!-- VERSION 2025-11-01 00:00:00.000000 +0000 UTC -->

## 📄 License

MIT License. See [LICENSE](LICENSE).

---

## 🤝 Contributing

Contributions are welcome! Report bugs, suggest features, and contribute code:

- 🐛 **Found a mistake?** Open an issue on GitHub with reproduction steps
- 💡 **Have a feature idea?** Create an issue to discuss the suggestion
- 📖 **Documentation confusing?** Report it so we can improve
- 🚀 **Need new features?** Share the use cases to help us understand requirements
- ⚡ **Performance issue?** Help us optimize through reporting slow operations
- 🔧 **Configuration problem?** Ask questions about complex setups
- 📢 **Follow project progress?** Watch the repo to get new releases and features
- 🌟 **Success stories?** Share how this package improved the workflow
- 💬 **Feedback?** We welcome suggestions and comments

---

## 🔧 Development

New code contributions, follow this process:

1. **Fork**: Fork the repo on GitHub (using the webpage UI).
2. **Clone**: Clone the forked project (`git clone https://github.com/yourname/repo-name.git`).
3. **Navigate**: Navigate to the cloned project (`cd repo-name`)
4. **Branch**: Create a feature branch (`git checkout -b feature/xxx`).
5. **Code**: Implement the changes with comprehensive tests
6. **Testing**: (Golang project) Ensure tests pass (`go test ./...`) and follow Go code style conventions
7. **Documentation**: Update documentation to support client-facing changes and use significant commit messages
8. **Stage**: Stage changes (`git add .`)
9. **Commit**: Commit changes (`git commit -m "Add feature xxx"`) ensuring backward compatible code
10. **Push**: Push to the branch (`git push origin feature/xxx`).
11. **PR**: Open a merge request on GitHub (on the GitHub webpage) with detailed description.

Please ensure tests pass and include relevant documentation updates.

---

## 🌟 Support

Welcome to contribute to this project via submitting merge requests and reporting issues.

**Project Support:**

- ⭐ **Give GitHub stars** if this project helps you
- 🤝 **Share with teammates** and (golang) programming friends
- 📝 **Write tech blogs** about development tools and workflows - we provide content writing support
- 🌟 **Join the ecosystem** - committed to supporting open source and the (golang) development scene

**Have Fun Coding with this package!** 🎉🎉🎉

<!-- TEMPLATE (EN) END: STANDARD PROJECT FOOTER -->

---

## GitHub Stars

[![starring](https://starchart.cc/yyle88/goi18n.svg?variant=adaptive)](https://starchart.cc/yyle88/goi18n)
