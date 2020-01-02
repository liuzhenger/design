package design

import "testing"

func Test_14_1(t *testing.T) {
	q1 := new(ARequest)
	q1.Init(1, "问题1")
	q2 := new(BRequest)
	q2.Init(2, "问题2")
	q3 := new(CRequest)
	q3.Init(3, "问题3")

	h1 := new(AHandler)
	h1.Level = 1
	h2 := new(BHandler)
	h2.Level = 2

	m := new(ManagerHandler)
	m.Add(h1)
	m.Add(h2)
	m.HandlerMessage(q2)
	m.HandlerMessage(q1)
	m.HandlerMessage(q3)
}
