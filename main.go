package main

import (
	"github.com/atotto/clipboard"
	"serial-number/format"
)

func main() {
	config := format.Config{
		Headers:        []string{"##", "###", "####", "#####", "######", "#######"},
		ReplaceRegexps: []string{`^[\d.]+`},
		LevelsToString: format.DefaultLevelsToString,
	}
	// 读取剪切版内容
	text, _ := clipboard.ReadAll()
	// 添加序号
	newText := format.AddSerialNumber(text, config)
	// 写入到剪切版
	clipboard.WriteAll(newText)
}