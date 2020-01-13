package design

import "fmt"

// 产品接口
type ICement interface {
	DoSomeThing()
}

type IGround interface {
	DoSomeThing()
}

type IRoof interface {
	DoSomeThing()
}

// 产品实例
type Cement struct {
}

func (c *Cement) DoSomeThing() {

}

type Ground struct {
}

func (g *Ground) DoSomeThing() {

}

type Roof struct {
}

func (r *Roof) DoSomeThing() {
	
}

// 产品
type Product struct {
	C ICement
	G IGround
	R IRoof
}

// 建造者接口
type IBuilder interface {
	BuildCement()
	BuildGround()
	BuildRoof()
	Build() *Product
}

// 建造者
type Builder struct {
	p Product
}

func (b *Builder) BuildCement() {
	fmt.Println("build cement")
	b.p.C = new(Cement)
}

func (b *Builder) BuildGround() {
	fmt.Println("build ground")
	b.p.G = new(Ground)
}

func (b *Builder) BuildRoof() {
	fmt.Println("build roof")
	b.p.R = new(Roof)
}

func (b *Builder) Build() *Product {
	return &b.p
}

// 监工，控制建造顺序
type Director struct {
	IBuilder
}

func (d *Director) BuildProduct() *Product {
	d.BuildCement()
	d.BuildGround()
	d.BuildRoof()
	return d.Build()
}
