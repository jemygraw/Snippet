package main

import (
	"fmt"
)

// TreeNode is a tree
type TreeNode struct {
	Value    string
	Children []TreeNode
}

func (t *TreeNode) String() string {
	if t.Children == nil {
		return t.Value
	}

	buf := t.Value
	for _, child := range t.Children {
		buf += child.String()
	}
	return buf
}

// Height returns the height of the tree node
func (t *TreeNode) Height() int {
	if t.Children == nil {
		return 1
	}

	max := 0
	for _, child := range t.Children {
		childHeight := child.Height()
		if childHeight > max {
			max = childHeight
		}
	}
	return 1 + max
}

// Find finds the tree node of an item
func (t *TreeNode) Find(val string) (target *TreeNode) {
	if t.Value == val {
		target = t
		return
	}

	for _, child := range t.Children {
		targetTemp := child.Find(val)
		if targetTemp != nil {
			target = targetTemp
			return
		}
	}
	return
}

func main() {
	root := TreeNode{}
	root.Value = "A"
	root.Children = []TreeNode{
		TreeNode{
			Value: "B",
			Children: []TreeNode{
				TreeNode{
					Value:    "D",
					Children: nil,
				},
				TreeNode{
					Value:    "E",
					Children: nil,
				},
			},
		},
		TreeNode{
			Value: "C",
			Children: []TreeNode{
				TreeNode{
					Value:    "F",
					Children: nil,
				},
				TreeNode{
					Value:    "G",
					Children: nil,
				},
				TreeNode{
					Value: "H",
					Children: []TreeNode{
						TreeNode{
							Value:    "I",
							Children: nil,
						},
					},
				},
			},
		},
	}

	fmt.Println(root.String())
	fmt.Println(root.Height())

	fmt.Println("---")
	testCases := []string{"A", "B", "C", "D", "E", "F", "G", "H", "I"}
	//testCases = []string{"B"}
	for _, val := range testCases {
		target := root.Find(val)
		fmt.Println("target--->", target)
		if target != nil {
			fmt.Println(val, ", height:", target.Height())
		} else {
			fmt.Println("not found error", val)
		}
		fmt.Println("")
	}
}
