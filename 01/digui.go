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

func main() {
	//下面做了一个简单的斐波那契数列算法

	Fibonacci(1, 1, 10, 0)
	fmt.Println(s)

}
