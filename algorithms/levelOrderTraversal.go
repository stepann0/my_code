package main

import (
	"fmt"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func (t *TreeNode) String() string {
	if t == nil {
		return "_"
	}
	// return fmt.Sprintf("{%d L%s R%s}", t.Val, t.Left, t.Right)
	return fmt.Sprintf("%d", t.Val)
}

type Queue[T any] struct {
	slice []T
}

func (q *Queue[T]) push(val T) {
	q.slice = append(q.slice, val)
}

func (q *Queue[T]) pop() T {
	node := q.slice[0]
	q.slice = q.slice[1:]
	return node
}

func (q *Queue[T]) len() int {
	return len(q.slice)
}

func levelOrder(root *TreeNode) [][]int {
	all_levels := [][]int{}
	queue := Queue[*TreeNode]{}
	queue.push(root)

	for queue.len() > 0 {
		level_nodes := []int{}
		qLen := queue.len()
		for i := 0; i < qLen; i++ {
			node := queue.pop()
			if node != nil {
				level_nodes = append(level_nodes, node.Val)
				queue.push(node.Left)
				queue.push(node.Right)
			}
		}

		if len(level_nodes) > 0 {
			all_levels = append(all_levels, level_nodes)
		}
	}
	return all_levels
}

func main() {
	seven, four := TreeNode{7, nil, nil}, TreeNode{4, nil, nil}
	five := TreeNode{5, &seven, nil}
	three := TreeNode{3, nil, &four}
	two := TreeNode{2, nil, &five}
	one := TreeNode{1, &two, &three}
	fmt.Println(levelOrder(&one))
}
