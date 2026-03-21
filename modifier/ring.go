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

// Ring draws an outline ring around a widget (like Tailwind `ring-2 ring-blue-500`).
type Ring struct {
	Width  unit.Dp
	Color  color.NRGBA
	Offset unit.Dp
	Radius unit.Dp
}

func (r Ring) Layout(gtx layout.Context, w layout.Widget) layout.Dimensions {
	macro := op.Record(gtx.Ops)
	dims := w(gtx)
	call := macro.Stop()

	call.Add(gtx.Ops)

	// Draw ring
	rw := gtx.Dp(r.Width)
	ro := gtx.Dp(r.Offset)
	rr := gtx.Dp(r.Radius)

	ringRect := image.Rectangle{
		Min: image.Pt(-ro, -ro),
		Max: image.Pt(dims.Size.X+ro, dims.Size.Y+ro),
	}

	paint.FillShape(gtx.Ops, r.Color,
		clip.Stroke{
			Path:  clip.UniformRRect(ringRect, rr+ro).Path(gtx.Ops),
			Width: float32(rw),
		}.Op(),
	)

	return dims
}
