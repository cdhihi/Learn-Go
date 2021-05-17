package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func main() {
	filePath := "C:\\Users\\Admin\\go\\src\\Learn-Go\\06\\log\\20210517.log"

	//写入操作
	//file,_ := os.OpenFile(filePath,os.O_WRONLY | os.O_CREATE | os.O_APPEND,0666)
	//
	//file.WriteString("=======>file := 20210516.log \n")
	//
	//
	//file.WriteString("=======>结束  \n")

	//读操作

	//file,_ := os.Open(filePath)
	//
	//b , _ :=ioutil.ReadAll(file)

	//fmt.Println(string(b))

	//分段读取
	//buf := make([]byte,128)
	//
	//
	//var content []byte
	//
	//for  {
	//	n,err := file.Read(buf)
	//
	//	if err != nil {
	//		fmt.Println("error ",err)
	//
	//		if err == io.EOF {
	//			fmt.Println("文件读取完成")
	//			break
	//		}
	//		break
	//	}
	//
	//	fmt.Println(1111)
	//	fmt.Println(n)
	//	fmt.Println(buf[:n])
	//	fmt.Println(buf)
	//	content = append(content,buf[:n]...)
	//}
	//
	//fmt.Println(string(content))

	//使用buffer

	//file,_ := os.OpenFile(filePath,os.O_CREATE|os.O_WRONLY,0666)

	//bw := bufio.NewWriter(file)
	//
	//bw.WriteString("=======>file := 202105167.log bufio \n")
	//
	//
	//bw.WriteString("=======>结束  \n")
	//
	//bw.Flush()
	file, _ := os.Open(filePath)
	//buf := make([]byte,128) //设置每次读取多少

	//var content []byte
	var content string

	br := bufio.NewReader(file)
	for true {

		//n,err := br.Read(buf)

		data, err := br.ReadString('\n')

		if err != nil {
			fmt.Println("error ", err)

			if err == io.EOF {
				fmt.Println("读取结束 ", err)
			}
			break
		}

		//content = append(content,buf[:n]...)

		content += data

	}

	fmt.Println(content)

}
