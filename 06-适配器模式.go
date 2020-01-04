package design

import "fmt"

// 结构
// target 需要的接口，目标需求
// adapted 现有对象，并不符合target的需求
// adapter 适配器，用adapted实现target需求

type ITarget interface {
	Method1()
	Method2()
}

type Adapted struct {
}

func (a *Adapted) MethodA() {

}

func (a *Adapted) MethodB() {

}

// 面向对象语言中的类适配器模式
// 就是适配器类继承adapted类
// go中不支持继承，有时可以用组合实现
type Adapter struct {
	Adapted
}

func (a *Adapter) Method1() {
	fmt.Print("do something 1")
	a.MethodA()
}

func (a *Adapter) Method2() {
	fmt.Print("do something 2")
	a.MethodB()
}

// 面向对象中的对象适配器模式
// 就是适配器对象中包含adapted的引用
type Adapter1 struct {
	*Adapted
}

func (a *Adapter1) Method1() {
	fmt.Print("do something 1")
	a.MethodA()
}

func (a *Adapter1) Method2() {
	fmt.Print("do something 2")
	a.MethodB()
}
