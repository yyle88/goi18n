package goi18n

import (
	"path/filepath"
	"strings"

	"github.com/iancoleman/strcase"
	"github.com/yyle88/must"
)

type Options struct {
	OutputPath string
	PkgName    string
}

func NewOptions() *Options {
	return &Options{}
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
