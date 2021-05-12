# Google Tranlsate CLI

用 Go 编写的简单 CLI 工具，用于调用 Google Translate API 进行翻译。

## 安装

```shell
go get github.com/onlyice/google-translate-cli
```

会在 `$GOPATH/bin` 或者 `$HOME/go/bin` 下安装 `google-translate-cli` 二进制文件。请把这个路径加入 `PATH` 环境变量中，方便使用。

## 使用

```shell
> google-translate-api --text "hello world"

# 输出帮助
> google-translate-api -h
```

## 搭配 GoldenDict 使用

本工具提供了 `--html` 选项用于生成 HTML；而 GoldenDict 的词典来源可以是某个程序的输出结果。因此你可以搭配 GoldenDict 来翻译。具体配置方法你摸索一下就懂了。