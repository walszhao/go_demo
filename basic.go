package main

import (
	"fmt"
	"math"
	"math/cmplx"
)

/**
go语言引入了 complex 类型  用来表示复数(分为实部和虚部)
*/

/*
函数外定义变量  不可用 :=
不是全局变量  是包内部变量
*/
var (
	aa = 123
	ss = "oooo"
)

/**
变量定义
var  变量名  变量类型
*/
func variableZeroValue() {
	var a int
	var s string
	fmt.Printf("%d %q \n", a, s)
}

/**
变量定义
var  变量名，变量名  变量类型
可同时定义同类型的变量
*/
func variableInitialValue() {
	var a, b int = 3, 4
	var s string = "kkkkk"
	fmt.Printf("%d  %d  %q \n", a, b, s)
}

/**
变量定义
var  变量名，变量名  = value01, value02
Go语言可以根据初始化赋值情况进行类型推断  前提是 不定义变量类型
*/
func variableTypeDefine() {
	var a, b, c, s = 3, 4, false, "kkk"

	fmt.Println(a, b, c, s)
}

/**
变量定义
 变量名，变量名  := value01, value02
  省略 var 用:= 进行变量定义和赋值
  在对变量重新赋值是 不需要 :
*/
func variableShorter() {
	a, b, c, s := 3, 4, false, "kkk"
	a = 7
	fmt.Println(a, b, c, s)
}

func euler() {
	c := 3 + 4i
	fmt.Println(cmplx.Abs(c))
}

/**
go语言在类型转换时不会进行隐式转换  必须进行显示的强转
需要注意的是  写法上和java 强转不同
int(math.Sqrt(float64(a*a + b*b)))
*/
func triangle() {
	var a, b = 3, 4
	var c int
	c = int(math.Sqrt(float64(a*a + b*b)))
	fmt.Println(c)
}

/**
常量的定义 用关键字 const
可以定义 类型  也可以不用定义类型
也就是说  在对变量进行运算时  仅仅是进行了字符替换
不定义类型的情况下  a b 可以做int 也可以做 float
*/
func contants() {
	const (
		a, b = 1, 2
		s    = "111"
	)
	var c int
	// triangle 函数中  必须对 a*a + b*b进行强制类型转换
	// 这里定义 a b 为常量 不定义类型 可以不进行强转
	c = int(math.Sqrt(a*a + b*b))

	fmt.Println(s, c)
}

/**
go语言没有特定的枚举类型关键字
用常量的来定义
*/
func emuns() {
	/*普通 枚举类型定义
	const (
		cpp    = 0
		java   = 1
		golang = 2
	)*/
	/** 自增值枚举类型
	iota 表达式的作用是  变量自增
	*/
	const (
		cpp = iota
		java
		golang
	)

	//用iota 表示 b kb mb gb tb pb
	const (
		b = 1 << (10 * iota)
		kb
		mb
		gb
		tb
		pb
	)

	fmt.Println(cpp, java, golang)
	fmt.Println(b, kb, mb, gb, tb, pb)
}

func main() {
	/*fmt.Println("aaaaaa")
	variableZeroValue()
	variableInitialValue()
	variableTypeDefine()
	variableShorter()
	fmt.Println(aa, ss)*/

	euler()
	triangle()
	contants()
	emuns()
}
