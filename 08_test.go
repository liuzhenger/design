package design

import (
	"fmt"
	"testing"
)

func Test_08_1(t *testing.T) {
	filter := func(r Request) bool {
		if r.Length > 5 {
			fmt.Println("data too lang")
			return false
		}
		return true
	}

	s := new(Server)
	s.Init()
	s.AddFilter(filter)
	s.Handler("login", func(r Request) {
		fmt.Println("登陆中：", r.Name)
	})
	r := Request{
		Url:    "login",
		Name:   "小明",
		Length: 10,
	}
	s.Do(r)
	r.Length = 1
	s.Do(r)
}
