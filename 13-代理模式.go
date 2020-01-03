package design

import "fmt"

// 静态代理
// 委托类接口，委托类，代理类
// 委托类：真正执行业务逻辑的类
// 委托类接口：代理类和委托类要提供的方法
// 代理类：利用委托类重新实现委托类接口，可以做一些增强的实现

// 使用实例
type IBuy interface {
	Login(name, pwd string)
	BuyTicket()
}

type Buy struct {
}

func (b *Buy) Login(name, pwd string) {
	fmt.Println("登陆：", name, pwd)
}

func (b *Buy) BuyTicket() {
	fmt.Println("买票")
}

type BuyProxy struct {
	*Buy
}

func (b *BuyProxy) Login(name, pwd string) {
	b.Buy.Login(name, pwd)
}

func (b *BuyProxy) BuyTicket() {
	b.Before()
	b.Buy.BuyTicket()
	b.After()
}

func (b *BuyProxy) Before() {
	fmt.Println("开始刷票")
}

func (b *BuyProxy) After() {
	fmt.Println("买票成功，短信通知客户")
}

// 动态代理
// 个人理解，使用反射动态获取委托类的类型和方法，然后利用反射调用委托类的方法实现代理模式
// 动态代理和静态代理的区别，静态代理需要在代理类中明确指明委托类；而动态代理类中只需要知道要代理什么类，
// 并知道要怎么使用委托类的方法，然后具体代理类会做什么，需要在确定了委托类之后才能确定
