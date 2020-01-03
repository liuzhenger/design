package design

import "testing"

func Test_13_1(t *testing.T) {
	b := new(Buy)
	p := new(BuyProxy)
	p.Buy = b
	p.Login("Tom", "123456")
	p.BuyTicket()
}
