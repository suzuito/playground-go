package main

import "fmt"

type Node struct {
	Data     string
	Children []*Node
}

func dfs(node *Node, visit func(*Node)) {
	visit(node)
	for _, child := range node.Children {
		dfs(child, visit)
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
	dfs(&root, func(n *Node) {
		fmt.Println(n.Data)
	})
}
