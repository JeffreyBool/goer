package node

import (
	"fmt"
)

type Node struct {
	Value       int
	Left, Right *Node
}

func (n Node) Print() {
	fmt.Println(n.Value)
}

//指针类型有可能是空指针,使用的时候需要进行是否为空判断.
func (n *Node) SetValue(value int) {
	if n == nil {
		fmt.Println("Setting Value to nil " +
			"node. Ignored.")
		return
	}

	n.Value = value
}

func CreateNode(value int) *Node {
	return &Node{Value: value}
}
