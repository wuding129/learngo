package main

import "fmt"

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

func main() {
	fmt.Println("hello world")
	variableZeroValue()
	variableInitialVAlue()
	variableTypeDeduction()
	variableShorter()
	fmt.Println(aa,bb,ss)
}
