// Package modifier provides Tailwind-style visual modifiers for Gio widgets.
//
// These are composable decorators that can be applied to any widget:
//
//	modifier.Shadow(modifier.ShadowLg).Layout(gtx, myWidget)
//	modifier.Bg(theme.Primary).Layout(gtx, myWidget)
//	modifier.Rounded(theme.RoundedLg).Layout(gtx, myWidget)
package modifier

import (
	"image"
	"image/color"

	"gioui.org/layout"
	"gioui.org/op"
	"gioui.org/op/clip"
	"gioui.org/op/paint"
	"gioui.org/unit"
)

// Bg adds a background color to a widget.
type Bg struct {
	Color  color.NRGBA
	Radius unit.Dp
}

func (b Bg) Layout(gtx layout.Context, w layout.Widget) layout.Dimensions {
	macro := op.Record(gtx.Ops)
	dims := w(gtx)
	call := macro.Stop()

	rr := gtx.Dp(b.Radius)
	rrect := clip.UniformRRect(image.Rectangle{Max: dims.Size}, rr)
	defer rrect.Push(gtx.Ops).Pop()
	paint.ColorOp{Color: b.Color}.Add(gtx.Ops)
	paint.PaintOp{}.Add(gtx.Ops)

	call.Add(gtx.Ops)
	return dims
}
