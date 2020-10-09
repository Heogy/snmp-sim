package main

import (
	"fmt"
	"strconv"
)

type OidNode struct {
	id     int
	value  string
	childs []OidNode
}



func main() {
	root := initData()
	walk := root.walk()
	for _, s := range walk {
		fmt.Printf(s)
	}
}

func (node OidNode) walk() []string {
	return node.walkPrefix("")
}

func (node OidNode) walkPrefix(prefix string) []string{

	var result []string
	for _, child := range node.childs {
		result = append(result, child.walkPrefix(prefix + "." + strconv.Itoa(node.id))...)
	}
	return result
}

func initData() OidNode {

	var child214 = OidNode{
		id:    4,
		value: "the child 214",
	}

	var child21 = OidNode{
		id:     1,
		value:  "the child 21",
		childs: []OidNode{child214},
	}

	var child2 = OidNode{
		id:     2,
		value:  "the child 2",
		childs: []OidNode{child21},
	}
	var child1 = OidNode{
		id:    1,
		value: "the child 1",
	}
	var root = OidNode{
		id:     1,
		value:  "the root",
		childs: []OidNode{child1, child2},
	}
	return root
}
