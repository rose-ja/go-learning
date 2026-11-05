package main

import (
	"fmt"
	"strings"
)

func WordCount(text string) {
	// 统一转换为小写
	text = strings.ToLower(text)

	// 按空格切分单词
	words := strings.Fields(text)

	// 统计每个单词出现次数
	wordCount := make(map[string]int)
	for _, word := range words {
		wordCount[word]++
	}

	// 打印统计结果
	for word, count := range wordCount {
		fmt.Printf("%s: %d\n", word, count)
	}
}
