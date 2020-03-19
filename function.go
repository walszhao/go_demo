package main

import "fmt"

/**
go语言的函数可以有多个返回值
如果只需接收一个返回值  另外一个用_代替
*/
func eval(a, b int, op string) (int, error) {

	switch op {
	case "+":
		return a + b, nil
	case "-":
		return a - b, nil
	case "*":
		return a * b, nil
	case "/":
		return a / b, nil
	default:
		return 0, fmt.Errorf("error op:%s", op)
	}

}

/**
go语言支持像Java 一样的 可变参数列表  传入的参数 相当于一个数组
*/
func sumArgs(numbers ...int) int {
	sum := 0
	for i := range numbers {
		sum += numbers[i]
	}
	return sum

}

/**
go语言中 参数传递是值传递
调用方拷贝参数的副本交给函数执行  不会改变调用方的值
如果需要改变调用方的值  则需要将参数的指针传递给函数
*/
func swapValue(a, b int) {
	a, b = b, a
}

func swapAddress(a, b *int) {
	*a, *b = *b, *a
}

func main() {
	//p, error := eval(3, 4, "*")

	/*p, _ := eval(3, 4, "+")
	fmt.Println(p)*/

	if p, error := eval(3, 4, "+"); error != nil {
		fmt.Println("Error:", error)
	} else {
		fmt.Println(p)
	}

	fmt.Println(sumArgs(1, 2, 3, 4, 5, 6, 7, 8, 9, 10))

	//值传递  没有交换
	a, b := 3, 4
	swapValue(a, b)
	fmt.Println(a, b)

	//地址传递  ab交换了值
	swapAddress(&a, &b)
	fmt.Println(a, b)
}
