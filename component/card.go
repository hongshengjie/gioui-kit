package component

import (
	"image"

	"gioui.org/font"
	"gioui.org/layout"
	"gioui.org/op/clip"
	"gioui.org/op/paint"

	"github.com/hongshengjie/gioui-kit/theme"
)

// Card is a DaisyUI-style card container.
type Card struct {
	Bordered bool
	Compact  bool
	th       *theme.Theme
}

func NewCard(th *theme.Theme) *Card {
	return &Card{th: th}
}

func (c *Card) WithBorder() *Card {
	c.Bordered = true
	return c
}

func (c *Card) WithCompact() *Card {
	c.Compact = true
	return c
}

func (c *Card) Layout(gtx layout.Context, body layout.Widget) layout.Dimensions {
	th := c.th
	radius := gtx.Dp(th.RoundedXl)
	padding := layout.UniformInset(th.Space6)
	if c.Compact {
		padding = layout.UniformInset(th.Space4)
	}

	return layout.Stack{}.Layout(gtx,
		layout.Expanded(func(gtx layout.Context) layout.Dimensions {
			sz := gtx.Constraints.Min
			outerStack := clip.UniformRRect(image.Rectangle{Max: sz}, radius).Push(gtx.Ops)
			defer outerStack.Pop()

			if c.Bordered {
				// Paint-over border: fill outer with border color, then overpaint center with background.
				// Avoids clip.Stroke which can loop infinitely on zero-height rectangles.
				paint.ColorOp{Color: th.Base300}.Add(gtx.Ops)
				paint.PaintOp{}.Add(gtx.Ops)
				bw := gtx.Dp(1)
				inner := image.Rect(bw, bw, sz.X-bw, sz.Y-bw)
				if inner.Dx() > 0 && inner.Dy() > 0 {
					innerRadius := radius - bw
					if innerRadius < 0 {
						innerRadius = 0
					}
					innerStack := clip.UniformRRect(inner, innerRadius).Push(gtx.Ops)
					paint.ColorOp{Color: th.Base100}.Add(gtx.Ops)
					paint.PaintOp{}.Add(gtx.Ops)
					innerStack.Pop()
				}
			} else {
				paint.ColorOp{Color: th.Base100}.Add(gtx.Ops)
				paint.PaintOp{}.Add(gtx.Ops)
			}

			return layout.Dimensions{Size: sz}
		}),
		layout.Stacked(func(gtx layout.Context) layout.Dimensions {
			return padding.Layout(gtx, body)
		}),
	)
}

// CardWithHeader renders a card with a separate title section.
func (c *Card) CardWithHeader(gtx layout.Context, title string, body layout.Widget) layout.Dimensions {
	th := c.th
	return c.Layout(gtx, func(gtx layout.Context) layout.Dimensions {
		return layout.Flex{Axis: layout.Vertical}.Layout(gtx,
			layout.Rigid(func(gtx layout.Context) layout.Dimensions {
				return layout.Inset{Bottom: th.Space3}.Layout(gtx, func(gtx layout.Context) layout.Dimensions {
					return drawText(gtx, th, title, th.BaseContent, th.H3Size, font.Bold)
				})
			}),
			layout.Rigid(body),
		)
	})
}
