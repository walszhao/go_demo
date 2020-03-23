package main

import "fmt"

func main() {

	//切片   是对一个数组的视图
	//切片 是址传递

	arr := [...]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}

	fmt.Println("arr[2:6]  = ", arr[2:6])
	fmt.Println("arr[:6]  = ", arr[:6])
	fmt.Println("arr[2:]  = ", arr[2:])
	fmt.Println("arr[:]  = ", arr[:])

	// reslice   可以在切片上再进行切片  得到的切片依然是原数组的切片
	s1 := arr[:5]
	s1 = s1[2:]
	fmt.Println(s1)

	//这里涉及一个 切片扩展
	// arrExtending1 的值为  3,4,5
	// arrExtending2 的值在 arrExtending1 的基础上 结合 arr 进行扩展
	// 扩展不可超越 arr 的长度
	// 只可向后 扩展  不可向前扩展
	// 详细原理  参考  /picture/sliceExtend01.png  /picture/sliceExtend02.png
	arrExtending1 := arr[3:6]
	arrExtending2 := arrExtending1[2:5]
	fmt.Println(arrExtending1, arrExtending2)

	s2 := append(s1, 10)
	s3 := append(s2, 11)
	s4 := append(s3, 12)

	fmt.Println("s1,s2,s3,s4 = ", s1, s2, s3, s4)
	fmt.Println("arr = ", arr)

}
