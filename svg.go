package treemap

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
				<g transform="translate(%f,%f)">
					<text 
						data-notex="1" 
						class="slicetext" 
						text-anchor="start"
						style="%s" 
						data-math="N">%s</text>
				</g>
			</g>
			`,
			q.X,
			q.Y,
			q.W,
			q.H,
			"fill:blue;opacity:0.5;fill-opacity:1;stroke:grey;stroke-width:1px;stroke-opacity:1;",
			q.X,
			q.Y,
			"font-size: 12px; fill: rgb(68, 68, 68); fill-opacity: 1; white-space: pre;",
			q.Title,
		)
		s += box + "\n"
	}

	s += `</svg>`

	return []byte(s)
}
