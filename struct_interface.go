package main

import (
	"fmt"
	"math"
)

// 练习 1：Method
type Student struct {
	Name  string
	Score int
}

func (s Student) GetInfo() string {
	return fmt.Sprintf(
		"姓名：%s，成绩：%d",
		s.Name,
		s.Score,
	)
}

func (s *Student) SetScore(score int) {
	s.Score = score
}

func Struct() {
	s := Student{}
	s.SetScore(80)
	fmt.Println(s.Score)
}

// 练习 2：Interface
type Shape interface {
	Area() float64
}

type Rectangle struct {
	Width  float64
	Height float64
}

func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}

type Circle struct {
	Radius float64
}

func (c Circle) Area() float64 {
	return math.Pi * c.Radius * c.Radius
}

func PrintArea(s Shape) {
	fmt.Println(s.Area())
}

func Interface() {
	r := Rectangle{
		Width:  30,
		Height: 30,
	}
	c := Circle{
		Radius: 20,
	}
	PrintArea(r)
	PrintArea(c)
}

// 练习 3：理解接口隐式实现
type Runner interface {
	Run()
}

type Person struct{}

func (p *Person) Run() {
	fmt.Println("running")
}

// 实现的 Run 方法是指针接收者方法，所以只有 *Person 才有这个方法
// var r Runner = &Person{}
