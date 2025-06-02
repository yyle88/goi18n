# Overview
This is an easy way to write code for translations.  
By using `{{.PluralCount}}` in the text, you can pick the right plural form and include the number directly with `PluralCount: num`.

```go
localizer.MustLocalize(&i18n.LocalizeConfig{
    DefaultMessage: catsMessage,
    PluralCount:    num,
    // TemplateData: map[string]interface{}{
    //    "PluralCount": num,
    // }, // Since PluralCount is a special name, you can skip this part and still get the same result.
})
```

While, I think this approach might be too complex.  
As I improve my coding skills, I might find it more useful. But as a beginner, I was really surprised when I first saw this method.

---

# 知识总结
这是种比较省略的写法
通过配置 {{.PluralCount}} 这个特殊的名字为翻译的内容，就能够借由 `PluralCount: num` 既选中复数模版还把复数的参数传到模板里

```
localizer.MustLocalize(&i18n.LocalizeConfig{
    DefaultMessage: catsMessage,
    PluralCount:    num,
    // TemplateData: map[string]interface{}{
    //    "PluralCount": num,
    // }, // 因为 PluralCount 已经是特殊的名字，这里就能省掉这块，达到同样的效果
}
```

但我认为这种写法没卵用，这属于是过度设计的结果
当然随着以后的熟练，我或许也会常用这种写法，但目前作为新手我只能说，初见这个方法时我感到大受震撼。
