# Overview

In Chinese, there’s no plural concept, so only "other" is used, and setting "One" / "Zero" or "Two" has no effect.

```go
catsMessage := &i18n.Message{
    ID:    "Cats",
    Zero:  "I have no cats", // doesn’t work: English only uses "one" (count=1) and "other" (all other counts)
    One:   "I have one cat.",
    Two:   "I have two cats.", // doesn’t work: English only uses "one" (count=1) and "other" (all other counts)
    Other: "I have {{.count}} cats.",
}
```

This is a hidden rule. I haven’t found where this rule is coded or how to turn it off (though I’m not sure if I really want to).
It’s just there by default, so we have to remember it.

---

# 知识总结

中文没有复数概念，所以只能用 "other" 类别，配置 "One"、"Zero" 或 "Two" 也不会生效。

```go
catsMessage := &i18n.Message{
    ID:    "Cats",
    Zero:  "我没有猫", // 没用的配置，中文只用 "other"
    One:   "我有一只猫", // 没用的配置，中文只用 "other"
    Two:   "我有两只猫", // 没用的配置，中文只用 "other"
    Other: "我有 {{.count}} 只猫",
}
```

这是个隐形规则，我甚至还没找到限制的代码位置，也不知道该如何关闭它（虽然我也不知真想关闭它，我就是想探索探索）。
但它确实是默认生效的，我们只能记住这个规则。
