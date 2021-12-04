package treemap

import (
	"fmt"
	"log"
)

const (
	fontSize             int     = 12
	textHeightMultiplier float64 = 2
	textWidthMultiplier  float64 = 0.8
)

type SVGRenderer struct{}

func (r SVGRenderer) Render(box UIBox) []byte {
	s := fmt.Sprintf(`
		<svg 
			xmlns="http://www.w3.org/2000/svg" 
			xmlns:xlink="http://www.w3.org/1999/xlink" 
			viewBox="0 0 %f %f" 
			style="%s"
		>`,
		box.W,
		box.H,
		"background: white none repeat scroll 0% 0%;",
	)

	var q UIBox
	que := []UIBox{box}
	for len(que) > 0 {
		q, que = que[0], que[1:]
		que = append(que, q.Children...)

		box := fmt.Sprintf(
			`
			<g>
				<rect x="%f" y="%f" width="%f" height="%f" style="%s" />
				%s
			</g>
			`,
			q.X,
			q.Y,
			q.W,
			q.H,
			"fill: rgb(255, 255, 255);opacity:0.5;fill-opacity:1;stroke:grey;stroke-width:1px;stroke-opacity:1;",
			TextSVG(q.Title),
		)
		s += box + "\n"
	}

	s += `</svg>`

	return []byte(s)
}

func textWidth(text string, fontSize float64) float64 {
	return fontSize * float64(len(text)) * textWidthMultiplier
}

func textHeight(text string, fontSize float64) float64 {
	return fontSize * textHeightMultiplier
}

func TextSVG(t *UIText) string {
	if t == nil {
		return ""
	}

	w := textWidth(t.Text, float64(fontSize))
	h := textHeight(t.Text, float64(fontSize))
	log.Printf("%#v %#v\n", w, h)

	// compute bounding and scale

	s := fmt.Sprintf(`
		<text 
			data-notex="1" 
			text-anchor="start"
			transform="translate(%f,%f) scale(%f)"
			style="%s" 
			data-math="N">%s</text>
		`,
		t.X,
		t.Y,
		t.Scale,
		fmt.Sprintf("font-size: %dpx; fill: rgb(68, 68, 68); fill-opacity: 1; white-space: pre;", fontSize),
		t.Text,
	)
	return s
}
