package design

import "fmt"

// 发布订阅模式

// 订阅者接口
type Observer interface {
	Update(n int)
}

// 发布者接口
type ISubject interface {
	Attach(o Observer)
	Detach(o Observer)
	Notify(n int)
}

// 发布者实例
type Platform struct {
	list []Observer
}

func (p *Platform) Attach(o Observer) {
	for _,v := range p.list {
		if v == o {
			return
		}
	}
	p.list = append(p.list,o)
}

func (p *Platform) Detach(o Observer) {
	for k,v := range p.list {
		if v == o {
			p.list = append(p.list[:k],p.list[k+1:]...)
			return
		}
	}
}

func (p *Platform) Notify(n int) {
	for _,v := range p.list {
		v.Update(n)
	}
}

// 订阅者实例
type Reader struct {
	Name string
}

func (r *Reader) Update(n int) {
	fmt.Println(r.Name," 收到通知 ",n)
}



