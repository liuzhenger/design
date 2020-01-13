package design

import "testing"

func Test_02_01(t *testing.T) {
	factory := new(MeiDiFactory)
	tv := factory.NewTv().Product()
	television := factory.NewTelevision().Product()

	tv.DoSomeThing()
	television.DoSomeThing()
}
