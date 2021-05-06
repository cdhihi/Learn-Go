package main

import "fmt"

//用来比较的结构体
type num struct {
	a   int
	b   int
	sym string
}

//大于比较方法
func (num *num) mor_than(ret1, ret2 interface{}) interface{} {
	if num.a > num.b {
		return ret1
	} else {
		return ret2
	}

}

//小于比较方法
func (num *num) less_than(ret1, ret2 interface{}) interface{} {
	if num.a < num.b {
		return ret1
	} else {
		return ret2
	}

}

//不等于比较方法
func (num *num) not_equal(ret1, ret2 interface{}) interface{} {
	if num.a != num.b {
		return ret1
	} else {
		return ret2
	}

}

//等于比较方法
func (num *num) equal(ret1, ret2 interface{}) interface{} {
	if num.a == num.b {
		return ret1
	} else {
		return ret2
	}
}

func (num *num) calculation(ret1, ret2 interface{}) interface{} {

	//不用的运算符用不同的方法
	switch num.sym {
	case ">":

		return num.mor_than(ret1, ret2)

	case "<":

		return num.less_than(ret1, ret2)
	case "=":

		return num.equal(ret1, ret2)
	case "!=":

		return num.not_equal(ret1, ret2)
	default:
		return nil
	}

}

func main() {
	//写一个简单的三元运算符

	a := num{
		a:   2,
		b:   2,
		sym: "=",
	}

	//开始计算
	c := a.calculation("验证成功", "验证失败")

	fmt.Println(c)
}
