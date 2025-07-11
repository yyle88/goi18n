[![GitHub Workflow Status (branch)](https://img.shields.io/github/actions/workflow/status/yyle88/goi18n/release.yml?branch=main&label=BUILD)](https://github.com/yyle88/goi18n/actions/workflows/release.yml?query=branch%3Amain)
[![GoDoc](https://pkg.go.dev/badge/github.com/yyle88/goi18n)](https://pkg.go.dev/github.com/yyle88/goi18n)
[![Coverage Status](https://img.shields.io/coveralls/github/yyle88/goi18n/master.svg)](https://coveralls.io/github/yyle88/goi18n?branch=main)
![Supported Go Versions](https://img.shields.io/badge/Go-1.22%2C%201.23-lightgrey.svg)
[![GitHub Release](https://img.shields.io/github/release/yyle88/goi18n.svg)](https://github.com/yyle88/goi18n/releases)
[![Go Report Card](https://goreportcard.com/badge/github.com/yyle88/goi18n)](https://goreportcard.com/report/github.com/yyle88/goi18n)

# Goi18n

`goi18n` replace map[string]interface{} with generic parameters to make go-i18n more concise.

## Overview

`goi18n` is a Go library and code generation tool that simplifies internationalization (i18n) in Go applications. It processes i18n message files (e.g., YAML) and generates type-safe Go code, including structs and functions for message handling. The generated code integrates seamlessly with the `go-i18n` library, enabling efficient and safe translation rendering for multiple languages.

## CHINESE README

[中文说明](README.zh.md)

## Features

- **Code Generation**: Automatically generates Go structs and functions from i18n message files.
- **Type Safety**: Creates structs for named parameters (e.g., `ErrorAlreadyExistParam`) and functions for anonymous parameters (e.g., `NewConfirmAction`).
- **Flexible Output**: Supports custom output paths and package names, derived from the target directory.
- **Multi-Language Support**: Tested with English (`en-US`), Simplified Chinese (`zh-CN`), and Khmer (`km-KH`).
- **Integration with go-i18n**: Generates `I18n*` functions that return `i18n.LocalizeConfig` for direct use with `go-i18n`.
- **Clean Naming**: Converts message IDs (e.g., `ERROR_ALREADY_EXIST`) to PascalCase (e.g., `ErrorAlreadyExist`) using `strcase`.
- **Format and Imports**: Ensures generated code is well-formatted (`formatgo`) and includes necessary imports (`syntaxgo_ast`).

## Installation

```bash
go get github.com/yyle88/goi18n
```

## Usage

### Step 1: Prepare i18n Message Files

Create YAML files for each supported locale in a directory (e.g., `i18n/`):

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

Import the generated package and use the functions for translation:

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

    // Using I18nSayHello
    config := example1message.I18nSayHello(&example1message.SayHelloParam{Name: "杨亦乐"})
    msg, err := localizer.Localize(config)
    if err != nil {
        panic(err)
    }
    fmt.Println(msg) // Output: 你好，杨亦乐！

    // Using NewSayHello
	messageID, data := example1message.NewSayHello(&example1message.SayHelloParam{Name: "杨亦乐"})
    msg, err = localizer.Localize(&i18n.LocalizeConfig{
        MessageID:    messageID,
        TemplateData: data,
    })
    fmt.Println(msg) // Output: 你好，杨亦乐！
}
```

## Example

```
goi18n/
├── goi18n.go         # Main logic for the goi18n package
├── internal/
│   └── examples/
│       ├── example1/ # Example and test code for YAML-based internationalization
│       │   ├── example1_test.go
│       │   └── internal/
│       │       └── message1/
│       │           ├── en-US.yaml
│       │           ├── zh-CN.yaml
│       │           └── km-KH.yaml
│       ├── example2/ # Example and test code for JSON-based internationalization
│       │   ├── example2_test.go
│       │   └── internal/
│       │       └── message2/
│       │           ├── trans.en-US.json
│       │           └── trans.zh-CN.json
│       └── example3/ # Example and test code for 【chinese-first】 internationalization
│           ├── example3_test.go
│           └── internal/
│               └── message3/
│                   ├── msg.en-US.yaml
│                   └── msg.zh-CN.yaml
```


- See [generate logic for example1](internal/examples/example1/internal/message1/i18n.gen_test.go) and [usage examples for example1](internal/examples/example1/example1_test.go).
- See [generate logic for example2](internal/examples/example2/internal/message2/i18n.gen_test.go) and [usage examples for example2](internal/examples/example2/example2_test.go).
- See [generate logic for example3](internal/examples/example3/internal/message3/i18n.gen_test.go) and [usage examples for example3](internal/examples/example3/example3_test.go).

## Testing

The project includes tests in `internal/examples/example1/example1_test.go`, covering various use cases:

- Named parameters: `I18nSayHello`, `I18nErrorAlreadyExist`
- Anonymous parameters: `I18nPleaseConfirm`
- No parameters: `I18nWelcome`, `I18nSuccess`

---

## License

MIT License. See [LICENSE](LICENSE).

---

## Contributing

Contributions are welcome! To contribute:

1. Fork the repo on GitHub (using the webpage interface).
2. Clone the forked project (`git clone https://github.com/yourname/repo-name.git`).
3. Navigate to the cloned project (`cd repo-name`)
4. Create a feature branch (`git checkout -b feature/xxx`).
5. Stage changes (`git add .`)
6. Commit changes (`git commit -m "Add feature xxx"`).
7. Push to the branch (`git push origin feature/xxx`).
8. Open a pull request on GitHub (on the GitHub webpage).

Please ensure tests pass and include relevant documentation updates.

---

## Support

Welcome to contribute to this project by submitting pull requests and reporting issues.

If you find this package valuable, give me some stars on GitHub! Thank you!!!

**Thank you for your support!**

**Happy Coding with this package!** 🎉

Give me stars. Thank you!!!

---

## GitHub Stars

[![starring](https://starchart.cc/yyle88/goi18n.svg?variant=adaptive)](https://starchart.cc/yyle88/goi18n)
