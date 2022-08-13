package main

import "fmt"

type Signed interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64
}

type Unsigned interface {
	~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64 | ~uintptr
}

type Integer interface {
	Signed | Unsigned
}

type Float interface {
	~float32 | ~float64
}

type Ordered interface {
	Integer | Float | ~string
}

type Node[T Ordered] struct {
	Val   T
	Left  *Node[T]
	Right *Node[T]
}

func Insert[T Ordered](n *Node[T], val T) *Node[T] {
	if n == nil {
		return &Node[T]{Val: val, Left: nil, Right: nil}
	}
	if val < n.Val {
		n.Left = Insert(n.Left, val)
	} else if val > n.Val {
		n.Right = Insert(n.Right, val)
	}
	return n
}

func Search[T Ordered](n *Node[T], val T) bool {
	if n == nil {
		return false
	}
	if n.Val < val {
		return Search(n.Right, val)
	} else if n.Val > val {
		return Search(n.Left, val)
	} else {
		return true
	}
}

func main() {
	vals := []int{3, 1, 4, 3, 1, 5}
	root := Node[int]{Val: vals[0], Left: nil, Right: nil}
	for i := 1; i < len(vals); i++ {
		Insert(&root, vals[i])
	}
	fmt.Println(root.Left.Left)
}
