package design

import "testing"

func Test_04_01(t *testing.T) {
	d := new(Director)
	d.IBuilder = new(Builder)
	d.BuildProduct()
}
