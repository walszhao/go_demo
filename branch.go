package main

import (
	"fmt"
	"io/ioutil"
)

/**
go语言 的条件不需要用括号
*/
func branch() {
	const fileName = "abc.txt"
	contents, err := ioutil.ReadFile(fileName)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("%s\n", contents)
	}
}

/**
go语言可以在条件语句里写代码  分号隔开  再接条件语句
如果在条件语句定义变量  那么变量的作用域为条件语句block 而不是函数
*/
func branchCondition() {

	const fileName = "abc.txt"
	if contents, err := ioutil.ReadFile(fileName); err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("%s\n", contents)
	}
}

/**
switch  语句
go 语言的case 块不需要break 自动默认break
*/
func grade(sore int) string {
	g := ""
	switch {
	case sore < 0 || sore > 100:
		panic(fmt.Sprintf("Wrong sore: %d", sore))
	case sore <= 60:
		g = "F"
	case sore <= 80:
		g = "D"
	case sore <= 90:
		g = "B"
	case sore <= 100:
		g = "A"
	}
	return g

}

func main() {
	//branch()
	//branchCondition()

	fmt.Println(
		grade(55),
		//grade(-1),
		grade(99))

}
