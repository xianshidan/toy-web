package main

import "fmt"

func main() {
	son := Son{
		Parent{},
	}

	// java: i am son
	// golang: i am parent
	//son.SayHello()
	// 显示调用
	son.Parent.SayHello()
}

type Parent struct {
}

func (p Parent) SayHello() {
	fmt.Println("I am " + p.Name())
}

func (p Parent) Name() string {
	return "Parent"
}

type Son struct {
	Parent
}

// Name 定义了自己的 Name() 方法
func (s Son) Name() string {
	return "Son"
}
