package main

import "fmt"

func Variable() {
	// 练习一：变量声明
    name := "张三"
    age := 22
    height := 1.75
    isGraduated := false

    fmt.Println(name, age, height, isGraduated)

    // 练习二：零值
    var (
        zeroInt    int
        zeroString string
        zeroBool   bool
        zeroFloat  float64
    )

    fmt.Println(zeroInt, zeroString, zeroBool, zeroFloat)

    // 练习三：类型转换
    a := 10
    b := float64(a)

    price := 19.99
    c := int(price)

    fmt.Println(a, b, c)

    // 练习四：常量
    const PI = 3.14

    const (
        StatusPending = iota
        StatusActive
        StatusDisabled
    )

    fmt.Println(PI, StatusPending, StatusActive, StatusDisabled)
}