package treemap

import "strings"

// for numerical stability
const minHeatDifferenceForHeatmap float64 = 0.0000001

type Node struct {
	Path    string
	Name    string
	Size    float64
	Heat    float64
	HasHeat bool
}

type Tree struct {
	Nodes map[string]Node
	To    map[string][]string
	Root  string
}

func (t Tree) HasHeat() bool {
	minHeat, maxHeat := t.HeatRange()
	return (maxHeat - minHeat) > minHeatDifferenceForHeatmap
}

func (t Tree) HeatRange() (minHeat float64, maxHeat float64) {
	first := true
	for _, node := range t.Nodes {
		if !node.HasHeat {
			continue
		}
		h := node.Heat

		if first {
			minHeat = h
			maxHeat = h
			first = false
			continue
		}

		if h > maxHeat {
			maxHeat = h
		}
		if h < minHeat {
			minHeat = h
		}
	}
	return minHeat, maxHeat
}

func (t Tree) NormalizeHeat() {
	minHeat, maxHeat := t.HeatRange()

	if (maxHeat - minHeat) < minHeatDifferenceForHeatmap {
		return
	}

	for path, node := range t.Nodes {
		if !node.HasHeat {
			continue
		}

		n := Node{
			Path:    node.Path,
			Name:    node.Name,
			Size:    node.Size,
			Heat:    (node.Heat - minHeat) / (maxHeat - minHeat),
			HasHeat: true,
		}
		t.Nodes[path] = n
	}
}

// SetNamesFromPaths will update each node to its path leaf as name.
func SetNamesFromPaths(t *Tree) {
	if t == nil {
		return
	}
	for path, node := range t.Nodes {
		parts := strings.Split(node.Path, "/")
		if len(parts) == 0 {
			continue
		}

		t.Nodes[path] = Node{
			Path:    node.Path,
			Name:    parts[len(parts)-1],
			Size:    node.Size,
			Heat:    node.Heat,
			HasHeat: node.HasHeat,
		}
	}
}

// CollapseRoot will remove root nodes up to first child that have multiple children.
// Will set name of this node to joined path from roots.
// Will set size and heat to this child's size and heat.
func CollapseRoot(t *Tree) {
	if t == nil {
		return
	}

	q := t.Root
	for children := t.To[q]; len(children) == 1; {
		q = children[0]
		children = t.To[q]
	}

	t.Root = q

	node := t.Nodes[q]
	t.Nodes[q] = Node{
		Path:    node.Path,
		Name:    node.Path,
		Size:    node.Size,
		Heat:    node.Heat,
		HasHeat: node.HasHeat,
	}
}
