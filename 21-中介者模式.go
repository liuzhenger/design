package design

import "fmt"

// 迪米特法则的典型应用

// 同事角色
// 一般定义的方法是给中介者使用
type IColleague interface {
	DoSelfMethod()
}

type Colleague struct {
	Name string
	M IMediator // 中介者
}

// 不依赖中介者的方法
func (c *Colleague) DoSelfMethod() {
	fmt.Println(c.Name," do self method")
}

// 依赖中介者的方法
func (c *Colleague) DepMethod() {
	fmt.Println(c.Name," DepMethod...")
	c.M.DoSomeThing()
}

// 中介者角色
// 包含有依赖关系的各种对象
// 一般定义的方法是给其它对象使用，体现对象之间的依赖关系
type IMediator interface {
	DoSomeThing()
}

type Mediator struct {
	Name string
	C1,C2 IColleague
}

func (b *Mediator) DoSomeThing() {
	fmt.Println(b.Name," DoSomeThing...")
	b.C1.DoSelfMethod()
	b.C2.DoSelfMethod()
}

func (b *Mediator) DoSelfMethod() {
	fmt.Println(b.Name," do self method")
}

// 实例
// 客户和房产中介
type IPlayer interface {
	GetList() []string
	Receiver(name,str string)
}

type Player struct {
	HM IHouseMediator
	Name string
	list []string
}

func (p *Player) Publish(str string) {
	fmt.Println("发布：",p.Name," ",str)
	p.list = append(p.list,str)
	p.HM.Publish(p.Name,str)
}

func (p *Player) Find(name string) {
	fmt.Println("查询：",name)
	list := p.HM.GetList(name)
	for _,v := range list {
		fmt.Println(name," ",v)
	}
}

func (p *Player) GetList() []string {
	return p.list
}

func (p *Player) Receiver(name,str string)  {
	fmt.Println("收到发布消息：",name," ",str)
}

type IHouseMediator interface {
	Publish(name ,str string)
	GetList(name string) []string
}

type HouseMediator struct {
	players []*Player
}

func (h *HouseMediator) AddPlayer(p *Player) {
	if p == nil {
		return
	}
	for _,v := range h.players {
		if v == p {
			return
		}
	}
	h.players = append(h.players,p)
}

func (h *HouseMediator) RemovePlayer(p *Player) {
	if p == nil {
		return
	}
	for k,v := range h.players {
		if v == p {
			h.players = append(h.players[:k],h.players[k+1:]...)
			return
		}
	}
}

func (h *HouseMediator) Publish(name, str string) {
	for _,v := range h.players {
		if v.Name != name {
			v.Receiver(name,str)
		}
	}
}

func (h *HouseMediator) GetList(name string) []string {
	for _,v := range h.players {
		if v.Name == name {
			return v.GetList()
		}
	}
	return []string{}
}
