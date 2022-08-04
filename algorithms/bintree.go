package main

import "fmt"

type node struct {
	Val   int
	Left  *node
	Right *node
}

func insert(n *node, val int) *node {
	if n == nil {
		n = &node{Val: val, Left: nil, Right: nil}
	} else if val < n.Val {
		n.Left = insert(n.Left, val)
	} else if val > n.Val {
		n.Right = insert(n.Right, val)
	}
	return n
}

func search(n *node, val int) bool {
	if n == nil {
		return false
	}
	if n.Val < val {
		return search(n.Right, val)
	} else if n.Val > val {
		return search(n.Left, val)
	} else {
		return true
	}
}

func main() {
	vals := []int{3, 1, 4, 3, 1, 5}
	root := node{Val: vals[0], Left: nil, Right: nil}
	for i := 1; i < len(vals); i++ {
		insert(&root, vals[i])
	}
	fmt.Println(root.Left.Left)
}
