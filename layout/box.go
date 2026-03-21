package layout

import (
	"image"
	"image/color"

	"gioui.org/layout"
	"gioui.org/op"
	"gioui.org/op/clip"
	"gioui.org/op/paint"
	"gioui.org/unit"
)

// Box is a simple rectangular container with optional padding, bg, and radius.
// Similar to a <div> with Tailwind classes like `p-4 bg-white rounded-lg`.
type Box struct {
	Padding    layout.Inset
	Background color.NRGBA
	Radius     unit.Dp
	MinWidth   unit.Dp
	MinHeight  unit.Dp
	MaxWidth   unit.Dp
	Border     Border
}

// Border represents a CSS-like border.
type Border struct {
	Color color.NRGBA
	Width unit.Dp
}

// Layout renders the box.
func (b Box) Layout(gtx layout.Context, w layout.Widget) layout.Dimensions {
	// Apply max width constraint
	if b.MaxWidth > 0 {
		maxPx := gtx.Dp(b.MaxWidth)
		if gtx.Constraints.Max.X > maxPx {
			gtx.Constraints.Max.X = maxPx
		}
	}

	// Apply min dimensions
	if b.MinWidth > 0 {
		minPx := gtx.Dp(b.MinWidth)
		if gtx.Constraints.Min.X < minPx {
			gtx.Constraints.Min.X = minPx
		}
	}
	if b.MinHeight > 0 {
		minPx := gtx.Dp(b.MinHeight)
		if gtx.Constraints.Min.Y < minPx {
			gtx.Constraints.Min.Y = minPx
		}
	}

	return b.Padding.Layout(gtx, func(gtx layout.Context) layout.Dimensions {
		return layout.Stack{}.Layout(gtx,
			// Background
			layout.Expanded(func(gtx layout.Context) layout.Dimensions {
				sz := gtx.Constraints.Min
				rr := gtx.Dp(b.Radius)
				rrect := clip.UniformRRect(image.Rectangle{Max: sz}, rr)
				defer rrect.Push(gtx.Ops).Pop()

				// Fill background
				if b.Background.A > 0 {
					paint.ColorOp{Color: b.Background}.Add(gtx.Ops)
					paint.PaintOp{}.Add(gtx.Ops)
				}

				// Draw border
				if b.Border.Width > 0 {
					bw := gtx.Dp(b.Border.Width)
					drawBorder(gtx.Ops, sz, rr, bw, b.Border.Color)
				}

				return layout.Dimensions{Size: sz}
			}),
			// Content
			layout.Stacked(w),
		)
	})
}

// drawBorder draws a rectangular border.
func drawBorder(ops *op.Ops, sz image.Point, radius, width int, col color.NRGBA) {
	// Simple border: draw outer rect, then mask inner rect
	// For simplicity, we use stroke-based approach
	r := image.Rectangle{Max: sz}
	paint.FillShape(ops, col,
		clip.Stroke{
			Path:  clip.UniformRRect(r, radius).Path(ops),
			Width: float32(width),
		}.Op(),
	)
}

// Container is a centered max-width container (like Tailwind `container mx-auto`).
type Container struct {
	MaxWidth unit.Dp
	Padding  unit.Dp
}

// Layout centers the content with max-width.
func (c Container) Layout(gtx layout.Context, w layout.Widget) layout.Dimensions {
	maxW := gtx.Dp(c.MaxWidth)
	if maxW <= 0 {
		maxW = gtx.Dp(1280) // default xl breakpoint
	}
	pad := gtx.Dp(c.Padding)

	return layout.Flex{Alignment: layout.Middle}.Layout(gtx,
		layout.Flexed(1, func(gtx layout.Context) layout.Dimensions {
			if gtx.Constraints.Max.X > maxW {
				gtx.Constraints.Max.X = maxW
			}
			return layout.Inset{
				Left:  unit.Dp(pad),
				Right: unit.Dp(pad),
			}.Layout(gtx, w)
		}),
	)
}
