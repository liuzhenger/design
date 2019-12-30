package design

import "testing"

func Test_22_1(t *testing.T) {
	p := &Platform{}
	r1 := &Reader{"R1"}
	r2 := &Reader{"R2"}
	p.Attach(r1)
	p.Attach(r2)
	p.Notify(5)
	p.Detach(r1)
	p.Notify(6)
}
