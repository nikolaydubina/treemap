package treemap

// UIText is spec on how to render text.
type UIText struct {
	Text  string
	X     float64
	Y     float64
	H     float64
	W     float64
	Scale float64
}

// UIBox is spec on how to render a box. Could be Root.
type UIBox struct {
	Title    *UIText
	X        float64
	Y        float64
	W        float64
	H        float64
	Children []UIBox
}

func NewUIBox(node string, tree Tree, x, y, w, h, margin float64, padding float64) UIBox {
	t := UIBox{
		X: x + margin,
		Y: y + margin,
		W: w - (2 * margin),
		H: h - (2 * margin),
	}

	var textHeight float64
	if title := tree.Nodes[node].Name(); title != "" {
		// TODO: estimation of text length
		textHeight = 20
		t.Title = &UIText{
			Text:  title,
			X:     t.X + padding,
			Y:     t.Y + padding + textHeight,
			W:     t.W - (2 * padding),
			H:     textHeight,
			Scale: textHeight / 12,
		}
	}

	if len(tree.To[node]) == 0 {
		return t
	}

	areas := make([]float64, 0, len(tree.To[node]))
	for _, toPath := range tree.To[node] {
		areas = append(areas, nodeSize(tree, toPath))
	}

	childrenContainer := Box{
		X: t.X + padding,
		Y: t.Y + padding + textHeight,
		W: t.W - (2 * padding),
		H: t.H - (2 * padding) - textHeight,
	}
	boxes := Squarify(childrenContainer, areas)

	for i, toPath := range tree.To[node] {
		if boxes[i] == NilBox {
			continue
		}
		box := NewUIBox(
			toPath,
			tree,
			boxes[i].X,
			boxes[i].Y,
			boxes[i].W,
			boxes[i].H,
			margin,
			padding,
		)
		t.Children = append(t.Children, box)
	}

	return t
}

func nodeSize(tree Tree, node string) float64 {
	if n, ok := tree.Nodes[node]; ok {
		return n.Size
	}
	var s float64
	for _, child := range tree.To[node] {
		s += nodeSize(tree, child)
	}
	return s
}
