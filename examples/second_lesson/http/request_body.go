package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hi, this  is home page")
}

func readBodyOnce(w http.ResponseWriter, r *http.Request) {
	body, err := io.ReadAll(r.Body)
	if err != nil {
		fmt.Fprintf(w, "read body failed: %v", err)
		// 记住要返回，不然就还会执行后面的代码
		return
	}
	// 类型转换，将 []byte 转换为 string
	fmt.Fprintf(w, "read the data: %s \n", string(body))

	// 尝试再次读取，啥也读不到，但是也不会报错
	body, err = io.ReadAll(r.Body)
	if err != nil {
		// 不会进来这里
		fmt.Fprintf(w, "read the data one more time got error: %v", err)
		return
	}
	fmt.Fprintf(w, "read the data one more time: [%s] and read data length %d \n", string(body), len(body))
}

func getBodyIsNil(w http.ResponseWriter, r *http.Request) {

	//GetBody：原则上是可以多次读取，但是在原生的http。Request！里面，这个是nil
	// 中间件处理后，可以这样读取两次
	//body, _ := r.GetBody()
	//io.ReadAll(body)

	//body, _ = r.GetBody()
	//io.ReadAll(body)

	if r.GetBody == nil {
		fmt.Fprint(w, "GetBody is nil \n")
	} else {
		fmt.Fprintf(w, "GetBody not nil \n")
	}
}

func queryParams(w http.ResponseWriter, r *http.Request) {
	values := r.URL.Query()
	//fmt.Printf("values type %T\n", values)
	name := values["name"][2]
	fmt.Println(name)
	fmt.Fprintf(w, "query is %v\n", values)
}

func wholeUrl(w http.ResponseWriter, r *http.Request) {
	data, _ := json.Marshal(r.URL)
	fmt.Fprintf(w, string(data))
}

// Go会自动将header名字转为标准名字一其实就是大小change
func header(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("type: %T\n", r.Header["Wpw-Sid-Abc"])
	fmt.Println("aaar", r.Header["Wpw-Sid-Abc"])
	fmt.Fprintf(w, "header is %v\n", r.Header)
}

func form(w http.ResponseWriter, r *http.Request) {
	// 要先调用ParseForm
	// 建议加上Content--Type:application/x-www-form-urlencoded

	fmt.Fprintf(w, "before parse form %v\n", r.Form)
	err := r.ParseForm()
	if err != nil {
		fmt.Fprintf(w, "parse form error %v\n", r.Form)
	}
	fmt.Fprintf(w, "before parse form %v\n", r.Form)
}

func main() {
	http.HandleFunc("/", home)
	http.HandleFunc("/body/once", readBodyOnce)
	http.HandleFunc("/body/multi", getBodyIsNil)
	http.HandleFunc("/url/query", queryParams)
	http.HandleFunc("/header", header)
	http.HandleFunc("/wholeUrl", wholeUrl)
	http.HandleFunc("/form", form)
	if err := http.ListenAndServe(":8080", nil); err != nil {
		panic(err)
	}
}
