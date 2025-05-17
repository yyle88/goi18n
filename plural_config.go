package goi18n

type numType interface {
	int | int8 | int16 | int32 | int64 | uint | uint8 | uint16 | uint32 | uint64 | float32 | float64
}

/*
PluralConfig 配置复数翻译的参数，通过复数的数量选择是 "单数翻译" 还是 "复数翻译"，接着再填模板参数完成语句

但是在使用中发现，中文没有单复数形式的翻译，即使配置页无法使用，原因：

CLDR 如何定义其他语言的复数规则？
不同语言的复数规则差异很大，例如：

语言	复数规则类别	示例（Count=0,1,2,5）
英语	one（1）, other（其他）	"1 apple", "3 apples"
俄语	one, few, many, other	1 яблоко, 2 яблока, 5 яблок
阿拉伯语	zero, one, two, few, many, other	6 种分类
中文	other	所有数量均用 other

因此结论是，很多语言是不支持单复数的，即使配置也是无效的，因此建议中文用户就不要分别配置单复数翻译啦
*/
type PluralConfig struct {
	PluralCount any
}

func NewPluralConfig[Num numType](pluralCount Num) *PluralConfig {
	return &PluralConfig{
		PluralCount: pluralCount,
	}
}
