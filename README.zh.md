[![GitHub Workflow Status (branch)](https://img.shields.io/github/actions/workflow/status/yyle88/goi18n/release.yml?branch=main&label=BUILD)](https://github.com/yyle88/goi18n/actions/workflows/release.yml?query=branch%3Amain)
[![GoDoc](https://pkg.go.dev/badge/github.com/yyle88/goi18n)](https://pkg.go.dev/github.com/yyle88/goi18n)
[![Coverage Status](https://img.shields.io/coveralls/github/yyle88/goi18n/main.svg)](https://coveralls.io/github/yyle88/goi18n?branch=main)
[![Supported Go Versions](https://img.shields.io/badge/Go-1.22--1.25-lightgrey.svg)](https://github.com/yyle88/goi18n)
[![GitHub Release](https://img.shields.io/github/release/yyle88/goi18n.svg)](https://github.com/yyle88/goi18n/releases)
[![Go Report Card](https://goreportcard.com/badge/github.com/yyle88/goi18n)](https://goreportcard.com/report/github.com/yyle88/goi18n)

# goi18n

`goi18n` 是一个 Go 包和代码生成工具集，通过使用泛型参数替换 `map[string]interface{}`，使 `go-i18n` 的使用更加简洁。

## 概述

`goi18n` 简化了 Go 应用程序中的国际化（i18n）开发。它能够处理 i18n 消息文件（如 YAML），并生成类型安全的 Go 代码，包括消息处理的结构体和函数。生成的代码与 `go-i18n` 包集成，支持多种语言的高效、安全翻译渲染。

<!-- TEMPLATE (ZH) BEGIN: LANGUAGE NAVIGATION -->
## 英文文档

[ENGLISH README](README.md)
<!-- TEMPLATE (ZH) END: LANGUAGE NAVIGATION -->

## 为什么使用 goi18n？

在使用优秀的 [go-i18n](https://github.com/nicksnyder/go-i18n) 包时，你经常需要编写重复的代码：

```go
// 传统方式 - 冗长且容易出错
localizer.Localize(&i18n.LocalizeConfig{
    MessageID: "ERROR_NOT_EXIST",
    TemplateData: map[string]interface{}{
        "what": "User",
        "code": "12345",
    },
})
```

这种方式存在多个问题：
- **缺乏类型安全**：容易在消息 ID 和参数名称中出现拼写错误
- **运行时错误**：错误在运行时才能发现，而不是编译时
- **IDE 支持有限**：无法对消息 ID 和参数进行自动补全
- **维护负担**：难以追踪每个消息需要哪些参数

**goi18n 解决了这些问题**，通过生成类型安全的代码：

```go
// 使用 goi18n - 类型安全且简洁
message1.I18nErrorNotExist(&message1.ErrorNotExistParam{
    What: "User",
    Code: "12345",
})
```

优势：
- ✅ **编译时安全**：在部署前捕获错误
- ✅ **IDE 自动补全**：完整的 IntelliSense 支持
- ✅ **重构友好**：可以自信地重命名参数
- ✅ **自文档化**：生成的结构体显示所需参数
- ✅ **零运行时开销**：性能与手写代码相同

## 功能特性

- **代码生成**：从 i18n 消息文件自动生成 Go 结构体和函数。
- **类型安全**：为命名参数生成结构体（如 `ErrorAlreadyExistParam`），为匿名参数生成函数（如 `NewConfirmAction`）。
- **灵活输出**：支持自定义输出路径和包名，从目标目录自动推导。
- **多语言支持**：已测试支持英语（`en-US`）、简体中文（`zh-CN`）和高棉语（`km-KH`）。
- **与 go-i18n 集成**：生成返回 `i18n.LocalizeConfig` 的 `I18n*` 函数，直接用于 `go-i18n`。
- **规范命名**：将消息 ID（如 `ERROR_ALREADY_EXIST`）转换为 PascalCase（如 `ErrorAlreadyExist`），使用 `strcase`。
- **代码格式化**：使用 `formatgo` 确保代码格式规范，通过 `syntaxgo_ast` 自动添加必要导入。
- **Unicode 支持**：使用自定义命名函数处理非 ASCII 消息 ID。
- **复数形式**：支持多种语言的 CLDR 复数规则（one、other、few、many、zero、two）。

## 生成代码结构

针对每种消息类型，goi18n 生成不同的代码模式：

### 1. 带命名参数的消息

**输入（YAML）：**
```yaml
ERROR_NOT_EXIST: "{{ .what }} {{ .code }} does not exist"
```

**生成的代码：**
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

### 2. 带匿名参数的消息

**输入（YAML）：**
```yaml
PLEASE_CONFIRM: "Please confirm {{ . }}"
```

**生成的代码：**
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

### 3. 无参数的消息

**输入（YAML）：**
```yaml
SUCCESS: "Success"
```

**生成的代码：**
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

## 配置选项

goi18n 通过 `Options` 类型提供灵活的配置：

```go
options := goi18n.NewOptions().
    WithOutputPath("internal/i18n/messages.go").
    WithPkgName("i18n").
    WithGenerateNewMessage(true).
    WithAllowNonAsciiRune(false)
```

可用选项：

- **`WithOutputPath(path string)`**：设置输出文件路径（必须以 `.go` 结尾）
- **`WithPkgName(name string)`**：设置生成代码中使用的包名
- **`WithOutputPathWithPkgName(path string)`**：设置输出路径并从现有文件/父 DIR 自动推断包名
- **`WithGenerateNewMessage(bool)`**：启用生成返回 `(messageID, templateData)` 元组的 `New*` 函数
- **`WithAllowNonAsciiRune(bool)`**：启用对消息 ID 中非 ASCII 字符的支持
- **`WithUnicodeMessageName(func)`**：自定义 Unicode 消息 ID 的命名函数
- **`WithUnicodeStructName(func)`**：自定义 Unicode 结构体名称的命名函数
- **`WithUnicodeFieldName(func)`**：自定义 Unicode 字段名称的命名函数

## 安装

```bash
go get github.com/yyle88/goi18n
```

## 使用方法

### 步骤 1：准备 i18n 消息文件

在目录（如 `i18n/`）中为每种支持的语言创建 YAML 文件：

**`i18n/en-US.yaml`**：
```yaml
SAY_HELLO: "Hello, {{ .name }}!"
WELCOME: "Welcome to this app!"
SUCCESS: "Success"
PLEASE_CONFIRM: "Please confirm {{ . }}"
ERROR_NOT_EXIST: "{{ .what }} {{ .code }} does not exist"
ERROR_ALREADY_EXIST: "{{ .what }} {{ .code }} already exists"
```

**`i18n/zh-CN.yaml`**：
```yaml
SAY_HELLO: "你好，{{ .name }}！"
WELCOME: "欢迎使用此应用！"
SUCCESS: "成功"
PLEASE_CONFIRM: "请确认{{ . }}"
ERROR_NOT_EXIST: "{{ .what }} {{ .code }} 不存在"
ERROR_ALREADY_EXIST: "{{ .what }} {{ .code }} 已存在"
```

### 步骤 2：生成代码

使用 `goi18n.Generate` 函数处理消息文件并生成 Go 代码：

```go
package main

import (
    "github.com/nicksnyder/go-i18n/v2/i18n"
    "github.com/yyle88/goi18n"
	"github.com/yyle88/rese"
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

这将生成一个文件（`output/message.go`），包名为 `output`，包含结构体（如 `ErrorAlreadyExistParam`）和函数（如 `NewSayHello`、`I18nSayHello`）。

### 步骤 3：使用生成的代码

导入生成的包并使用函数进行翻译：

```go
package main

import (
    "fmt"
    "github.com/nicksnyder/go-i18n/v2/i18n"
    "github.com/yyle88/goi18n/internal/examples/example1/example1generate/example1message"
    "golang.org/x/text/language"
    "gopkg.in/yaml.v3"
)

func main() {
    bundle := i18n.NewBundle(language.AmericanEnglish)
    bundle.RegisterUnmarshalFunc("yaml", yaml.Unmarshal)
    bundle.MustLoadMessageFile("i18n/zh-CN.yaml")

    localizer := i18n.NewLocalizer(bundle, "zh-CN")

    // 使用 I18nSayHello
    config := example1message.I18nSayHello(&example1message.SayHelloParam{Name: "杨亦乐"})
    msg, err := localizer.Localize(config)
    if err != nil {
        panic(err)
    }
    fmt.Println(msg) // 输出：你好，杨亦乐！

    // 使用 NewSayHello
    messageID, data := example1message.NewSayHello(&example1message.SayHelloParam{Name: "杨亦乐"})
    msg, err = localizer.Localize(&i18n.LocalizeConfig{
        MessageID:    messageID,
        TemplateData: data,
    })
    fmt.Println(msg) // 输出：你好，杨亦乐！
}
```

## 示例

```
goi18n/
├── goi18n.go         # 核心逻辑
├── internal/
│   └── examples/
│       ├── example1/ # 展示如何读取 yaml 配置获取翻译信息
│       │   ├── example1_test.go
│       │   └── internal/
│       │       └── message1/
│       │           ├── en-US.yaml
│       │           ├── zh-CN.yaml
│       │           └── km-KH.yaml
│       ├── example2/ # 展示如何读取 json 配置获取翻译信息
│       │   ├── example2_test.go
│       │   └── internal/
│       │       └── message2/
│       │           ├── trans.en-US.json
│       │           └── trans.zh-CN.json
│       └── example3/ # 展示如何使用【中文优先】的国际化配置
│           ├── example3_test.go
│           └── internal/
│               └── message3/
│                   ├── msg.en-US.yaml
│                   └── msg.zh-CN.yaml
```

- See [生成逻辑1](internal/examples/example1/internal/message1/i18n.gen_test.go) and [使用示例1](internal/examples/example1/example1_test.go).
- See [生成逻辑2](internal/examples/example2/internal/message2/i18n.gen_test.go) and [使用示例2](internal/examples/example2/example2_test.go).
- See [生成逻辑3](internal/examples/example3/internal/message3/i18n.gen_test.go) and [使用示例3](internal/examples/example3/example3_test.go).

## 测试

项目在 `internal/examples/example1/example1_test.go` 中包含测试用例，覆盖以下场景：

- 命名参数：`I18nSayHello`、`I18nErrorAlreadyExist`
- 匿名参数：`I18nPleaseConfirm`
- 无参数：`I18nWelcome`、`I18nSuccess`

<!-- TEMPLATE (ZH) BEGIN: STANDARD PROJECT FOOTER -->
<!-- VERSION 2025-11-01 00:00:00.000000 +0000 UTC -->

## 📄 许可证类型

MIT 许可证。详情请参阅 [LICENSE](LICENSE)。

---

## 🤝 贡献新代码

欢迎贡献代码！报告错误、建议功能和贡献代码：

- 🐛 **发现错误？** 在 GitHub 上提出问题并附上复现步骤
- 💡 **有功能想法？** 创建问题讨论建议
- 📖 **文档令人困惑？** 报告它以便我们改进
- 🚀 **需要新功能？** 分享使用场景以帮助我们理解需求
- ⚡ **性能问题？** 通过报告慢速操作帮助我们优化
- 🔧 **配置问题？** 询问关于复杂设置的问题
- 📢 **关注项目进展？** Watch 仓库以获取新版本和功能
- 🌟 **成功案例？** 分享这个包如何改进工作流程
- 💬 **反馈？** 我们欢迎建议和评论

---

## 🔧 开发流程

新代码贡献，请遵循以下流程：

1. **Fork**：在 GitHub 上 Fork 仓库（使用网页界面）。
2. **Clone**：克隆 Forked 项目（`git clone https://github.com/yourname/repo-name.git`）。
3. **Navigate**：进入克隆的项目（`cd repo-name`）
4. **Branch**：创建功能分支（`git checkout -b feature/xxx`）。
5. **Code**：实现更改并编写全面的测试
6. **Testing**：（Golang 项目）确保测试通过（`go test ./...`）并遵循 Go 代码风格约定
7. **Documentation**：更新文档以支持面向客户端的更改，并使用有意义的提交消息
8. **Stage**：暂存更改（`git add .`）
9. **Commit**：提交更改（`git commit -m "Add feature xxx"`）确保向后兼容的代码
10. **Push**：推送到分支（`git push origin feature/xxx`）。
11. **PR**：在 GitHub 上打开合并请求（在 GitHub 网页上）并附上详细描述。

请确保测试通过并包含相关文档更新。

---

## 🌟 贡献与支持

欢迎通过提交合并请求和报告问题为此项目做出贡献。

**项目支持：**

- ⭐ **给 GitHub 点星** 如果这个项目对你有帮助
- 🤝 **与团队成员分享** 和（golang）编程朋友
- 📝 **撰写技术博客** 关于开发工具和工作流程 - 我们提供内容写作支持
- 🌟 **加入生态系统** - 致力于支持开源和（golang）开发场景

**祝编程愉快！** 🎉🎉🎉

<!-- TEMPLATE (ZH) END: STANDARD PROJECT FOOTER -->

---

## GitHub Stars

[![starring](https://starchart.cc/yyle88/goi18n.svg?variant=adaptive)](https://starchart.cc/yyle88/goi18n)
