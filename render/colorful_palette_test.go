package render

import (
	"testing"

	"github.com/lucasb-eyer/go-colorful"
)

func TestGetPaletteBasic(t *testing.T) {
	palette, _ := GetPalette("RdBu")

	expColor := colorful.Color{R: 0.403921568627451, G: 0, B: 0.12156862745098039}
	if palette[0].Col != expColor {
		t.Errorf("exp(%#v) != got(%#v)", expColor, palette[0].Col)
	}
}
