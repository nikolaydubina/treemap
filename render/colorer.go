package render

import (
	"image/color"

	"github.com/nikolaydubina/treemap"
)

type NoneColorer struct{}

func (s NoneColorer) Color(tree treemap.Tree, node string) color.Color {
	return color.Transparent
}

// TreeHueColorer this algorithm will split Hue in NCL ranges such that
// deeper nodes have more specific hue.
// The advantage of this coloring is that nodes in that belong topologically close will have similar hue.
type TreeHueColorer struct{}

func (s TreeHueColorer) Color(tree treemap.Tree, node string) color.Color {
	return nil
}

// HeatColorer will use heat field of nodes.
// If not present, then will pick midrange.
// This is proxy for go-colorful palette.
type HeatColorer struct {
	Palette ColorfulPalette
}

func (s HeatColorer) Color(tree treemap.Tree, node string) color.Color {
	n, ok := tree.Nodes[node]
	if !ok {
		return s.Palette.GetInterpolatedColorFor(0.5)
	}
	return s.Palette.GetInterpolatedColorFor(n.Heat)
}
