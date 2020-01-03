package design

import "testing"

func Test_12_1(t *testing.T) {
	f := new(CircleFactory)
	f.Init()
	c := f.Get("red").(*Circle)
	c.X = 1
	c.Y = 1
	c.Draw()
	c1 := f.Get("red").(*Circle)
	c1.X = 2
	c1.Y = 2
	c1.Draw()
	c2 := f.Get("blue").(*Circle)
	c2.X = 3
	c2.Y = 3
	c2.Draw()
}
