package main

import "net/http"

type Routable interface {
	Router(method string, pattern string, handlerFunc func(ctx *Context))
}

type Handler interface {
	http.Handler
	Routable
}

type HandleBaseOnMap struct {
	// key 应该是method+url
	handlers map[string]func(ctx *Context)
}

func (h *HandleBaseOnMap) Router(method string, pattern string, handlerFunc func(ctx *Context)) {
	key := h.keys(method, pattern)
	h.handlers[key] = handlerFunc
}

func (h *HandleBaseOnMap) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	key := h.keys(request.Method, request.URL.Path)
	if handlers, ok := h.handlers[key]; ok {
		handlers(NewContext(writer, request))
	} else {
		writer.WriteHeader(http.StatusNotFound)
		writer.Write([]byte("NOT FOUND"))
	}
}

func (h *HandleBaseOnMap) keys(method string, pattern string) string {
	return method + "#" + pattern
}

// 确保HandleBaseOnMap一定实现了接口
var _ Handler = &HandleBaseOnMap{}

// NewHandleBaseOnMap 类似构造函数，返回接口？
func NewHandleBaseOnMap() Handler {
	return &HandleBaseOnMap{
		handlers: make(map[string]func(ctx *Context), 4),
	}
}
