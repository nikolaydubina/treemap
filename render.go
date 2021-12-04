package treemap

type UIBox struct {
	Title      string
	X          float64
	Y          float64
	W          float64
	H          float64
	TextHeight float64 // ratio of height we leave for text
	Padding    float64 // ratio of height and width we spend for padding node contents (including text and inner boxes)
	Children   []UIBox
}

func NewUIBox(node string, tree Tree, x, y, w, h, textHeight, padding float64) UIBox {
	t := UIBox{
		Title:      tree.Nodes[node].Name(),
		X:          x,
		Y:          y,
		W:          w,
		H:          h,
		TextHeight: textHeight,
		Padding:    padding,
	}

	if len(tree.To[node]) == 0 {
		return t
	}

	areas := make([]float64, 0, len(tree.To[node]))
	for _, toPath := range tree.To[node] {
		areas = append(areas, nodeSize(tree, toPath))
	}

	childrenContainer := Box{
		X: x + (w * padding),
		Y: y + (h * padding) + (h * textHeight),
		W: w * (1 - (2 * padding)),
		H: h * (1 - (2 * padding) - textHeight),
	}
	boxes := Squarify(childrenContainer, areas)

	for i, toPath := range tree.To[node] {
		box := NewUIBox(
			toPath,
			tree,
			boxes[i].X,
			boxes[i].Y,
			boxes[i].W,
			boxes[i].H,
			textHeight,
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
