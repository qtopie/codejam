package main

import "fmt"

type Node struct {
	Val   int
	Left  *Node
	Right *Node
}

func preorderTraverse(p *Node) {
	if p == nil {
		return
	}

	fmt.Println(p.Val)
	preorderTraverse(p.Left)
	preorderTraverse(p.Right)
}

func inorderTraverse(p *Node) {
	if p == nil {
		return
	}

	inorderTraverse(p.Left)
	fmt.Println(p.Val)
	inorderTraverse(p.Right)
}

func postorderTraverse(p *Node) {
	if p == nil {
		return
	}

	postorderTraverse(p.Left)
	postorderTraverse(p.Right)
	fmt.Println(p.Val)
}

type GraphNode struct {
	Val      int
	Children []*GraphNode
}

func dfs(p *GrahpNode) {
	if p == nil {
		return
	}

	fmt.Println(p.Val)
	if len(p.Children) > 0 {
		for _, c := range p.Children {
			dfs(c)
		}
	}
}

func bfs(root *GraphNode) {
	if root == nil {
		return
	}

	queue := []*GraphNode{root}

	for len(queue) > 0 {
		p := queue[0]
		fmt.Println(p.Val)

		queue = queue[1:]
		// add children to queue
		for _, c := range p.Children {
			queue = append(queue, c)
		}

	}
}

func main() {
	fmt.Println("vim-go")
}
