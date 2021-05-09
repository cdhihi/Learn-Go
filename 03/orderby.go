package main

import (
	"fmt"
	"sort"
	"github.com/cdhihi/shanyuan"
)

func main()  {
	arr := map[int]string{}
	
	arr[1] = "a"
	arr[4] = "d"
	arr[5] = "e"
	arr[3] = "c"
	arr[2] = "b"


	var keys []int

	for k,_ := range arr{

		keys = append(keys, k)
	}

	//切片排序数字
	sort.Ints(keys)

	//字符串排序
	//sort.Strings(keys)


	for _,v := range keys{

		fmt.Printf("key=%d val = %s \n",v,arr[v])

	}


	//发布的三元运算符测试
	println(shanyuan.If(1>2,"true","false").(string))
	
}
