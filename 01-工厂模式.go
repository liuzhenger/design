package design

import "fmt"

type ISend interface {
	Send()
}

type NilSend struct {
}

func (n *NilSend) Send() {
	fmt.Println("发送失败")
}

type Sms struct {
}

func (s *Sms) Send() {
	fmt.Println("发送短信")
}

type Email struct {
}

func (s *Email) Send() {
	fmt.Println("发送邮件")
}

// 普通简单工厂
type SendFactory1 struct {
}

func GetSender(name string) ISend {
	switch name {
	case "sms":
		return new(Sms)
	case "email":
		return new(Email)
	default:
		return new(NilSend)
	}
}

// 多方法简单工厂
// 多方法简单工厂和普通简单工厂其实并没有什么不同，只是为了避免使用普通简单工厂创建对象时参数传错
type SendFactory2 struct {
}

func (s *SendFactory2) SendSms() ISend {
	return new(Sms)
}

func (s *SendFactory2) SendEmail() ISend {
	return new(Email)
}

// 静态方法简单工厂
// 可以避免创建工厂类，其实工厂类全局创建一次就可以了，和静态方法简单工厂一样
// go语言没有静态方法，可以在init()方法中创建全局工厂对象的方式代替或者结合单例模式创建一个工厂对象

// 工厂方法模式
// 新增了一个工厂接口
type IProduct interface {
	Product() ISend
}

type SmsFactory struct {
}

func (s *SmsFactory) Product() ISend {
	return new(Sms)
}

type EmailFactory struct {
}

func (s *EmailFactory) Product() ISend {
	return new(Email)
}
