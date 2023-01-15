package main

import (
	"encoding/json"
	"errors"
	"fmt"
	webv2 "geektime/toy-web/pkg/v2"
	"io"
	"net/http"
)

func home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "这是主页")
}

func user(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "这是用户")
}

func createUser(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "这是创建用户")
}

func order(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "这是订单")
}

type Server interface {
	Route(pattern string, handlerFunc http.HandlerFunc)
	Start(address string) error
}

type sdkHttpServer struct {
	Name string
}

func SignUp(ctx *Context) {
	req := &signUpReq{}

	err := ctx.ReadJson(req)
	if err != nil {
		ctx.BadRequestJson(err)
		return
	}

	resp := &commonResponse{
		Data: 123,
	}
	err = ctx.WriteJson(http.StatusOK, resp)
	if err != nil {
		fmt.Printf("WriteJson fail: %v ", err)
	}

	respJson, err := json.Marshal(resp)
	if err != nil {

	}
	// 返回一个虚拟的 user id 表示注册成功了
	fmt.Fprintf(ctx.W, "%s", respJson)
}

// SignUpWithoutContext 在没有 context 抽象的情况下，是长这样的
func SignUpWithoutContext(w http.ResponseWriter, r *http.Request) {
	req := &signUpReq{}
	body, err := io.ReadAll(r.Body)
	if err != nil {
		fmt.Fprintf(w, "read body failed: %v", err)
		// 要返回掉，不然就会继续执行后面的代码
		return
	}
	err = json.Unmarshal(body, req)
	if err != nil {
		fmt.Fprintf(w, "deserialized failed: %v", err)
		// 要返回掉，不然就会继续执行后面的代码
		return
	}

	resp := &commonResponse{
		Data: 123,
	}
	respJson, err := json.Marshal(resp)
	if err != nil {

	}

	// 返回一个虚拟的 user id 表示注册成功了
	fmt.Fprintf(w, "%s", respJson)
}

func SignUpWithoutWrite(w http.ResponseWriter, r *http.Request) {
	c := webv2.NewContext(w, r)
	req := &signUpReq{}
	err := c.ReadJson(req)
	if err != nil {
		resp := &commonResponse{
			BizCode: 4, // 假如说我们这个代表输入参数错误
			Msg:     fmt.Sprintf("invalid request: %v", err),
		}
		respBytes, _ := json.Marshal(resp)
		fmt.Fprint(w, string(respBytes))
		return
	}
	// 这里又得来一遍 resp 转json
	fmt.Fprintf(w, "invalid request: %v", err)
}

type signUpReq struct {
	Email             string `json:"email"`
	Password          string `json:"password"`
	ConfirmedPassword string `json:"confirmed_password"`
}

type commonResponse struct {
	BizCode int         `json:"biz_code"`
	Msg     string      `json:"msg"`
	Data    interface{} `json:"data"`
}

func main() {
	server := NewServer("test-server")
	//server.Router("/", home)
	//server.Router("/user", user)
	//server.Router("/user/create", createUser)
	//server.Router("/order", order)
	server.Router(http.MethodGet, "/user/signup", SignUp)
	err := server.Start(":8080")
	er := errors.New("error")
	fmt.Println(er)
	if err != nil {
		panic(err)
	}

	//http.HandleFunc("/", home)
	//http.HandleFunc("/user", user)
	//http.HandleFunc("/user/create", createUser)
	//http.HandleFunc("/order", order)
	//http.ListenAndServe(":8080", nil)
}
