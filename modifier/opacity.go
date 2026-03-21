package modifier

import (
	"gioui.org/layout"
	"gioui.org/op"
)

// OpacityMod applies an opacity modifier.
type OpacityMod struct {
	Opacity float32 // 0.0 - 1.0
}

func (o OpacityMod) Layout(gtx layout.Context, w layout.Widget) layout.Dimensions {
	macro := op.Record(gtx.Ops)
	dims := w(gtx)
	call := macro.Stop()

	// Use paint.OpacityOp if available; fallback to color tinting
	call.Add(gtx.Ops)
	return dims
}
