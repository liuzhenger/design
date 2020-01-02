package design

import "sort"

// 抽象策略角色
// 定义算法或业务执行方法
type Strategy interface {
	MySort(list []int)
}

// 具体策略角色
// 算法或业务具体实现
type BigSort struct {
}

func (b *BigSort) MySort(list []int) {
	sort.Ints(list)
	for i := 0; i < len(list)/2; i++ {
		list[i], list[len(list)-i-1] = list[len(list)-i-1], list[i]
	}
}

type MinSort struct {
}

func (m *MinSort) MySort(list []int) {
	sort.Ints(list)
}

// 上下文角色
// 保存策略的具体实现，并提供调用方法
type Context struct {
	s Strategy
}

func (c *Context) Init(s Strategy) {
	c.s = s
}

func (c *Context) DoSort(list []int) {
	c.s.MySort(list)
}
