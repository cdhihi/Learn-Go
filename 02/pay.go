package main

import "fmt"

const (
	wx = iota
	al = iota
)

//阿里结构体
type Alipay struct {
	id int
	name  string
	amount float64
}

//微信结构体
type WeChatPay struct {
	id int
	name string
	amount float64
}

type payinfo interface {
	pay(float64)
	info()
}


type getinfo interface {
	tostring()
}


func (p *Alipay) pay(amount float64)  {

	fmt.Printf("%s:支付金额：%f元成功",p.name,amount)
}

func (p *Alipay) info()  {
	fmt.Println(p)
}

func (p *WeChatPay) info()  {
	fmt.Println(p)
}

func (p *WeChatPay) pay(amount float64)  {

	fmt.Printf("%s:支付金额:%f元成功",p.name,amount)
}

func (ali *Alipay) tostring()  {

	fmt.Println(ali)
}

func (wx *WeChatPay) tostring()  {

	fmt.Println(wx)
}

func order(p payinfo,amount float64)  {
	p.pay(amount)
	p.info()
}


//选择支付
func selPay(s string) interface{} {

	switch s {
	case "wx":
		//return createPay(wx)
		return createPay(wx)
	case "al":
		return createPay(al)

	}
	return nil

}

//创建支付结构体
func createPay(i int) interface{} {

	switch i {
	case 0:
		return &WeChatPay{
			name: "微信",
		}
	case 1:
		return &Alipay{
			name: "支付宝",
		}

	}
	

	return nil

}

func Try(f func(), han func(interface{}))  {
	defer func() {

		if e := recover();e != nil {
			han(e)
		}
	}()
	f()
}

func main(){



	//amount := 10.00
	//
	//order(&wx,amount)



	Try(func() {
		a := selPay("wx")
		if a == nil {

			panic("创建支付结构体异常")

		}else {
			b := a.(payinfo)
			b.pay(10.33)
			fmt.Println(b)
			fmt.Printf("type:%T",b)
		}

	}, func(e interface{}) {

		fmt.Println(e.(string))
	})







}
