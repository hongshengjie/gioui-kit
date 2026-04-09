package scaffold

import (
	"image"
	"image/color"

	"gioui.org/layout"
	"gioui.org/op"
	"gioui.org/op/clip"
	"gioui.org/op/paint"
	"gioui.org/unit"
	"gioui.org/widget"

	"github.com/hongshengjie/gioui-kit/theme"
)

type DrawerSide int

const (
	DrawerLeft DrawerSide = iota
	DrawerRight
)

// Drawer renders a slide-in panel overlay.
type Drawer struct {
	Visible bool
	Side    DrawerSide
	Width   unit.Dp
	OnClose func() // called when the backdrop is tapped
	th      *theme.Theme
	back    widget.Clickable
}

func NewDrawer(th *theme.Theme) *Drawer {
	return &Drawer{
		Width: 300,
		Side:  DrawerLeft,
		th:    th,
	}
}

func (d *Drawer) Toggle() { d.Visible = !d.Visible }
func (d *Drawer) Open()   { d.Visible = true }
func (d *Drawer) Close()  { d.Visible = false }

func (d *Drawer) Layout(gtx layout.Context, content layout.Widget) layout.Dimensions {
	if !d.Visible {
		return layout.Dimensions{}
	}
	th := d.th
	w := gtx.Dp(d.Width)

	return layout.Stack{}.Layout(gtx,
		// Backdrop (tappable to close)
		layout.Expanded(func(gtx layout.Context) layout.Dimensions {
			if d.back.Clicked(gtx) && d.OnClose != nil {
				d.OnClose()
			}
			return d.back.Layout(gtx, func(gtx layout.Context) layout.Dimensions {
				sz := gtx.Constraints.Max
				paint.FillShape(gtx.Ops, color.NRGBA{A: 100},
					clip.Rect{Max: sz}.Op(),
				)
				return layout.Dimensions{Size: sz}
			})
		}),
		// Drawer panel
		layout.Stacked(func(gtx layout.Context) layout.Dimensions {
			var offsetX int
			if d.Side == DrawerRight {
				offsetX = gtx.Constraints.Max.X - w
			}
			defer op.Offset(image.Pt(offsetX, 0)).Push(gtx.Ops).Pop()

			sz := image.Pt(w, gtx.Constraints.Max.Y)
			paint.FillShape(gtx.Ops, th.Base100,
				clip.Rect{Max: sz}.Op(),
			)

			// Draw right border for left drawer
			if d.Side == DrawerLeft {
				borderRect := image.Rect(w-1, 0, w, sz.Y)
				paint.FillShape(gtx.Ops, th.Base300,
					clip.Rect(borderRect).Op(),
				)
			}

			cgtx := gtx
			cgtx.Constraints.Min = sz
			cgtx.Constraints.Max = sz
			return content(cgtx)
		}),
	)
}
