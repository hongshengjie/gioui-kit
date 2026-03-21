package scaffold

import (
	"image/color"

	"gioui.org/font"
	"gioui.org/font/gofont"
	"gioui.org/layout"
	"gioui.org/op"
	"gioui.org/op/paint"
	"gioui.org/text"
	"gioui.org/unit"
	"gioui.org/widget"

	"github.com/hongshengjie/gioui-kit/theme"
)

func drawLabel(gtx layout.Context, th *theme.Theme, txt string, col color.NRGBA, size unit.Sp, weight font.Weight) layout.Dimensions {
	lbl := widget.Label{MaxLines: 1}
	f := font.Font{Weight: weight}
	paint.ColorOp{Color: col}.Add(gtx.Ops)
	if th.Shaper == nil {
		th.Shaper = text.NewShaper(text.NoSystemFonts(), text.WithCollection(gofont.Collection()))
	}
	return lbl.Layout(gtx, th.Shaper, f, size, txt, op.CallOp{})
}
