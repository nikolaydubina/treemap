package main

import (
	"flag"
	"fmt"
	"io"
	"log"
	"os"

	"github.com/nikolaydubina/treemap"
)

const doc string = `
TODO

TODO 1
`

func main() {
	var (
		w float64
		h float64
	)

	flag.Usage = func() {
		fmt.Fprintf(flag.CommandLine.Output(), doc, os.Args[0])
		flag.PrintDefaults()
	}
	flag.Float64Var(&w, "w", 1028, "width of output")
	flag.Float64Var(&h, "h", 1028, "height of output")
	flag.Parse()

	in, err := io.ReadAll(os.Stdin)
	if err != nil {
		log.Fatal(err)
	}

	parser := treemap.CSVTreeParser{}
	tree, err := parser.ParseString(string(in))
	if err != nil || tree == nil {
		log.Fatal(err)
	}

	spec := treemap.NewUIBox(tree.Root, *tree, 0, 0, w, h, 0.2, 0.05)
	renderer := treemap.SVGRenderer{}

	os.Stdout.Write(renderer.Render(spec))
}
