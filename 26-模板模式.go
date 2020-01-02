package design

import "fmt"

// 模板接口
// 包含模板方法，自定义方法，公共方法等
type ITemplate interface {
	TemplateMethod()
	DoSomething()
	PublicMethod()
}

// 具体模板接口
type ATemplate interface {
	DoSelfMethod()
}

// 模板实例
type Template struct {
	Temp ATemplate
}

func (t *Template) TemplateMethod() {
	fmt.Println("模板方法：")
	t.PublicMethod()
	t.Temp.DoSelfMethod()
}

func (t *Template) DoSomething() {
	fmt.Println("普通方法")
}

func (t *Template) PublicMethod() {
	fmt.Println("公共方法")
}

// 具体模板实例
type One struct {
}

func (o *One) DoSelfMethod() {
	fmt.Println("具体模板方法")
}
