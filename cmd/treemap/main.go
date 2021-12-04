package main

import (
	"flag"
	"fmt"
	"io"
	"log"
	"os"

	"github.com/nikolaydubina/treemap/parser"
	"github.com/nikolaydubina/treemap/render"
)

const doc string = `
Generate treemaps from STDIN in header-less CSV as follows:

</ delimitered path>,<size>,<heat>

Example:

'''csv
Africa/Algeria,33333216.0,72.301
Africa/Angola,12420476.0,42.731
Africa/Benin,8078314.0,56.728
'''

$ cat file.csv | treemap > out.svg
`

func main() {
	var (
		w       float64
		h       float64
		margin  float64
		padding float64
	)

	flag.Usage = func() {
		fmt.Fprint(flag.CommandLine.Output(), doc)
		flag.PrintDefaults()
	}
	flag.Float64Var(&w, "w", 1028, "width of output")
	flag.Float64Var(&h, "h", 1028, "height of output")
	flag.Float64Var(&margin, "margin", 5, "margin between boxes")
	flag.Float64Var(&padding, "padding", 5, "padding between box border and content")
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

	spec := render.NewUIBox(tree.Root, *tree, 0, 0, w, h, margin, padding)
	renderer := render.SVGRenderer{}

	os.Stdout.Write(renderer.Render(spec))
}
