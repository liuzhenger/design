package design

// 请求消息
type Request struct {
	Url    string
	Name   string
	Length int
}

// 过滤器
type Filter func(request Request) bool

// 服务器
type Server struct {
	fs      []Filter
	handler map[string]func(request Request)
}

func (s *Server) Init() {
	s.handler = make(map[string]func(request Request))
}

func (s *Server) AddFilter(f Filter) {
	s.fs = append(s.fs, f)
}

func (s *Server) Handler(name string, f func(request Request)) {
	s.handler[name] = f
}

func (s *Server) Do(r Request) {
	for _, v := range s.fs {
		if !v(r) {
			return
		}
	}
	// 执行业务逻辑
	f := s.handler[r.Url]
	f(r)
}
