package design

import "testing"

func Test_26_1(t *testing.T) {
	temp := new(Template)
	temp.Temp = new(One)
	temp.PublicMethod()
	temp.DoSomething()
	temp.TemplateMethod()
}
