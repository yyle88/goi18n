# Overview

In this example, I wanted to try what happens if the `MessageID` is in Chinese and if the template parameter key name is in Chinese.  
It turns out this works, and it’s the same as using English.

```toml
[Cats]
Description = "Cats"
other = "我有 {{.count}} 只猫。"

["我有几只猫"]
Description = "Cats"
other = "我有 {{.猫的数量}} 只猫。"
```

```go
fmt.Println(localizer.MustLocalize(&i18n.LocalizeConfig{
    MessageID:   "Cats",
    PluralCount: num,
    TemplateData: map[string]interface{}{
        "count": num,
    },
}))
fmt.Println(localizer.MustLocalize(&i18n.LocalizeConfig{
    MessageID:   "我有几只猫",
    PluralCount: num,
    TemplateData: map[string]interface{}{
        "猫的数量": num,
    },
}))
```

This works well. It lets me do things like writing message configurations mainly in Chinese and then adding English translations.  
Since my team and customers mostly use Chinese, we usually focus on Chinese prompts and error messages first. We add English translations when the project is about 80% done.  
So, this experiment is invaluable.

---

# 知识总结

在这个例子里，我想试试假如 "MessageID" 是【中文】的会怎么样，和，假如模板参数的键名称为【中文】是怎么样的。
实践证明这也是能运行的，而且和配置成英文是等效的。

```toml
[Cats]
Description = "Cats"
other = "我有 {{.count}} 只猫。"

["我有几只猫"]
Description = "Cats"
other = "我有 {{.猫的数量}} 只猫。"
```

```go
fmt.Println(localizer.MustLocalize(&i18n.LocalizeConfig{
    MessageID:   "Cats",
    PluralCount: num,
    TemplateData: map[string]interface{}{
        "count": num,
    },
}))
fmt.Println(localizer.MustLocalize(&i18n.LocalizeConfig{
    MessageID:   "我有几只猫",
    PluralCount: num,
    TemplateData: map[string]interface{}{
        "猫的数量": num,
    },
}))
```

目前看来这确实是可以的，这样就能让我做很多事情，比如把消息的配置文件写成以中文为主体的的，再辅以英文翻译。
由于我的同事都是中文开发者，而我们的客户大多也是用中文的，我们通常开发思路就是先管好中文的提示语和报错信息，在项目开发至80%以上的时候再去翻译英语。
因此这样的探索还是很有必要的。
