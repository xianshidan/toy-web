package main

import (
	"fmt"
)

func main() {

	// 因为 u 是结构体，所以方法调用的时候它数据是不会变的
	u := User{
		Name: "Tom",
		Age:  10,
	}
	u.ChangeName("Tom Changed!")
	u.ChangeAge(100)
	fmt.Printf("%v \n", u)

	// 因为 up 指针，所以内部的数据是可以被改变的
	up := &User{
		Name: "Jerry",
		Age:  12,
	}

	// 因为 ChangeName 的接收器是结构体
	// 所以 up 的数据还是不会变
	up.ChangeName("Jerry Changed!")
	up.ChangeAge(120)

	fmt.Printf("%v \n", up)
}

type User struct {
	Name string
	Age  int
}

// ChangeName 结构体接收器
// 为结构体创建方法
// 设计不可变对象
func (u User) ChangeName(newName string) {
	//http.HandlerFunc() 包的方法
	// 结构接收器的方法不会影响数据
	u.Name = newName
}

// ChangeAge 指针接收器
// 遇事不决，用指针
func (u *User) ChangeAge(newAge int) {
	// 指针接收器方法会影响数据
	u.Age = newAge
}

type HandleFunc func()

// Hello 不用指针
func (h HandleFunc) Hello() {
}
