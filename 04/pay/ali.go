package pay

type Ali struct {
}

func (payinfo *Ali) Pay(data string) string {

	return "alipay => data" + data
}
