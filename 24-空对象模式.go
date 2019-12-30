package design

import (
	"fmt"
)

type IGood interface {
	fmt.Stringer
	GetName() string
	GetPrice() int
}

type Good struct {
	name string
	price int
}

func NewGood(name string,price int) *Good {
	return &Good{name,price}
}

func (g *Good) GetName() string {
	return g.name
}

func (g *Good) GetPrice() int {
	return g.price
}

func (g *Good) String() string {
	return fmt.Sprintf("name: %s, price: %d\n",g.name,g.price)
}

type nullGood struct {
}

func (n *nullGood) GetName() string {
	return ""
}

func (n *nullGood) GetPrice() int {
	return 0
}

func (n *nullGood) String() string {
	return "没有商品信息"
}

var NullGood = new(nullGood)

type FactoryGoods struct {
}

func (f *FactoryGoods) Find(name string) IGood {
	switch name {
	case "a":
		return NewGood("a",1)
	case "b":
		return NewGood("b",2)
	default:
		return NullGood
	}
}