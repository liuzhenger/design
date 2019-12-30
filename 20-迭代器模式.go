package design

// 迭代器接口
type Iterator interface {
	Next() interface{}
	HasNext() bool
}

// 迭代器实例
type ConcreteIterator struct {
	a     IAggregate
	index int
}

func (c *ConcreteIterator) Init(a IAggregate) {
	c.a = a
}

func (c *ConcreteIterator) Next() interface{} {
	defer func() {
		c.index++
	}()
	if c.index < c.a.Size() {
		return c.a.Get(c.index)
	}
	return nil
}

func (c *ConcreteIterator) HasNext() bool {
	return c.index < c.a.Size()
}

// 抽象聚集接口
type IAggregate interface {
	Add(e interface{})
	CreateIterator() Iterator
	Size() int
	Get(n int) interface{}
}

// 聚集实例
type ConcreteAggregate struct {
	list []interface{}
}

func (c *ConcreteAggregate) Add(e interface{}) {
	c.list = append(c.list, e)
}

func (c *ConcreteAggregate) CreateIterator() Iterator {
	a := new(ConcreteIterator)
	a.Init(c)
	return a
}

func (c *ConcreteAggregate) Size() int {
	return len(c.list)
}

func (c *ConcreteAggregate) Get(i int) interface{} {
	if i >= 0 && i < len(c.list) {
		return c.list[i]
	}
	return nil
}
