package main

import (
	"fmt"
	"iter"
)

type Node struct {
	Data     string
	Children []*Node
}

func dfs(node *Node, yield func(*Node) bool) {
	if !yield(node) {
		return
	}
	for _, child := range node.Children {
		dfs(child, yield)
	}
}

func all(node *Node) iter.Seq[*Node] {
	return func(yield func(*Node) bool) {
		dfs(node, yield)
	}
}

func main() {
	root := Node{
		Data: "root",
		Children: []*Node{
			{
				Data: "hoge",
			},
			{
				Data: "fuga",
				Children: []*Node{
					{
						Data: "bar",
					},
				},
			},
			{
				Data: "foo",
			},
		},
	}
	for node := range all(&root) {
		fmt.Println(node.Data)
	}
}
