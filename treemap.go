package treemap

import "strings"

// for numerical stability
const minHeatDifferenceForHeatmap float64 = 0.0000001

type Node struct {
	Path    string
	Size    float64
	Heat    float64
	HasHeat bool
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
			Size:    node.Size,
			Heat:    (node.Heat - minHeat) / (maxHeat - minHeat),
			HasHeat: true,
		}
		t.Nodes[path] = n
	}
}
