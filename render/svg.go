package render

import (
	"fmt"
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

func TextSVG(t *UIText) string {
	if t == nil {
		return ""
	}

	s := fmt.Sprintf(`
		<text 
			data-notex="1" 
			text-anchor="start"
			transform="translate(%f,%f) scale(%f)"
			style="%s" 
			data-math="N">%s</text>
		`,
		t.X,
		t.Y+t.H,
		t.Scale,
		fmt.Sprintf("font-family: Open Sans, sans-serif !important; font-size: %dpx; fill: rgb(68, 68, 68); fill-opacity: 1; white-space: pre;", fontSize),
		t.Text,
	)
	return s
}
