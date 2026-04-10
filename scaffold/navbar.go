package scaffold

import (
	"image"
	"image/color"

	"gioui.org/layout"
	"gioui.org/op/clip"
	"gioui.org/op/paint"
	"gioui.org/unit"

	"github.com/ossprovider/gioui-kit/theme"
)

// Navbar renders a top navigation bar (like DaisyUI navbar).
type Navbar struct {
	Height     unit.Dp
	Background color.NRGBA
	Bordered   bool
	th         *theme.Theme
}

func NewNavbar(th *theme.Theme) *Navbar {
	return &Navbar{
		Height:     56,
		Background: th.Base100,
		Bordered:   true,
		th:         th,
	}
}

func (n *Navbar) Layout(gtx layout.Context, start, center, end layout.Widget) layout.Dimensions {
	th := n.th
	h := gtx.Dp(n.Height)

	return layout.Stack{}.Layout(gtx,
		layout.Expanded(func(gtx layout.Context) layout.Dimensions {
			sz := image.Pt(gtx.Constraints.Max.X, h)

			// Background
			paint.FillShape(gtx.Ops, n.Background,
				clip.Rect{Max: sz}.Op(),
			)

			// Bottom border
			if n.Bordered {
				borderRect := image.Rect(0, h-1, sz.X, h)
				paint.FillShape(gtx.Ops, th.Base300,
					clip.Rect(borderRect).Op(),
				)
			}
			return layout.Dimensions{Size: sz}
		}),
		layout.Stacked(func(gtx layout.Context) layout.Dimensions {
			gtx.Constraints.Min.Y = h
			gtx.Constraints.Max.Y = h
			return layout.Inset{Left: th.Space4, Right: th.Space4}.Layout(gtx, func(gtx layout.Context) layout.Dimensions {
				return layout.Flex{
					Alignment: layout.Middle,
				}.Layout(gtx,
					// Start section
					layout.Rigid(func(gtx layout.Context) layout.Dimensions {
						if start == nil {
							return layout.Dimensions{}
						}
						return start(gtx)
					}),
					// Center spacer + center content
					layout.Flexed(1, func(gtx layout.Context) layout.Dimensions {
						if center == nil {
							return layout.Dimensions{Size: image.Pt(gtx.Constraints.Max.X, 0)}
						}
						return layout.Center.Layout(gtx, center)
					}),
					// End section
					layout.Rigid(func(gtx layout.Context) layout.Dimensions {
						if end == nil {
							return layout.Dimensions{}
						}
						return end(gtx)
					}),
				)
			})
		}),
	)
}
