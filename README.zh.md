# goi18n

`goi18n` 是一个 Go 库和代码生成工具，通过使用泛型参数替换 `map[string]interface{}`，使 `go-i18n` 的使用更加简洁。

## 概述

`goi18n` 简化了 Go 应用程序中的国际化（i18n）开发。它能够处理 i18n 消息文件（如 YAML），并生成类型安全的 Go 代码，包括消息处理的结构体和函数。生成的代码与 `go-i18n` 库无缝集成，支持多种语言的高效、安全翻译渲染。

## 英文文档

[English README](README.md)

## 功能特性

- **代码生成**：从 i18n 消息文件自动生成 Go 结构体和函数。
- **类型安全**：为命名参数生成结构体（如 `ErrorAlreadyExistParam`），为匿名参数生成函数（如 `NewConfirmAction`）。
- **灵活输出**：支持自定义输出路径和包名，从目标目录自动推导。
- **多语言支持**：已测试支持英语（`en-US`）、简体中文（`zh-CN`）和高棉语（`km-KH`）。
- **与 go-i18n 集成**：生成返回 `i18n.LocalizeConfig` 的 `I18n*` 函数，直接用于 `go-i18n`。
- **规范命名**：将消息 ID（如 `ERROR_ALREADY_EXIST`）转换为 PascalCase（如 `ErrorAlreadyExist`），使用 `strcase`。
- **代码格式化**：使用 `formatgo` 确保代码格式规范，通过 `syntaxgo_ast` 自动添加必要导入。

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

## 许可

项目采用 MIT 许可证，详情请参阅 [LICENSE](LICENSE)。

## 贡献新代码

非常欢迎贡献代码！贡献流程：

1. 在 GitHub 上 Fork 仓库 （通过网页界面操作）。
2. 克隆Forked项目 (`git clone https://github.com/yourname/repo-name.git`)。
3. 在克隆的项目里 (`cd repo-name`)
4. 创建功能分支（`git checkout -b feature/xxx`）。
5. 添加代码 (`git add .`)。
6. 提交更改（`git commit -m "添加功能 xxx"`）。
7. 推送分支（`git push origin feature/xxx`）。
8. 发起 Pull Request （通过网页界面操作）。

请确保测试通过并更新相关文档。

## 贡献与支持

欢迎通过提交 pull request 或报告问题来贡献此项目。

如果你觉得这个包对你有帮助，请在 GitHub 上给个 ⭐，感谢支持！！！

**感谢你的支持！**

**祝编程愉快！** 🎉

Give me stars. Thank you!!!
