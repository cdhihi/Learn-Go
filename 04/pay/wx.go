package pay

type Wx struct {
}

func (payinfo *Wx) Pay(data string) string {

	return "alipay => data" + data
}
