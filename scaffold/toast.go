package scaffold

import (
	"image"
	"time"

	"gioui.org/font"
	"gioui.org/layout"
	"gioui.org/op"
	"gioui.org/op/clip"
	"gioui.org/op/paint"

	"github.com/hongshengjie/gioui-kit/theme"
)

type ToastPosition int

const (
	ToastBottom ToastPosition = iota
	ToastTop
	ToastTopRight
	ToastBottomRight
)

type Toast struct {
	Text     string
	Variant  int // reuse AlertVariant
	Position ToastPosition
	Visible  bool
	Duration time.Duration // auto-dismiss duration; 0 defaults to 3s
	showAt   time.Time
	th       *theme.Theme
}

func NewToast(th *theme.Theme) *Toast {
	return &Toast{Position: ToastBottom, th: th}
}

func (t *Toast) Show(text string) {
	t.Text = text
	t.Visible = true
	t.showAt = time.Now()
}

func (t *Toast) Layout(gtx layout.Context) layout.Dimensions {
	if !t.Visible || t.Text == "" {
		return layout.Dimensions{}
	}
	dur := t.Duration
	if dur == 0 {
		dur = 3 * time.Second
	}
	deadline := t.showAt.Add(dur)
	if gtx.Now.After(deadline) {
		t.Visible = false
		return layout.Dimensions{}
	}
	gtx.Execute(op.InvalidateCmd{At: deadline})
	th := t.th
	radius := gtx.Dp(th.RoundedLg)

	var alignment layout.Direction
	switch t.Position {
	case ToastTop:
		alignment = layout.N
	case ToastTopRight:
		alignment = layout.NE
	case ToastBottomRight:
		alignment = layout.SE
	default:
		alignment = layout.S
	}

	return alignment.Layout(gtx, func(gtx layout.Context) layout.Dimensions {
		return layout.Inset{
			Top: th.Space4, Bottom: th.Space4,
			Left: th.Space4, Right: th.Space4,
		}.Layout(gtx, func(gtx layout.Context) layout.Dimensions {
			return layout.Stack{Alignment: layout.Center}.Layout(gtx,
				layout.Expanded(func(gtx layout.Context) layout.Dimensions {
					sz := gtx.Constraints.Min
					defer clip.UniformRRect(image.Rectangle{Max: sz}, radius).Push(gtx.Ops).Pop()
					paint.ColorOp{Color: th.Neutral}.Add(gtx.Ops)
					paint.PaintOp{}.Add(gtx.Ops)
					return layout.Dimensions{Size: sz}
				}),
				layout.Stacked(func(gtx layout.Context) layout.Dimensions {
					return layout.Inset{
						Top: th.Space3, Bottom: th.Space3,
						Left: th.Space4, Right: th.Space4,
					}.Layout(gtx, func(gtx layout.Context) layout.Dimensions {
						return drawLabel(gtx, th, t.Text, th.NeutralContent, th.SmSize, font.Medium)
					})
				}),
			)
		})
	})
}
