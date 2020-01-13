package design

import "fmt"

// 抽象工厂模式
// 就是提供一组创建工厂类的接口

// 产品接口
type ITv interface {
	DoSomeThing()
}

type ITelevision interface {
	DoSomeThing()
}

// 产品实例
type MeiDiTv struct {
}

func (m *MeiDiTv) DoSomeThing() {
	fmt.Println("美的电视")
}

type MeiDiTelevision struct {
}

func (m *MeiDiTelevision) DoSomeThing() {
	fmt.Println("美的冰箱")
}

// 产品工厂接口
type TvFactory interface {
	Product() ITv
}

type TelevisionFactory interface {
	Product() ITelevision
}

// 产品工厂实例
type MeiDiTvFactory struct {
}

func (m *MeiDiTvFactory) Product() ITv {
	return new(MeiDiTv)
}

type MeiDiTelevisionFactory struct {
}

func (m *MeiDiTelevisionFactory) Product() ITelevision {
	return new(MeiDiTelevision)
}

// 抽象工厂接口，包含多个创建工厂对象的方法
type Factory interface {
	NewTv() TvFactory
	NewTelevision() TelevisionFactory
}

// 抽象工厂实例
type MeiDiFactory struct {
}

func (m *MeiDiFactory) NewTv() TvFactory {
	return new(MeiDiTvFactory)
}

func (m *MeiDiFactory) NewTelevision() TelevisionFactory {
	return new(MeiDiTelevisionFactory)
}
