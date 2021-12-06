package main

import (
	"flag"
	"fmt"
	"image/color"
	"io"
	"log"
	"os"

	"github.com/nikolaydubina/treemap"
	"github.com/nikolaydubina/treemap/parser"
	"github.com/nikolaydubina/treemap/render"
)

const doc string = `
Generate treemaps from STDIN in header-less CSV.

</ delimitered path>,<size>,<heat>

Example:

$ echo '
Africa/Algeria,33333216,72
Africa/Angola,12420476,42
Africa/Benin,8078314,56
' | treemap > out.svg

Command options:
`

func main() {
	var (
		w           float64
		h           float64
		marginBox   float64
		paddingBox  float64
		padding     float64
		colorScheme string
		colorBorder string
	)

	flag.Usage = func() {
		fmt.Fprint(flag.CommandLine.Output(), doc)
		flag.PrintDefaults()
	}
	flag.Float64Var(&w, "w", 1028, "width of output")
	flag.Float64Var(&h, "h", 640, "height of output")
	flag.Float64Var(&marginBox, "margin-box", 4, "margin between boxes")
	flag.Float64Var(&paddingBox, "padding-box", 4, "padding between box border and content")
	flag.Float64Var(&padding, "padding", 32, "padding around root content")
	flag.StringVar(&colorScheme, "color", "balance", "color scheme (RdBu, balance, none)")
	flag.StringVar(&colorBorder, "color-border", "white", "color of borders (white, black)")
	flag.Parse()

	in, err := io.ReadAll(os.Stdin)
	if err != nil {
		log.Fatal(err)
	}

	parser := parser.CSVTreeParser{}
	tree, err := parser.ParseString(string(in))
	if err != nil || tree == nil {
		log.Fatal(err)
	}

	sizeImputer := treemap.SumSizeImputer{EmptyLeafSize: 1}
	sizeImputer.ImputeSize(*tree)

	heatImputer := treemap.WeightedHeatImputer{EmptyLeafHeat: 0.5}
	heatImputer.ImputeHeat(*tree)

	tree.NormalizeHeat()

	var colorer render.Colorer

	palette, hasPalette := render.GetPalette(colorScheme)
	treeHueColorer := render.TreeHueColorer{
		Offset: 0,
		Hues:   map[string]float64{},
		C:      0.5,
		L:      0.5,
		DeltaH: 15,
		DeltaC: 0.1,
		DeltaL: 0.4,
	}

	switch {
	case colorScheme == "none":
		colorer = render.NoneColorer{}
	case colorScheme == "balanced":
		colorer = treeHueColorer
	case hasPalette && tree.HasHeat():
		colorer = render.HeatColorer{Palette: palette}
	case tree.HasHeat():
		palette, _ := render.GetPalette("RdBu")
		colorer = render.HeatColorer{Palette: palette}
	default:
		colorer = treeHueColorer
	}

	borderColor := color.White
	switch {
	case colorScheme == "none":
		borderColor = color.Black
	case colorBorder == "white":
		borderColor = color.White
	case colorBorder == "black":
		borderColor = color.Black
	default:
		borderColor = color.White
	}

	uiBuilder := render.UITreeMapBuilder{
		Colorer:     colorer,
		BorderColor: borderColor,
	}
	spec := uiBuilder.NewUITreeMap(*tree, w, h, marginBox, paddingBox, padding)
	renderer := render.SVGRenderer{}

	os.Stdout.Write(renderer.Render(spec, w, h))
}
