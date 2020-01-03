package design

import "fmt"

// 建立对象池，实现对象重用
// 重点是享元角色中的内部状态和外部状态的划分
// 内部状态和外部状态的变化不会相互影响
// *内部状态相同的对象其实是同一个对象，这一点要注意

// 使用场景，系统中存在大量相似对象，需要缓冲池的场景

// 角色结构
// 抽象享元角色
// 具体享元角色
// 享元工厂

type IShape interface {
	Draw()
}

type Circle struct {
	Color string
	X     int
	Y     int
}

func (c *Circle) Draw() {
	fmt.Printf("Circle Color:%s X:%d Y:%d\n", c.Color, c.X, c.Y)
}

type CircleFactory struct {
	cs map[string]IShape
}

func (c *CircleFactory) Init() {
	c.cs = make(map[string]IShape)
}

func (c *CircleFactory) Get(color string) IShape {
	res, ok := c.cs[color]
	if !ok {
		fmt.Println("创建对象")
		c.cs[color] = &Circle{Color: color}
		return c.cs[color]
	}
	return res
}
