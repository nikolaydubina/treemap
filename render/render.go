package render

import (
	"github.com/nikolaydubina/treemap"
	"github.com/nikolaydubina/treemap/layout"
)

const (
	fontSize             int     = 12
	textHeightMultiplier float64 = 0.8
	textWidthMultiplier  float64 = 0.8
	tooSmallBoxHeight    float64 = 5
	tooSmallBoxWidth     float64 = 5
)

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

func (f UIBox) IsEmpty() bool {
	return f.W == 0 || f.H == 0
}

func NewUIBox(node string, tree treemap.Tree, x, y, w, h, margin float64, padding float64) UIBox {
	if (w <= (2 * padding)) || (h <= (2 * padding)) || w < tooSmallBoxWidth || h < tooSmallBoxHeight {
		// too small, do not render
		return UIBox{}
	}

	t := UIBox{
		X: x + margin,
		Y: y + margin,
		W: w - (2 * margin),
		H: h - (2 * margin),
	}

	var textHeight float64
	if title := tree.Nodes[node].Name(); title != "" {
		// fit text
		w := t.W - (2 * padding)
		h := t.H - (2 * padding)
		if scale, th := fitText(title, fontSize, w); scale > 0 && th > 0 && th < h {
			textHeight = th
			// if enough space for text, then add
			t.Title = &UIText{
				Text:  title,
				X:     t.X + padding,
				Y:     t.Y + padding,
				W:     w,
				H:     textHeight,
				Scale: scale,
			}
		}
	}

	if len(tree.To[node]) == 0 {
		return t
	}

	areas := make([]float64, 0, len(tree.To[node]))
	for _, toPath := range tree.To[node] {
		areas = append(areas, nodeSize(tree, toPath))
	}

	childrenContainer := layout.Box{
		X: t.X + padding,
		Y: t.Y + padding + textHeight,
		W: t.W - (2 * padding),
		H: t.H - (2 * padding) - textHeight,
	}
	boxes := layout.Squarify(childrenContainer, areas)

	for i, toPath := range tree.To[node] {
		if boxes[i] == layout.NilBox {
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
		if box.IsEmpty() {
			continue
		}
		t.Children = append(t.Children, box)
	}

	return t
}

func nodeSize(tree treemap.Tree, node string) float64 {
	if n, ok := tree.Nodes[node]; ok {
		return n.Size
	}
	var s float64
	for _, child := range tree.To[node] {
		s += nodeSize(tree, child)
	}
	return s
}

// compute scale to fit worst dimension
func fitText(text string, fontSize int, W float64) (scale float64, h float64) {
	w := textWidth(text, float64(fontSize))
	h = textHeight(text, float64(fontSize))

	scale = 1.0
	if wscale := W / w; wscale < scale {
		scale = wscale
	}

	H := textHeight(text, float64(fontSize))
	if hscale := H / h; hscale < scale {
		scale = hscale
	}

	return scale, h
}

func textWidth(text string, fontSize float64) float64 {
	return fontSize * float64(len(text)) * textWidthMultiplier
}

func textHeight(text string, fontSize float64) float64 {
	return fontSize * textHeightMultiplier
}
