package component

import (
	"image"

	"gioui.org/font"
	"gioui.org/layout"
	"gioui.org/op/clip"
	"gioui.org/op/paint"
	"gioui.org/widget"

	"github.com/hongshengjie/gioui-kit/theme"
)

type Chip struct {
	Text     string
	Closable bool
	close    widget.Clickable
	th       *theme.Theme
}

func NewChip(th *theme.Theme, text string) *Chip {
	return &Chip{Text: text, th: th}
}

func (c *Chip) Layout(gtx layout.Context) layout.Dimensions {
	th := c.th

	return layout.Stack{Alignment: layout.Center}.Layout(gtx,
		layout.Expanded(func(gtx layout.Context) layout.Dimensions {
			sz := gtx.Constraints.Min
			radius := min(sz.X, sz.Y) / 2 // cap to avoid degenerate Bezier curves
			defer clip.UniformRRect(image.Rectangle{Max: sz}, radius).Push(gtx.Ops).Pop()
			paint.ColorOp{Color: th.Base200}.Add(gtx.Ops)
			paint.PaintOp{}.Add(gtx.Ops)
			return layout.Dimensions{Size: sz}
		}),
		layout.Stacked(func(gtx layout.Context) layout.Dimensions {
			inset := layout.Inset{Top: 4, Bottom: 4, Left: 12, Right: 12}
			return inset.Layout(gtx, func(gtx layout.Context) layout.Dimensions {
				return drawText(gtx, th, c.Text, th.BaseContent, th.SmSize, font.Medium)
			})
		}),
	)
}
