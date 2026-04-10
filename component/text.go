package component

import (
	"image"
	"image/color"

	"gioui.org/font"
	"gioui.org/font/gofont"
	"gioui.org/layout"
	"gioui.org/op"
	"gioui.org/op/paint"
	"gioui.org/text"
	"gioui.org/unit"
	"gioui.org/widget"

	"github.com/ossprovider/gioui-kit/theme"
)

// Text renders themed text (like Tailwind text-sm, text-lg, font-bold, etc.).
type Text struct {
	Content string
	Color   color.NRGBA
	Size    unit.Sp
	Weight  font.Weight
	th      *theme.Theme
}

func NewText(th *theme.Theme, content string) *Text {
	return &Text{
		Content: content,
		Color:   th.BaseContent,
		Size:    th.FontSize,
		Weight:  font.Normal,
		th:      th,
	}
}

func (t *Text) H1() *Text   { t.Size = t.th.H1Size; t.Weight = font.Bold; return t }
func (t *Text) H2() *Text   { t.Size = t.th.H2Size; t.Weight = font.Bold; return t }
func (t *Text) H3() *Text   { t.Size = t.th.H3Size; t.Weight = font.SemiBold; return t }
func (t *Text) H4() *Text   { t.Size = t.th.H4Size; t.Weight = font.SemiBold; return t }
func (t *Text) Sm() *Text   { t.Size = t.th.SmSize; return t }
func (t *Text) Xs() *Text   { t.Size = t.th.XsSize; return t }
func (t *Text) Bold() *Text { t.Weight = font.Bold; return t }
func (t *Text) WithColor(c color.NRGBA) *Text { t.Color = c; return t }

func (t *Text) Layout(gtx layout.Context) layout.Dimensions {
	return drawText(gtx, t.th, t.Content, t.Color, t.Size, t.Weight)
}

// drawText is a shared text drawing utility.
func drawText(gtx layout.Context, th *theme.Theme, txt string, col color.NRGBA, size unit.Sp, weight font.Weight) layout.Dimensions {
	lbl := widget.Label{MaxLines: 0}
	f := font.Font{Weight: weight}
	paint.ColorOp{Color: col}.Add(gtx.Ops)
	if th.Shaper == nil {
		th.Shaper = text.NewShaper(text.NoSystemFonts(), text.WithCollection(defaultFonts()))
	}
	return lbl.Layout(gtx, th.Shaper, f, size, txt, op.CallOp{})
}

// drawSpinner draws a simple loading spinner placeholder.
func drawSpinner(gtx layout.Context, col color.NRGBA, size unit.Sp) layout.Dimensions {
	sz := gtx.Sp(size)
	return layout.Dimensions{Size: image.Pt(int(sz), int(sz))}
}

// defaultFonts returns the built-in Go font collection.
func defaultFonts() []font.FontFace {
	return gofont.Collection()
}
