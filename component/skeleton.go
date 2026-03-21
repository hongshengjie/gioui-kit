package component

import (
	"image"

	"gioui.org/layout"
	"gioui.org/op/clip"
	"gioui.org/op/paint"
	"gioui.org/unit"

	"github.com/hongshengjie/gioui-kit/theme"
)

// Skeleton renders a loading skeleton placeholder.
type Skeleton struct {
	Width  unit.Dp
	Height unit.Dp
	Radius unit.Dp
	th     *theme.Theme
}

func NewSkeleton(th *theme.Theme) *Skeleton {
	return &Skeleton{Width: 200, Height: 20, Radius: th.RoundedMd, th: th}
}

func (s *Skeleton) Layout(gtx layout.Context) layout.Dimensions {
	w := gtx.Dp(s.Width)
	h := gtx.Dp(s.Height)
	rr := gtx.Dp(s.Radius)

	rect := image.Rect(0, 0, w, h)
	defer clip.UniformRRect(rect, rr).Push(gtx.Ops).Pop()
	paint.ColorOp{Color: s.th.Base300}.Add(gtx.Ops)
	paint.PaintOp{}.Add(gtx.Ops)

	return layout.Dimensions{Size: image.Pt(w, h)}
}
