package main

import (
	"fmt"
	"strconv"
)

func main()  {

	a := 0
	b := 1
	s := "打印开始："
	for i := 1;i < 20;i++{

		c := b + a
		a = b
		b = c
		s += strconv.Itoa(c) + ","
	}

	fmt.Println(s)


}