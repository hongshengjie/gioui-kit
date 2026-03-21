package layout

import (
	"image"
	"image/color"

	"gioui.org/layout"
	"gioui.org/op/clip"
	"gioui.org/op/paint"
	"gioui.org/unit"
)

// Center centers a widget within its parent constraints.
func Center(gtx layout.Context, w layout.Widget) layout.Dimensions {
	return layout.Center.Layout(gtx, w)
}

// DividerH draws a horizontal divider line.
type DividerH struct {
	Color     color.NRGBA
	Thickness unit.Dp
	Inset     layout.Inset
}

func (d DividerH) Layout(gtx layout.Context) layout.Dimensions {
	thickness := gtx.Dp(d.Thickness)
	if thickness <= 0 {
		thickness = 1
	}
	return d.Inset.Layout(gtx, func(gtx layout.Context) layout.Dimensions {
		sz := image.Pt(gtx.Constraints.Max.X, thickness)
		paint.FillShape(gtx.Ops, d.Color,
			clip.Rect{Max: sz}.Op(),
		)
		return layout.Dimensions{Size: sz}
	})
}

// SpaceH creates a horizontal spacer (width).
func SpaceH(dp unit.Dp) layout.Widget {
	return func(gtx layout.Context) layout.Dimensions {
		return layout.Dimensions{Size: image.Pt(gtx.Dp(dp), 0)}
	}
}

// SpaceV creates a vertical spacer (height).
func SpaceV(dp unit.Dp) layout.Widget {
	return func(gtx layout.Context) layout.Dimensions {
		return layout.Dimensions{Size: image.Pt(0, gtx.Dp(dp))}
	}
}
