package main

import "fmt"

type treeNode struct {
	value int
	left  *treeNode
	right *treeNode
}

func (node *treeNode) traverse() {
	if node == nil {
		return
	}
	node.left.traverse()
	fmt.Printf("%d ", node.value)
	node.right.traverse()
}

func (node *treeNode) traverseFunction(f func(node *treeNode)) {
	if node == nil {
		return
	}
	node.left.traverseFunction(f)
	f(node)
	node.right.traverseFunction(f)
}

func (node *treeNode) traverseWithChannel() chan *treeNode {
	c := make(chan *treeNode)
	go func() {
		node.traverseFunction(func(node *treeNode) {
			c <- node
		})
		close(c)
	}()
	return c
}

func initTreeNodeData(nodes *[5]treeNode) {
	for i := range nodes {
		nodes[i] = treeNode{value: i}
	}
	nodes[0].left = &nodes[1]
	nodes[1].left = &nodes[2]
	nodes[0].right = &nodes[3]
	nodes[3].right = &nodes[4]
}

func main() {
	var nodes [5]treeNode
	initTreeNodeData(&nodes)

	//normal traverse
	fmt.Println("----Normal Traverse----")
	nodes[0].traverse()
	fmt.Println()
	fmt.Println("----Traverse Function----")
	//traverseFunction
	nodeCount := 0
	nodes[0].traverseFunction(func(node *treeNode) {
		nodeCount++
	})
	fmt.Printf("Count: %d\n", nodeCount)

	fmt.Println("----Traverse With Channel----")
	channel := nodes[0].traverseWithChannel()
	for node := range channel {
		fmt.Printf("%d ", node.value)
	}
}
