package design

import "testing"

func Test_18_1(t *testing.T) {
	c1 := new(Colleague)
	c1.Name = "C1"
	c2 := new(Colleague)
	c2.Name = "C2"
	m := new(Mediator)
	m.Name = "M"
	c1.M = m
	c2.M = m
	m.C1 = c1
	m.C2 = c2
	c1.DoSelfMethod()
	c2.DoSelfMethod()
	c1.DepMethod()
	m.DoSelfMethod()
	m.DoSomeThing()
}

func Test_18_2(t *testing.T) {
	p1 := &Player{Name: "P1"}
	p2 := &Player{Name: "P2"}
	hm := new(HouseMediator)
	hm.AddPlayer(p1)
	hm.AddPlayer(p2)
	p1.HM = hm
	p2.HM = hm
	p1.Publish("1号房")
	p2.Publish("2号房")
	p1.Find("P1")
	p1.Find("P2")
	hm.RemovePlayer(p1)
	p2.Find("P1")
}
