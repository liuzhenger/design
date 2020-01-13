package design

import "sync"

type Object struct {
}

// 预加载
var A = Object{}

// 懒加载(并发不安全)
var a *Object

func GetOne() *Object {
	if a == nil {
		a = new(Object)
		return a
	}
	return a
}

// 懒加载(并发安全)
var once = sync.Once{}

func GetOne2() *Object {
	once.Do(func() {
		a = GetOne()
	})
	return a
}
