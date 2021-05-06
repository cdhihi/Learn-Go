package main

import "fmt"

type num struct {
	a int
	b int
	sym string
}


func (num *num) mor_than( ret1, ret2 interface{}) interface{}  {
	if num.a > num.b {
		return ret1
	}else {
		return ret2
	}

}

func (num *num) less_than( ret1, ret2 interface{}) interface{}  {
	if num.a < num.b {
		return ret1
	}else {
		return ret2
	}

}

func (num *num) not_equal( ret1, ret2 interface{}) interface{}  {
	if num.a != num.b {
		return ret1
	}else {
		return ret2
	}

}



func (num *num) equal( ret1, ret2 interface{}) interface{}  {
	if num.a == num.b {
		return ret1
	}else {
		return ret2
	}
}





func (num *num) calculation(ret1, ret2 interface{}) interface{} {

	switch num.sym {
	case ">":

		return num.mor_than(ret1 , ret2)

	case "<":

		return num.less_than(ret1 , ret2)
	case "=":

		return num.equal(ret1 , ret2)
	case "!=":

		return num.not_equal(ret1 , ret2)
	default:
		return  nil
	}


}


func main()  {
   a := num{
   	a : 2,
   	b : 2,
   	sym : "=",
   }
	c := a.calculation( "验证成功","验证失败")

	fmt.Println(c)
}