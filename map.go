package main

import "fmt"

func Map() {
	// 练习 1：Map 基础 CRUD
	scores := map[string]int{
    "张三": 90,
    "李四": 85,
    "王五": 95,
	}
	fmt.Println(scores["张三"])
	scores["李四"] = 100
	scores["赵六"] = 88
	delete(scores, "王五")
	fmt.Println(scores, len(scores))

	// 练习 2：判断 Key 是否存在
	scores2 := map[string]int {
		"张三": 90,
		"李四": 0,
	}
	names := []string{"张三", "李四", "王五"}
	for _, name := range names {
    score, ok := scores2[name]

    fmt.Printf(
        "姓名：%s, 成绩：%d, 是否存在：%t\n",
        name, score, ok,
    )
	}

	// 练习 3：Map 遍历
	prices := map[string]float64{
		"苹果": 5.5,
		"香蕉": 3.2,
		"橘子": 4.8,		
	}
	sum := 0.0
	for fruit, price := range prices {
		fmt.Printf("水果：%s, 价格：%.2f\n", fruit, price)
		sum += price
	}
	fmt.Printf("总价：%.2f\n", sum)

	// 练习 4：统计字符出现次数
	text := "hello golang"
	map_text := make(map[rune]int)
	for _, char := range text {
		map_text[char]++
	}
	for char, count := range map_text {
		fmt.Printf("%q: %d\n", char, count)
	}

	// 练习 5：函数传递
	scores3 := map[string]int{
    "张三": 90,
    "李四": 85,
    "王五": 95,
	}
	updateScore(scores3)
	fmt.Println(scores3)
}

func updateScore(scores map[string]int) {
	scores["张三"] = 100
	scores["王五"] = 95
}