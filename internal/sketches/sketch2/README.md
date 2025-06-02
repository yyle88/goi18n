# Overview

This is a simple way to write translation code.  
By using `{{.count}}` in the text, you can add the number to the translation.
The `PluralCount: num` picks the right plural form, and `TemplateData` with `"count": num` puts the number into the template.
In English, only "one" (for count=1) and "other" (for all other counts, including 0) are used.
So, setting "Zero" or "Two" in the code doesn’t work for English.
Similarly, in Chinese, there’s no plural concept, so only "other" is used, and setting "One" / "Zero" or "Two" has no effect.

```go
catsMessage := &i18n.Message{
    ID:    "Cats",
    Zero:  "I have no cats", // doesn’t work: English only uses "one" (count=1) and "other" (all other counts)
    One:   "I have one cat.",
    Two:   "I have two cats.", // doesn’t work: English only uses "one" (count=1) and "other" (all other counts)
    Other: "I have {{.count}} cats.",
}
```

This design feels a bit strange. When I first tried it, I spent hours wondering why the "One" translation for Chinese didn’t work, thinking my code was wrong.  
As a beginner, I found this approach confusing. I’d rather make my own mistakes than be limited by these hidden rules.
But now, we can’t change the rules, so we just have to remember them.

---

# 知识总结

虽然 i18n 允许配置 zero、one、two、few、many、other 这些复数类别，但在英文里只用 "one" (count=1) 和 "other" (其他所有计数)，其他配置没用。  
配置 "Zero" 或 "Two" 在英文里不起作用。  
这是 `Unicode CLDR plural rules` 的规则限制。  
同理，中文没有复数概念，所以只能用 "other" 类别，配置 "One"、"Zero" 或 "Two" 也不会生效。

```go
catsMessage := &i18n.Message{
    ID:    "Cats",
    Zero:  "我没有猫", // 没用的配置，中文只用 "other"
    One:   "我有一只猫", // 没用的配置，中文只用 "other"
    Two:   "我有两只猫", // 没用的配置，中文只用 "other"
    Other: "我有 {{.count}} 只猫",
}
```

这种设计真的很神奇。我一开始花了两个小时，甚至更久，研究为什么中文的 "One" 单数翻译不生效，还以为是自己的代码有问题。  
这设计挺让人困惑的，作为新手，我宁愿自己犯错，也不喜欢被这种隐形规则限制。  
但是现在，我们也没法去纠正它，只能记住这个规则。
