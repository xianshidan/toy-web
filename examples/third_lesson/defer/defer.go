package main

import "fmt"

func main() {
	// 从下往上执行，不要再循环里面使用
	defer func() {
		fmt.Println("aaa")
	}()

	defer func() {
		fmt.Println("bbb")
	}()

	defer func() {
		fmt.Println("ccc")
	}()
}
