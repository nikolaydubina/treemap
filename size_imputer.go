package treemap

// SumSizeImputer will set sum of children into empty parents and fill children with contant.
type SumSizeImputer struct {
	EmptyLeafSize float64
}

func (s SumSizeImputer) ImputeSize(t Tree) {
	s.ImputeSizeNode(t, t.Root)
}

func (s SumSizeImputer) ImputeSizeNode(t Tree, node string) {
	var sum float64
	for _, child := range t.To[node] {
		s.ImputeSizeNode(t, child)
		sum += t.Nodes[child].Size
	}

	if n, ok := t.Nodes[node]; !ok || n.Size == 0 {
		v := s.EmptyLeafSize
		if len(t.To[node]) > 0 {
			v = sum
		}

		t.Nodes[node] = Node{
			Path: node,
			Size: v,
		}
	}
}
