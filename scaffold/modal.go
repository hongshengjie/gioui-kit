package scaffold

import (
	"image"
	"image/color"

	"gioui.org/layout"
	"gioui.org/op/clip"
	"gioui.org/op/paint"
	"gioui.org/unit"
	"gioui.org/widget"

	"github.com/ossprovider/gioui-kit/theme"
)

// Modal renders a centered dialog overlay.
type Modal struct {
	Visible  bool
	MaxWidth unit.Dp
	backdrop widget.Clickable
	th       *theme.Theme
}

func NewModal(th *theme.Theme) *Modal {
	return &Modal{MaxWidth: 500, th: th}
}

func (m *Modal) Show() { m.Visible = true }
func (m *Modal) Hide() { m.Visible = false }

// Layout renders the modal overlay with content.
func (m *Modal) Layout(gtx layout.Context, content layout.Widget) layout.Dimensions {
	if !m.Visible {
		return layout.Dimensions{}
	}
	th := m.th

	return layout.Stack{Alignment: layout.Center}.Layout(gtx,
		// Backdrop — click to close
		layout.Expanded(func(gtx layout.Context) layout.Dimensions {
			sz := gtx.Constraints.Max
			return m.backdrop.Layout(gtx, func(gtx layout.Context) layout.Dimensions {
				if m.backdrop.Clicked(gtx) {
					m.Visible = false
				}
				paint.FillShape(gtx.Ops, color.NRGBA{A: 128},
					clip.Rect{Max: sz}.Op(),
				)
				return layout.Dimensions{Size: sz}
			})
		}),
		// Dialog
		layout.Stacked(func(gtx layout.Context) layout.Dimensions {
			maxW := gtx.Dp(m.MaxWidth)
			if gtx.Constraints.Max.X > maxW {
				gtx.Constraints.Max.X = maxW
			}
			gtx.Constraints.Min.X = gtx.Constraints.Max.X

			radius := gtx.Dp(th.Rounded2xl)
			padding := layout.UniformInset(th.Space6)

			return layout.Stack{}.Layout(gtx,
				layout.Expanded(func(gtx layout.Context) layout.Dimensions {
					sz := gtx.Constraints.Min
					defer clip.UniformRRect(image.Rectangle{Max: sz}, radius).Push(gtx.Ops).Pop()
					paint.ColorOp{Color: th.Base100}.Add(gtx.Ops)
					paint.PaintOp{}.Add(gtx.Ops)
					return layout.Dimensions{Size: sz}
				}),
				layout.Stacked(func(gtx layout.Context) layout.Dimensions {
					return padding.Layout(gtx, content)
				}),
			)
		}),
	)
}
