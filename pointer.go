package main

import "fmt"

func Pointer() {
	// 练习 1：指针基础
	age := 20
	p := &age
	fmt.Println(age, p, *p)
	*p = 25
	fmt.Println(age)

	// 练习 2：指针函数参数
	num := 10
	modify(&num)
	fmt.Println(num)

	// 练习 3：值传递对比
	x := 10
	changeValue(x)
	fmt.Println(x)
	changePointer(&x)
	fmt.Println(x)

	// 练习 4：指针零值与 new
	var p1 *int
	fmt.Println(p1)
	fmt.Println(p1 == nil)
	p2 := new(int)
	*p2 = 50
	fmt.Println(*p2)

	// 练习 5：综合挑战
	y := 10
	z := 20
	swap(&y, &z)
	fmt.Println(y, z)
}

func modify(num *int) {
	*num = 100
}

func changeValue(num int) {
	// 修改 num 为 100
	num = 100
}

func changePointer(num *int) {
	// 修改原变量为 100
	*num = 100
}

func swap(a *int, b *int) {
	*a, *b = *b, *a
}
