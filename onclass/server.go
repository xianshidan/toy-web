package main

import "net/http"

// Server1 可以放到其它包里面，这里不需要引用包，也可以实现接口
type Server1 interface {
	Start(address string) error
	Routable
}

// sdkHttpServer1 基于http实现
type sdkHttpServer1 struct {
	Name    string
	handler Handler
}

// Router  注册路由
func (s *sdkHttpServer1) Router(method string, pattern string, handlerFunc func(ctx *Context)) {

	// 方法一：
	//http.HandleFunc(pattern, func(writer http.ResponseWriter, request *http.Request) {
	//	// 比较恶心
	//	//if request.Method != method {
	//	//	writer.Write([]byte("error"))
	//	//}
	//	ctx := NewContext(writer, request)
	//	handlerFunc(ctx)
	//})

	// 方法二:
	s.Router(method, pattern, handlerFunc)
}

func (s *sdkHttpServer1) Start(address string) error {
	http.Handle("/", s.handler)
	return http.ListenAndServe(address, nil)
}

func NewServer(name string) Server1 {
	return &sdkHttpServer1{
		Name:    name,
		handler: NewHandleBaseOnMap(),
	}
}

//type Factroy func() Server1
//
//var factory Factroy
//
//func RegisterFactory(f Factroy) {
//	factory = f
//}
//
//func NewServer1() Server1 {
//	return factory()
//}
