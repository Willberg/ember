package sp

import (
	"fmt"
)

type Node struct {
	Name        string
	Val         int
	Left, Right *Node
}

func insert(root *Node, name string, val int) *Node {
	if root == nil {
		return &Node{Name: name, Val: val}
	}
	if root.Val <= val {
		root.Left = insert(root.Left, name, val)
	} else {
		root.Right = insert(root.Right, name, val)
	}
	return root
}

func BuildTree(names []string, vals []int) (root *Node) {
	for i, name := range names {
		root = insert(root, name, vals[i])
	}
	return root
}

func PrintTree(t *Node) {
	s := ""
	output(t, "", true, &s)
	fmt.Println(s)
}

func output(t *Node, prefix string, isTail bool, s *string) {
	if t.Right != nil {
		newPrefix := prefix
		if isTail {
			newPrefix += "│   "
		} else {
			newPrefix += "    "
		}
		output(t.Right, newPrefix, false, s)
	}
	*s += prefix
	if isTail {
		*s += "└── "
	} else {
		*s += "┌── "
	}
	*s += fmt.Sprintf("%d\n", t.Val)
	if t.Left != nil {
		newPrefix := prefix
		if isTail {
			newPrefix += "    "
		} else {
			newPrefix += "│   "
		}
		output(t.Left, newPrefix, true, s)
	}
}

func PrintTree2(root *Node, t, level int) {
	if root == nil {
		return
	}
	PrintTree2(root.Right, 2, level+1)
	switch t {
	case 0:
		fmt.Printf("%2d\n", root.Val)
	case 1:
		for i := 0; i < level; i++ {
			fmt.Print("\t")
		}
		fmt.Printf("\\ %2d\n", root.Val)
	case 2:
		for i := 0; i < level; i++ {
			fmt.Print("\t")
		}
		fmt.Printf("/ %2d\n", root.Val)
	}
	PrintTree2(root.Left, 1, level+1)
}

func Output(root *Node, n int) (res []*Node) {
	var dfs func(t *Node)
	dfs = func(t *Node) {
		if t == nil {
			return
		}
		dfs(t.Left)
		if len(res) == n {
			return
		}
		res = append(res, t)
		dfs(t.Right)
	}
	dfs(root)
	return
}
