package main

import "fmt"

func printSlice(s []int) {
	fmt.Printf("%v, length=%d, cap=%d\n", s, len(s), cap(s))
}

// 创建切面
func createSlice() {
	s1 := []int{1, 3, 5, 7, 9}
	printSlice(s1)

	// make 内建函数  指定长度
	s2 := make([]int, 16)
	printSlice(s2)

	// make 函数 指定长度和cap
	s3 := make([]int, 8, 32)
	printSlice(s3)
}

func copyingSlice() {
	s1 := []int{1, 3, 5, 7, 9}
	s2 := make([]int, 16)
	copy(s2, s1)
	printSlice(s2)
}

func deleteEle() {
	// 删除元素没有内建函数
	// 用append 进行拼接  漏掉要删除的那个
	// 比如要删除 5  就是前两个 append 后两个
	// 删除后cap 不变
	s1 := []int{1, 3, 5, 7, 9}

	// append 函数 第二个参数需要数组类型的  下面这种写法  表示切片里的每一个元素
	s1 = append(s1[:2], s1[3:]...)
	printSlice(s1)

}

func popingFromFront() {
	s1 := make([]int, 16)
	s1 = append(s1, 1, 3, 4, 5, 6)
	front := s1[0]
	s1 = s1[1:]
	fmt.Println(front)
	printSlice(s1)
}

func popingFromEnd() {
	s1 := make([]int, 16)
	s1 = append(s1, 1, 3, 4, 5, 6)
	end := s1[len(s1)-1]
	s1 = s1[:len(s1)-1]
	fmt.Println(end)
	printSlice(s1)
}

func main() {
	// 切片被定义为 nil 时  程序不会崩溃  默认加载  长度为0  cap 为0的切片
	//不断的添加元素时  长度依次增加  cap 成倍扩容
	/*var s []int
	for i := 0; i < 100; i++ {
		printSlice(s)
		s = append(s, 2*i+1)
	}
	fmt.Println(s)*/
	fmt.Println("create slice")
	createSlice()

	fmt.Println("copy slice")
	copyingSlice()

	fmt.Println("delete element")
	deleteEle()

	fmt.Println("poping from front")
	popingFromFront()

	fmt.Println("poping from end")
	popingFromEnd()

}
