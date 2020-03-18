package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

/**
for循环  实现一个十进制转二进制
*/
func convertToBin(n int) string {
	//循环  对2取模 得到的余数 倒叙相加
	var result string
	for ; n > 0; n /= 2 {
		lsb := n % 2
		result = strconv.Itoa(lsb) + result

	}
	return result
}

//循环的方式读取文件
func printFile(fileName string) {
	file, error := os.Open(fileName)
	if error != nil {
		panic(error)
	} else {
		scanner := bufio.NewScanner(file)
		//这里的 for省略起始条件和递增条件  只写终结条件
		//达到其他语言的 while 循环的效果
		//也可以什么条件都不写  是一个死循环
		for scanner.Scan() {
			fmt.Println(scanner.Text())
		}
	}
}

func main() {
	fmt.Println(convertToBin(3), convertToBin(13), convertToBin(53))

	printFile("abc.txt")
}
