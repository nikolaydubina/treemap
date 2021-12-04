package render

import (
	"image/color"

	"github.com/lucasb-eyer/go-colorful"
)

// This table contains the "keypoints" of the colorgradient you want to generate.
// The position of each keypoint has to live in the range [0,1]
// Ths is copied from go-colorful examples!!!
type ColorfulPalette []struct {
	Col colorful.Color
	Pos float64
}

// This is the meat of the gradient computation. It returns a HCL-blend between
// the two colors around `t`.
// Note: It relies heavily on the fact that the gradient keypoints are sorted.
func (gt ColorfulPalette) GetInterpolatedColorFor(t float64) color.Color {
	for i := 0; i < len(gt)-1; i++ {
		c1 := gt[i]
		c2 := gt[i+1]
		if c1.Pos <= t && t <= c2.Pos {
			// We are in between c1 and c2. Go blend them!
			t := (t - c1.Pos) / (c2.Pos - c1.Pos)
			return c1.Col.BlendHcl(c2.Col, t).Clamped()
		}
	}

	// Nothing found? Means we're at (or past) the last gradient keypoint.
	return gt[len(gt)-1].Col
}

func MustParseHex(s string) colorful.Color {
	c, err := colorful.Hex(s)
	if err != nil {
		panic("MustParseHex: " + err.Error())
	}
	return c
}

func GetPalette(name string) (ColorfulPalette, bool) {
	switch name {
	case "RdBu":
		p := ColorfulPalette{
			{MustParseHex("#67001f"), 0.00},
			{MustParseHex("#b2182b"), 0.10},
			{MustParseHex("#d6604d"), 0.20},
			{MustParseHex("#f4a482"), 0.30},
			{MustParseHex("#fddbc7"), 0.40},
			{MustParseHex("#f7f7f7"), 0.50},
			{MustParseHex("#d1e5f0"), 0.60},
			{MustParseHex("#92c5de"), 0.70},
			{MustParseHex("#4393c3"), 0.80},
			{MustParseHex("#2166ac"), 0.90},
			{MustParseHex("#053061"), 1.00},
		}
		return p, true
	default:
		return nil, false
	}
}
