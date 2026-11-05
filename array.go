package main

import "fmt"

func Array() {
	// 练习 1：数组 CRUD
	arr := [5]int{10, 20, 30, 40, 50}
	fmt.Println(arr[1])
	arr[2] = 100
	for index, value := range arr {
		fmt.Printf("Index: %d, Value: %d\n", index, value)
	}
	fmt.Println(len(arr))

	// 练习 2：数组值传递
	arr2 := arr
	arr2[0] = 200
	fmt.Println("arr:", arr)
	fmt.Println("arr2:", arr2)

	// 数组基础
	arr3 := [5]int{1, 2, 3, 4, 5}
	arr[2] = 100
	for index, value := range arr3 {
		fmt.Printf("Index: %d, Value: %d\n", index, value)
	}
	fmt.Println(len(arr3))

	// 切片 CRUD
	slice := []int{10, 20, 30, 40, 50}
	fmt.Println(slice[1])
	slice[1] = 100
	slice = append(slice, 60)
	slice = append(slice[:2], slice[3:]...)
	fmt.Println(slice)

	// len 与 cap
	slice2 := make([]int, 3, 5)
	slice2 = append(slice2, 1, 2)
	fmt.Println("slice2:", len(slice2), cap(slice2))
	slice2 = append(slice2, 3)
	fmt.Println("slice2:", len(slice2), cap(slice2))

	// 底层数组共享
	arr4 := [5]int{10, 20, 30, 40, 50}
	a := arr4[1:4]
	b := arr4[2:5]
	a[1] = 15
	fmt.Println(arr4, a, b)

	// copy
	slice3 := []int{10, 20, 30}
	b1 := make([]int, 3)
	copy(b1, slice3)
	b1[0] = 100
	fmt.Println(slice3, b1)

	// 函数传递
	slice4 := []int{10, 20, 30}
	change1(slice4)
	fmt.Println(slice4)
	change2(slice4)
	fmt.Println(slice4)
}

func change1(slice []int) {
	slice[0] = 100
}

func change2(slice []int) {
	slice = append(slice, 100)
}