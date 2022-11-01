package main

import (
	"strings"

	eng "github.com/kljensen/snowball/english"
)

// 小写过滤器
func lowercaseFilter(tokens []string) []string {
	r := make([]string, len(tokens))
	for i, token := range tokens {
		r[i] = strings.ToLower(token)
	}
	return r
}

var stopwords = map[string]struct{}{ // I wish Go had built-in sets.
	"a": {}, "and": {}, "be": {}, "have": {}, "i": {},
	"in": {}, "of": {}, "that": {}, "the": {}, "to": {},
}

// common-word过滤器
func stopwordFilter(tokens []string) []string {
	r := make([]string, 0, len(tokens))
	for _, token := range tokens {
		if _, ok := stopwords[token]; !ok {
			r = append(r, token)
		}
	}
	return r
}

// 还原单词基本形式
func stemmerFilter(tokens []string) []string {
	r := make([]string, len(tokens))
	for i, token := range tokens {
		r[i] = eng.Stem(token, false)
	}
	return r
}
