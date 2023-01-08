package main

func main() {

}

type Node struct {
	// 自引用只能使用指针，一般指针大小是固定的
	// 一般定义先计算大小，递归定义，会无穷大小
	//left Node
	//right Node

	left  *Node
	right *Node

	// 这个也会报错
	// nn NodeNode
}

type NodeNode struct {
	node Node
}
