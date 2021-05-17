package main

import (
	"Learn-Go/04/pay"
	"fmt"
)

type user struct {
	id   int
	name string
}

func init() {
	Register("wx", func() Pay {
		return new(pay.Wx)
	})

	Register("al", func() Pay {
		return new(pay.Ali)
	})

}

func main() {
	//fmt.Println(NewUser(1,"cdhihi"))

	payinfo, err := Create("wx")

	b := payinfo.Pay("发送")

	fmt.Println(err)
	fmt.Println(b)

}

//注册user
func NewUser(id int, name string) *user {
	return &user{id, name}
}
