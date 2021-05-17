package main

import (
	"fmt"
	"reflect"
)

type pay struct {
}

func getTypeAndValue(i interface{}) {

	t := reflect.TypeOf(i)
	v := reflect.ValueOf(i)

	fmt.Println("输出为 ", t, " 类型", "值为 ", v)

}

func reflectType(i interface{}) {

	t := reflect.TypeOf(i)
	v := reflect.ValueOf(i)

	fmt.Println("输出为 ", t, " 类型", "值为 ", v.Elem())

}

func main() {
	a := 1
	// //
	b := pay{}
	//b := 1
	getTypeAndValue(a)

	reflectType(&b)

}
