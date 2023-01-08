package main

import "net/http"

// Server1 可以放到其它包里面，这里不需要引用包，也可以实现接口
type Server1 interface {
	Router(pattern string, handlerFunc http.HandlerFunc)
	Start(address string) error
}

// sdkHttpServer1 基于http实现
type sdkHttpServer1 struct {
	Name string
}

// Router  注册路由
func (s *sdkHttpServer1) Router(pattern string, handlerFunc http.HandlerFunc) {
	http.HandleFunc(pattern, handlerFunc)
}

func (s *sdkHttpServer1) Start(address string) error {
	return http.ListenAndServe(address, nil)
}

func NewServer(name string) Server1 {
	return &sdkHttpServer1{
		Name: name,
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
