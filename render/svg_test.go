package render

import (
	"image/color"
	"testing"
)

func TestSVGRender(t *testing.T) {
	tests := []struct {
		name   string
		root   UIBox
		w      float64
		h      float64
		expSVG string
	}{
		{
			name: "basic",
			root: UIBox{
				Title: &UIText{
					Text:  "123",
					X:     1,
					Y:     2,
					W:     3,
					H:     4,
					Scale: 10,
					Color: color.Black,
				},
				X:           1,
				Y:           2,
				W:           3,
				H:           4,
				Children:    []UIBox{},
				IsRoot:      true,
				Color:       color.Black,
				BorderColor: color.Opaque,
			},
			w: 10,
			h: 10,
			expSVG: `
<svg 
	xmlns="http://www.w3.org/2000/svg" 
	xmlns:xlink="http://www.w3.org/1999/xlink" 
	viewBox="0 0 10.000000 10.000000" 
	style="background: white none repeat scroll 0% 0%;"
>
<g>
	<rect x="1.000000" y="2.000000" width="3.000000" height="4.000000" style="fill: rgba(0, 0, 0, 65535);opacity:1;fill-opacity:1;stroke:rgba(255,255,255,65535);stroke-width:1px;stroke-opacity:1;" />
	
<text 
	data-notex="1" 
	text-anchor="start"
	transform="translate(1.000000,6.000000) scale(10.000000)"
	style="font-family: Open Sans, verdana, arial, sans-serif !important; font-size: 12px; fill: rgb(0, 0, 0, 65535); fill-opacity: 1; white-space: pre;" 
	data-math="N">123</text>
		
</g>

</svg>`,
		},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			out := SVGRenderer{}.Render(tc.root, tc.w, tc.h)
			if string(out) != tc.expSVG {
				t.Errorf("expected(%s) != got(%s)", tc.expSVG, out)
			}
		})
	}
}
