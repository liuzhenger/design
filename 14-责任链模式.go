package design

import "fmt"

// 请求对象接口
type IRequest interface {
	GetLevel() int
	GetRequest() string
}

// 请求实例1
type ARequest struct {
	level   int
	request string
}

func (a *ARequest) Init(level int, request string) {
	a.level = level
	a.request = request
}

func (a *ARequest) GetLevel() int {
	return a.level
}

func (a *ARequest) GetRequest() string {
	return a.request
}

// 请求实例2
type BRequest struct {
	level   int
	request string
}

func (b *BRequest) Init(level int, request string) {
	b.level = level
	b.request = request
}

func (b *BRequest) GetLevel() int {
	return b.level
}

func (b *BRequest) GetRequest() string {
	return b.request
}

// 请求实例3
type CRequest struct {
	level   int
	request string
}

func (c *CRequest) Init(level int, request string) {
	c.level = level
	c.request = request
}

func (c *CRequest) GetLevel() int {
	return c.level
}

func (c *CRequest) GetRequest() string {
	return c.request
}

// 处理请求接口
type Handler interface {
	GetLevel() int
	Response(r IRequest)
}

// 处理请求的实例1
type AHandler struct {
	Level int
}

func (a *AHandler) GetLevel() int {
	return a.Level
}

func (a *AHandler) Response(r IRequest) {
	fmt.Println("处理请求：", r.GetLevel(), r.GetRequest())
}

// 处理请求的实例2
type BHandler struct {
	Level int
}

func (b *BHandler) GetLevel() int {
	return b.Level
}

func (b *BHandler) Response(r IRequest) {
	fmt.Println("处理请求：", r.GetLevel(), r.GetRequest())
}

// 处理器的管理器
// 维护处理器顺序，新增和移除处理器等等
type ManagerHandler struct {
	h []Handler
}

func (m *ManagerHandler) Add(h Handler) {
	if h == nil {
		return
	}
	m.h = append(m.h, h)
}

// 接收请求，寻找处理器并处理
func (m *ManagerHandler) HandlerMessage(r IRequest) {
	for _, v := range m.h {
		if r.GetLevel() <= v.GetLevel() {
			v.Response(r)
			return
		}
	}
	fmt.Println("not found handler")
}
