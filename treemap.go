package treemap

import "strings"

// for numerical stability
const minHeatDifferenceForHeatmap float64 = 0.0000001

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

func (t Tree) HasHeat() bool {
	minHeat, maxHeat := t.HeatRange()
	return (maxHeat - minHeat) > minHeatDifferenceForHeatmap
}

func (t Tree) HeatRange() (minHeat float64, maxHeat float64) {
	first := true
	for _, node := range t.Nodes {
		if first {
			minHeat = node.Heat
			maxHeat = node.Heat
			first = false
			continue
		}

		if node.Heat > maxHeat {
			maxHeat = node.Heat
		}
		if node.Heat < minHeat {
			minHeat = node.Heat
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
		if node.Heat == 0 {
			continue
		}

		n := Node{
			Path: node.Path,
			Size: node.Size,
			Heat: (node.Heat - minHeat) / (maxHeat - minHeat),
		}
		t.Nodes[path] = n
	}
}
