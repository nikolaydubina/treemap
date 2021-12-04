package parser

import (
	"encoding/csv"
	"errors"
	"fmt"
	"io"
	"strconv"
	"strings"

	"github.com/nikolaydubina/treemap"
)

type CSVTreeParser struct{}

func (s CSVTreeParser) ParseString(in string) (*treemap.Tree, error) {
	nodes, err := parseNodes(in)
	if err != nil {
		return nil, fmt.Errorf("can not parse nodes: %w", err)
	}

	tree, err := makeTree(nodes)
	if err != nil {
		return nil, fmt.Errorf("can not make tree: %w", err)
	}

	fillSizeIfAllEmpty(*tree)

	return tree, nil
}

func fillSizeIfAllEmpty(tree treemap.Tree) {
	allZeroSize := true
	for _, node := range tree.Nodes {
		if node.Size > 0 {
			allZeroSize = false
			break
		}
	}
	if allZeroSize {
		for i, node := range tree.Nodes {
			tree.Nodes[i] = treemap.Node{
				Path: node.Path,
				Size: 1,
				Heat: node.Heat,
			}
		}
	}
}

func parseNodes(in string) ([]treemap.Node, error) {
	var nodes []treemap.Node
	r := csv.NewReader(strings.NewReader(in))
	for {
		record, err := r.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			return nil, fmt.Errorf("can not parse: %w", err)
		}

		if len(record) == 0 {
			return nil, errors.New("no values in row")
		}

		node := treemap.Node{Path: record[0]}

		if len(record) >= 2 {
			v, err := strconv.ParseFloat(record[1], 64)
			if err != nil {
				return nil, fmt.Errorf("size(%s) is not float: %w", record[1], err)
			}
			node.Size = v
		}

		if len(record) >= 3 {
			v, err := strconv.ParseFloat(record[2], 64)
			if err != nil {
				return nil, fmt.Errorf("heat(%s) is not float: %w", record[2], err)
			}
			node.Heat = v
		}

		nodes = append(nodes, node)
	}
	return nodes, nil
}

// If node is in path, but not present, then it will be in To but not will have entry in Nodes.
// This is not terribly efficient, but should do its job for small graphs.
func makeTree(nodes []treemap.Node) (*treemap.Tree, error) {
	tree := treemap.Tree{
		Nodes: map[string]treemap.Node{},
		To:    map[string][]string{},
	}

	// for finding roots
	hasParent := map[string]bool{}

	for _, node := range nodes {
		if _, ok := tree.Nodes[node.Path]; ok {
			return nil, fmt.Errorf("duplicate node(%s)", node.Path)
		}
		tree.Nodes[node.Path] = node

		parts := strings.Split(node.Path, "/")
		hasParent[parts[0]] = false

		for parent, i := parts[0], 1; i < len(parts); i++ {
			child := parent + "/" + parts[i]

			tree.To[parent] = append(tree.To[parent], child)
			hasParent[child] = true

			parent = child
		}
	}

	for node, v := range tree.To {
		tree.To[node] = unique(v)
	}

	var roots []string
	for node, has := range hasParent {
		if !has {
			roots = append(roots, node)
		}
	}

	switch {
	case len(roots) == 0:
		return nil, errors.New("no roots, possible cycle in graph")
	case len(roots) > 1:
		tree.Root = "<some-secret-string>"
		tree.To[tree.Root] = roots
	default:
		tree.Root = roots[0]
	}

	return &tree, nil
}

func unique(a []string) []string {
	u := map[string]bool{}
	var b []string
	for _, q := range a {
		if _, ok := u[q]; !ok {
			u[q] = true
			b = append(b, q)
		}
	}
	return b
}
