package main

import "fmt"

func main() {

	var arr1 [2]int

	arr2 := [3]int{1, 2, 3}
	//不定义数组长度 用 ... 代替
	arr3 := [...]int{1, 2, 3, 4, 5, 6}

	//二维数组
	var grid [4][5]int
	fmt.Println(grid)

	//range 关键字可以同时循环 Key 和 Value
	for i, v := range arr3 {
		fmt.Println(i, v)
	}

	//数组是值类型 传入数据副本

	fmt.Println(arr1, arr2, arr3)

}
