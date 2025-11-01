package goi18n

type numType interface {
	int | int8 | int16 | int32 | int64 | uint | uint8 | uint16 | uint32 | uint64 | float32 | float64
}

// PluralConfig provides count param for messages with one/other forms
// Uses PluralCount to select between "one" and "other" translations
// Note: Chinese and some languages use "other" for all counts
//
// CLDR defines different plural rules for languages:
// - English: one (1), other (others) - "1 apple", "3 apples"
// - Russian: one, few, many, other - 1 яблоко, 2 яблока, 5 яблок
// - Arabic: zero, one, two, few, many, other - 6 categories
// - Chinese: other - all counts use "other"
//
// PluralConfig 配置复数翻译的参数，通过数量选择使用 "one" 还是 "other" 翻译，然后填充模板参数完成语句
//
// 但是在使用中发现，中文没有单复数形式的翻译，即使配置也无法使用，原因：
//
// CLDR 如何定义其他语言的复数规则？
// 不同语言的复数规则差异很大，例如：
//
// 语言	复数规则类别	示例（Count=0,1,2,5）
// 英语	one（1）, other（其他）	"1 apple", "3 apples"
// 俄语	one, few, many, other	1 яблоко, 2 яблока, 5 яблок
// 阿拉伯语	zero, one, two, few, many, other	6 种分类
// 中文	other	所有数量均用 other
//
// 因此结论是，很多语言是不支持单复数的，即使配置也是无效的，因此建议中文用户就不要分别配置单复数翻译啦
type PluralConfig struct {
	PluralCount any
}

// NewPluralConfig creates PluralConfig with numeric count
// Accepts any numeric type through generic constraint
//
// NewPluralConfig 创建带数值计数的 PluralConfig
// 通过泛型约束接受任意数值类型
func NewPluralConfig[Num numType](pluralCount Num) *PluralConfig {
	return &PluralConfig{
		PluralCount: pluralCount,
	}
}
