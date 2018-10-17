package main

import (
	"fmt"
	"math"
	"math/cmplx"
)

// 函数外定义变量,使用var,不能:=
//var aa = 3
//var ss = "kkk"
//var bb = true
// 或
var (
	aa = 3
	ss = "kkk"
	bb = true
)

// 定义变量
func variableZeroValue() {
	var a int
	var s string

	//fmt.Println(a, s)

	// 打印空字符串, %q会把引号打印出来
	fmt.Printf("%d %q\n", a, s)

}

// 初始化
func variableInitialVAlue() {
	var a, b int = 3, 4
	var s string = "abc"
	fmt.Println(a, b, s)

}

// 省略类型的初始化
func variableTypeDeduction() {
	var a, b, c, s = 3, 4, true, "def"
	fmt.Println(a, b, c, s)
}

// 变量简写
func variableShorter() {
	// 第一次定义变量可以用:=
	a, b, c, s := 3, 4, true, "def"
	b = 5
	fmt.Println(a, b, c, s)
}

// 内建变量类型
/*
bool, string
(u)int, (u)int8, (u)int16, (u)int32, (u)int64, uintptr ：‘指针’
byte：8位, rune：go的char类型，32位
float32, float64, complex64：复数，有实部和虚部, complex128

*/

// 验证欧拉公式e[i*PI]+1 = 0, e的i*PI次方加1为0，复数有实部和虚部
func euler() {
	//c := 3+4i
	//fmt.Println(cmplx.Abs(c))

	// 浮点数有误差，需要确定精度
	//fmt.Println(cmplx.Pow(math.E, 1i * math.Pi) + 1) 使用cmplx.Exp可直接写e的多少次方
	//fmt.Println(cmplx.Exp(1i * math.Pi) + 1)
	fmt.Printf("%.3f\n", cmplx.Exp(1i*math.Pi)+1)
}

// 类型转换，需要显示的强制转换
func triangle() {
	var a, b int = 3, 4
	var c int

	// todo 浮点数精度不准该怎么处理？
	c = int(math.Sqrt(float64(a * a + b * b)))
	fmt.Println(c)
}

func main() {
	fmt.Println("hello world")
	variableZeroValue()
	variableInitialVAlue()
	variableTypeDeduction()
	variableShorter()
	fmt.Println(aa, bb, ss)

	euler()
	triangle()
}
