package treemap

import "strings"

type Node struct {
	Path string
	Size float64
	Heat float64
}

func (n Node) Name() string {
	parts := strings.Split(n.Path, "/")
	if len(parts) == 0 {
		return ""
	}
	return parts[len(parts)-1]
}

type Tree struct {
	Nodes map[string]Node
	To    map[string][]string
	Root  string
}
