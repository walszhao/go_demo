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

func main() {
	//branch()
	branchCondition()
}
