package main

import "fmt"

func main()  {

	//封装if简单三元运算符
	a,b := 1,2

	c := IF(a > b,"1","2").(interface{})

	fmt.Println(c)
}

func IF(b bool, ret1, ret2 interface{}) interface{} {

	if b {
		return ret1
	}

	return ret2

}
