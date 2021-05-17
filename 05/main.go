package main

import (
	"fmt"
	"reflect"
)

type user struct {
	id   int
	name string
	info *userInfo
}

type userInfo struct {
	age  int
	ctiy string
}

func (u *user) GetName() string {

	return u.name
}

func (u *user) GetId(id int) (int, string) {

	return id, "www.baidu.com"
}

func main() {

	//反射

	//a := "baidu.com"
	//
	//v := reflect.TypeOf(a)
	//
	//t := reflect.ValueOf(a)
	//
	//
	//fmt.Println(v)
	//fmt.Println(t)
	//fmt.Println(v.Kind().String())
	//
	//
	//
	//fmt.Printf("v = %T , t= %T \n",v,t)
	//
	//
	//fmt.Println(v.Kind().String() == "string")

	u := user{
		10,
		"www.baidu.com",
		&userInfo{
			18,
			"四川",
		},
	}

	fmt.Println(u)

	t := reflect.TypeOf(&u)

	v := reflect.ValueOf(&u)

	fmt.Println(v, t)

	ufile, _ := t.Elem().FieldByName("id")

	fmt.Println(v.Elem().FieldByName("id"))
	fmt.Println(ufile.Name)

	fmt.Println(v.Elem().Field(1))
	fmt.Println(t.Elem().Field(1).Name)

	//反射方法
	a := v.MethodByName("GetId").Call([]reflect.Value{
		reflect.ValueOf(100),
	})

	//fmt.Println(v.MethodByName("GetId").Call([]reflect.Value{
	//	reflect.ValueOf(100),
	//}))

	fmt.Println(a[2])

}
