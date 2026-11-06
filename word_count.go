package main

import (
	"sort"
	"strings"
)

type WordFreq struct {
	Word  string
	Count int
}

func WordCount(text string) map[string]int {
	// 统一转换为小写
	text = strings.ToLower(text)

	// 按空格切分单词
	words := strings.Fields(text)

	// 统计每个单词出现次数
	wordCount := make(map[string]int)
	for _, word := range words {
		wordCount[word]++
	}

	return wordCount
}

func Top10(wordCount map[string]int) []WordFreq {
	// 1. 创建一个 WordFreq 类型的 Slice
	slice := make([]WordFreq, 0, len(wordCount))

	// 2. 遍历 wordCount，将每个键值对转换并添加到 Slice 中
	for word, count := range wordCount {
		slice = append(slice, WordFreq{Word: word, Count: count})
	}

	// 3. 使用 sort.Slice 按 Count 降序排序;频 次相同，按单词字典序升序排列
	sort.Slice(slice, func(i, j int) bool {
		if slice[i].Count == slice[j].Count {
			return slice[i].Word < slice[j].Word
		}

		return slice[i].Count > slice[j].Count
	})
	if len(slice) > 10 {
		slice = slice[:10]
	}

	return slice
}
