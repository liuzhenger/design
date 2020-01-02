package design

import (
	"fmt"
	"testing"
)

func Test_21_1(t *testing.T) {
	f := new(FactoryGoods)
	a := f.Find("a")
	b := f.Find("b")
	c := f.Find("c")
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
}
