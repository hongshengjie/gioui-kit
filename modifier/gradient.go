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

// GradientDir specifies gradient direction.
type GradientDir int

const (
	GradientToRight GradientDir = iota
	GradientToBottom
	GradientToBottomRight
)

// LinearGradient applies a linear gradient background.
type LinearGradient struct {
	From   color.NRGBA
	To     color.NRGBA
	Dir    GradientDir
	Radius unit.Dp
}

func (g LinearGradient) Layout(gtx layout.Context, w layout.Widget) layout.Dimensions {
	macro := op.Record(gtx.Ops)
	dims := w(gtx)
	call := macro.Stop()

	rr := gtx.Dp(g.Radius)

	// Draw gradient via multiple horizontal or vertical strips
	steps := 32
	for i := 0; i < steps; i++ {
		t := float32(i) / float32(steps)
		t2 := float32(i+1) / float32(steps)
		col := lerpColor(g.From, g.To, t)

		var stripRect image.Rectangle
		switch g.Dir {
		case GradientToRight:
			x1 := int(t * float32(dims.Size.X))
			x2 := int(t2 * float32(dims.Size.X))
			stripRect = image.Rect(x1, 0, x2, dims.Size.Y)
		case GradientToBottom:
			y1 := int(t * float32(dims.Size.Y))
			y2 := int(t2 * float32(dims.Size.Y))
			stripRect = image.Rect(0, y1, dims.Size.X, y2)
		case GradientToBottomRight:
			x1 := int(t * float32(dims.Size.X))
			x2 := int(t2 * float32(dims.Size.X))
			y1 := int(t * float32(dims.Size.Y))
			y2 := int(t2 * float32(dims.Size.Y))
			stripRect = image.Rect(x1, y1, x2, y2)
		}

		if i == 0 && rr > 0 {
			defer clip.UniformRRect(image.Rectangle{Max: dims.Size}, rr).Push(gtx.Ops).Pop()
		}
		paint.FillShape(gtx.Ops, col, clip.Rect(stripRect).Op())
	}

	call.Add(gtx.Ops)
	return dims
}

func lerpColor(a, b color.NRGBA, t float32) color.NRGBA {
	return color.NRGBA{
		R: uint8(float32(a.R)*(1-t) + float32(b.R)*t),
		G: uint8(float32(a.G)*(1-t) + float32(b.G)*t),
		B: uint8(float32(a.B)*(1-t) + float32(b.B)*t),
		A: uint8(float32(a.A)*(1-t) + float32(b.A)*t),
	}
}
