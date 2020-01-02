package design

import (
	"fmt"
	"github.com/influxdata/influxdb/pkg/slices"
)

// 元素允许访问的接口
// 实现这个接口就代表元素可以被访问
type IElement interface {
	Accept(visitor IVisitor)
}

// 实例元素A
type ElementA struct {
	Name      string
	LookPerms []string
}

func (e *ElementA) Accept(v IVisitor) {
	v.VisitA(e)
}

// 实例元素B
type ElementB struct {
	Name      string
	LookPerms []string
}

func (e *ElementB) Accept(v IVisitor) {
	v.VisitB(e)
}

// 访问者接口
// 定义所有元素的访问方法
type IVisitor interface {
	VisitA(e *ElementA)
	VisitB(e *ElementB)
}

// 访问者实例
type VisitorRoot struct {
}

func (v *VisitorRoot) VisitA(e *ElementA) {
	fmt.Println("管理员访问：", e.Name)
}

func (v *VisitorRoot) VisitB(e *ElementB) {
	fmt.Println("管理员访问：", e.Name)
}

type Visitor struct {
	Perm string
}

func (v *Visitor) VisitA(e *ElementA) {
	if len(e.LookPerms) == 0 || slices.Exists(e.LookPerms, v.Perm) {
		fmt.Println("普通人员访问：", e.Name)
	} else {
		fmt.Println("普通人员访问：", "权限不够，无法访问")
	}
}

func (v *Visitor) VisitB(e *ElementB) {
	if len(e.LookPerms) == 0 || slices.Exists(e.LookPerms, v.Perm) {
		fmt.Println("普通人员访问：", e.Name)
	} else {
		fmt.Println("普通人员访问：", "权限不够，无法访问")
	}
}

// 元素容器
type Data struct {
	l []IElement
}

func (d *Data) Add(e IElement) {
	if e == nil {
		return
	}
	d.l = append(d.l, e)
}

func (d *Data) Show(visitor IVisitor) {
	for _, v := range d.l {
		v.Accept(visitor)
	}
}
