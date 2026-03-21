package layout

import (
	"gioui.org/layout"
	"gioui.org/unit"
)

// W sets a fixed width constraint.
func W(dp unit.Dp) func(gtx layout.Context, w layout.Widget) layout.Dimensions {
	return func(gtx layout.Context, w layout.Widget) layout.Dimensions {
		px := gtx.Dp(dp)
		gtx.Constraints.Min.X = px
		gtx.Constraints.Max.X = px
		return w(gtx)
	}
}

// H sets a fixed height constraint.
func H(dp unit.Dp) func(gtx layout.Context, w layout.Widget) layout.Dimensions {
	return func(gtx layout.Context, w layout.Widget) layout.Dimensions {
		px := gtx.Dp(dp)
		gtx.Constraints.Min.Y = px
		gtx.Constraints.Max.Y = px
		return w(gtx)
	}
}

// MinW sets a minimum width.
func MinW(dp unit.Dp) func(gtx layout.Context, w layout.Widget) layout.Dimensions {
	return func(gtx layout.Context, w layout.Widget) layout.Dimensions {
		px := gtx.Dp(dp)
		if gtx.Constraints.Min.X < px {
			gtx.Constraints.Min.X = px
		}
		return w(gtx)
	}
}

// MaxW sets a maximum width.
func MaxW(dp unit.Dp) func(gtx layout.Context, w layout.Widget) layout.Dimensions {
	return func(gtx layout.Context, w layout.Widget) layout.Dimensions {
		px := gtx.Dp(dp)
		if gtx.Constraints.Max.X > px {
			gtx.Constraints.Max.X = px
		}
		return w(gtx)
	}
}

// WFull forces full width (like `w-full`).
func WFull(gtx layout.Context, w layout.Widget) layout.Dimensions {
	gtx.Constraints.Min.X = gtx.Constraints.Max.X
	return w(gtx)
}
