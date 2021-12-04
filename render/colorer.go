package render

import (
	"image/color"

	"github.com/lucasb-eyer/go-colorful"
	"github.com/nikolaydubina/treemap"
)

var (
	DarkTextColor  color.Color = color.Black
	LightTextColor color.Color = color.White
)

type NoneColorer struct{}

func (s NoneColorer) ColorBox(tree treemap.Tree, node string) color.Color {
	return color.Transparent
}

func (s NoneColorer) ColorText(tree treemap.Tree, node string) color.Color {
	return DarkTextColor
}

// TreeHueColorer this algorithm will split Hue in NCL ranges such that
// deeper nodes have more specific hue.
// The advantage of this coloring is that nodes in that belong topologically close will have similar hue.
type TreeHueColorer struct{}

func (s TreeHueColorer) ColorBox(tree treemap.Tree, node string) color.Color {
	return color.Transparent
}

func (s TreeHueColorer) ColorText(tree treemap.Tree, node string) color.Color {
	return DarkTextColor
}

// HeatColorer will use heat field of nodes.
// If not present, then will pick midrange.
// This is proxy for go-colorful palette.
type HeatColorer struct {
	Palette ColorfulPalette
}

func (s HeatColorer) ColorBox(tree treemap.Tree, node string) color.Color {
	n, ok := tree.Nodes[node]
	if !ok {
		return s.Palette.GetInterpolatedColorFor(0.5)
	}
	return s.Palette.GetInterpolatedColorFor(n.Heat)
}

func (s HeatColorer) ColorText(tree treemap.Tree, node string) color.Color {
	boxColor := s.ColorBox(tree, node).(colorful.Color)
	_, _, l := boxColor.Hcl()
	switch {
	case l > 0.5:
		return DarkTextColor
	default:
		return LightTextColor
	}
}
