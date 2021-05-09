package main

import "fmt"

var opts map[string]func(string1 string,string2 string) string


func array_reduce(a []string, f string) string {

	b := ""

	for _,val := range a {
		b = opts[f](b,val)
	}

	return b
}


func toString(v1 , v2 string) string {

	return v1 + "-" +v2

}


func init()  {
	opts = make(map[string]func(v1 string,v2 string) string,0)
	opts["toString"] = toString

}



func main()  {

	ar := []string{"大哥","二哥"}

	c := array_reduce(ar,"toString")

	fmt.Printf(c)
}
