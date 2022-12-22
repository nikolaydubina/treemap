package render

import (
	"math"
	"testing"
)

func TestTextWidth(t *testing.T) {
	tests := []struct {
		text     string
		expWidth float64
	}{
		{
			text:     "abc",
			expWidth: 3.0,
		},
		{
			text:     "♠♧♡",
			expWidth: 3.0,
		},
	}

	for _, tc := range tests {
		t.Run(tc.text, func(t *testing.T) {
			if w := textWidth(tc.text, 1.25); math.Abs(tc.expWidth-w) > 0.0001 {
				t.Errorf("wrong text width: exp(%f) != got(%f)", tc.expWidth, w)
			}
		})
	}
}
