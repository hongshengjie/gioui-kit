package theme

import (
	"image/color"

	"gioui.org/font"
)

// ---------- Font helpers ----------

// FontFace returns a font.Font with the given weight.
func FontFace(weight font.Weight) font.Font {
	return font.Font{Weight: weight}
}

var (
	FontThin      = FontFace(font.Thin)
	FontLight     = FontFace(font.Light)
	FontNormal    = FontFace(font.Normal)
	FontMedium    = FontFace(font.Medium)
	FontSemiBold  = FontFace(font.SemiBold)
	FontBold      = FontFace(font.Bold)
	FontExtraBold = FontFace(font.ExtraBold)
)

// ---------- Color utilities ----------

func rgb(hex uint32) color.NRGBA {
	return color.NRGBA{
		R: uint8(hex >> 16),
		G: uint8(hex >> 8),
		B: uint8(hex),
		A: 0xff,
	}
}

// WithAlpha returns a copy of c with the given alpha (0–255).
func WithAlpha(c color.NRGBA, a uint8) color.NRGBA {
	c.A = a
	return c
}

// Opacity returns a color with the given opacity fraction (0.0–1.0).
func Opacity(c color.NRGBA, opacity float32) color.NRGBA {
	c.A = uint8(float32(c.A) * opacity)
	return c
}

// Lerp linearly interpolates between two colors.
func Lerp(a, b color.NRGBA, t float32) color.NRGBA {
	return color.NRGBA{
		R: uint8(float32(a.R)*(1-t) + float32(b.R)*t),
		G: uint8(float32(a.G)*(1-t) + float32(b.G)*t),
		B: uint8(float32(a.B)*(1-t) + float32(b.B)*t),
		A: uint8(float32(a.A)*(1-t) + float32(b.A)*t),
	}
}

// RGB creates a color from hex value.
func RGB(hex uint32) color.NRGBA {
	return rgb(hex)
}
