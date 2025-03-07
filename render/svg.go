package render

import (
	"fmt"
	"image/color"
	"strings"
)

var xmlEscaper = strings.NewReplacer(
	"&", "&amp;",
	"<", "&lt;",
	">", "&gt;",
	"\"", "&quot;",
	"'", "&apos;",
)

type SVGRenderer struct{}

func (r SVGRenderer) Render(root UIBox, w, h float64) []byte {
	if !root.IsRoot {
		return nil
	}

	s := fmt.Sprintf(`
<svg 
	xmlns="http://www.w3.org/2000/svg" 
	xmlns:xlink="http://www.w3.org/1999/xlink" 
	viewBox="0 0 %f %f" 
	style="%s"
>`,
		w,
		h,
		"background: white none repeat scroll 0% 0%;",
	)

	var q UIBox
	que := []UIBox{root}
	for len(que) > 0 {
		q, que = que[0], que[1:]
		que = append(que, q.Children...)
		s += BoxSVG(q) + "\n"
	}

	s += `</svg>`

	return []byte(s)
}

func BoxSVG(q UIBox) string {
	if q.IsInvisible {
		return ""
	}
	r, g, b, a := color.White.RGBA()
	if q.Color != color.Opaque {
		r, g, b, a = q.Color.RGBA()
	}

	r = r >> 8
	g = g >> 8
	b = b >> 8
	o := float64(a>>8) / 255.0

	br, bg, bb, ba := color.White.RGBA()
	if q.BorderColor != color.Opaque {
		br, bg, bb, ba = q.BorderColor.RGBA()
	}
	br = br >> 8
	bg = bg >> 8
	bb = bb >> 8
	bo := float64(ba>>8) / 255.0

	return fmt.Sprintf(`
<g>
	<rect x="%f" y="%f" width="%f" height="%f" style="%s" />
	%s
</g>
`,
		q.X,
		q.Y,
		q.W,
		q.H,
		fmt.Sprintf("fill: rgb(%d, %d, %d);opacity:1;fill-opacity:%.2f;stroke:rgb(%d,%d,%d);stroke-width:1px;stroke-opacity:%.2f;", r, g, b, o, br, bg, bb, bo),
		TextSVG(q.Title),
	)
}

func TextSVG(t *UIText) string {
	if t == nil {
		return ""
	}

	r, g, b, a := color.Black.RGBA()
	if t.Color != color.Opaque {
		r, g, b, a = t.Color.RGBA()
	}

	r = r >> 8
	g = g >> 8
	b = b >> 8
	o := float64(a>>8) / 255.0

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
		fmt.Sprintf("font-family: Open Sans, verdana, arial, sans-serif !important; font-size: %dpx; fill: rgb(%d, %d, %d); fill-opacity: %.2f; white-space: pre;", fontSize, r, g, b, o),
		xmlEscaper.Replace(t.Text),
	)
	return s
}
