package main

import (
	"strings"
	"unicode"
)

// 将文本转换为token列表
func tokenize(text string) []string {
	// 实现在单词边界上分割文本并删除标点符号
	return strings.FieldsFunc(text, func(r rune) bool {
		return !unicode.IsLetter(r) && !unicode.IsNumber(r)
	})
}

// 分析器
func analyze(text string) []string {
	tokens := tokenize(text)
	tokens = lowercaseFilter(tokens)
	tokens = stopwordFilter(tokens)
	tokens = stemmerFilter(tokens)
	return tokens
}
