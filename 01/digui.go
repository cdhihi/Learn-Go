package main

import (
	"fmt"
	"strconv"
)

var (
	s = "1,1,"
)

func Fibonacci(a, b, num, strnum int) {

	a, b = b, (a + b)
	s += strconv.Itoa(b) + ","

	strnum++
	if strnum < num {
		Fibonacci(a, b, num, strnum)
	}
}

var MIN = 50

func main() {
	//下面做了一个简单的斐波那契数列算法

	Fibonacci(1, 1, 10, 0)
	fmt.Println(s)

	arr := fibarray(MIN)
	fmt.Print(arr)

}

func fibarray(term int) []int {
	//创建一个数组,以数组的形式来存放数输出的内容格式
	farr := make([]int, term)
	fmt.Println(farr)
	farr[0], farr[1] = 1, 1
	for i:= 2; i < term; i++ {
		farr[i] = farr[i-1] + farr[i-2]
	}
	return farr
}
