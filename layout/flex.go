// Package layout provides TailwindCSS-inspired layout primitives for Gio UI.
//
// Naming follows Tailwind conventions:
//   - Flex, FlexRow, FlexCol (flex layout)
//   - Grid (grid layout)
//   - Stack (absolute positioning / z-index)
//   - Container, Box (wrapper utilities)
//   - Gap, Padding, Margin modifiers
//
// Alignment uses Tailwind names:
//   - ItemsCenter, ItemsStart, ItemsEnd, ItemsStretch
//   - JustifyCenter, JustifyBetween, JustifyAround, JustifyEnd
package layout

import (
	"image"

	"gioui.org/layout"
	"gioui.org/unit"
)

// FlexRow lays out children horizontally (like Tailwind `flex flex-row`).
type FlexRow struct {
	Gap       unit.Dp
	Alignment layout.Alignment // cross-axis: ItemsStart, ItemsCenter, ItemsEnd
	Spacing   layout.Spacing   // main-axis: JustifyStart, JustifyCenter, etc.
	Wrap      bool
}

// Layout renders children in a horizontal flex row.
func (f FlexRow) Layout(gtx layout.Context, children ...layout.FlexChild) layout.Dimensions {
	flex := layout.Flex{
		Axis:      layout.Horizontal,
		Alignment: f.Alignment,
		Spacing:   f.Spacing,
	}
	if f.Gap > 0 {
		children = insertGaps(gtx, layout.Horizontal, f.Gap, children)
	}
	return flex.Layout(gtx, children...)
}

// FlexCol lays out children vertically (like Tailwind `flex flex-col`).
type FlexCol struct {
	Gap       unit.Dp
	Alignment layout.Alignment
	Spacing   layout.Spacing
}

// Layout renders children in a vertical flex column.
func (f FlexCol) Layout(gtx layout.Context, children ...layout.FlexChild) layout.Dimensions {
	flex := layout.Flex{
		Axis:      layout.Vertical,
		Alignment: f.Alignment,
		Spacing:   f.Spacing,
	}
	if f.Gap > 0 {
		children = insertGaps(gtx, layout.Vertical, f.Gap, children)
	}
	return flex.Layout(gtx, children...)
}

// insertGaps inserts spacer FlexChildren between items to simulate gap.
func insertGaps(gtx layout.Context, axis layout.Axis, gap unit.Dp, children []layout.FlexChild) []layout.FlexChild {
	if len(children) <= 1 {
		return children
	}
	result := make([]layout.FlexChild, 0, len(children)*2-1)
	gapPx := gtx.Dp(gap)
	spacer := layout.Rigid(func(gtx layout.Context) layout.Dimensions {
		if axis == layout.Horizontal {
			return layout.Dimensions{Size: image.Pt(gapPx, 0)}
		}
		return layout.Dimensions{Size: image.Pt(0, gapPx)}
	})
	for i, child := range children {
		if i > 0 {
			result = append(result, spacer)
		}
		result = append(result, child)
	}
	return result
}

// Rigid wraps layout.Rigid for convenience.
func Rigid(w layout.Widget) layout.FlexChild {
	return layout.Rigid(w)
}

// Flexed wraps layout.Flexed (like flex-grow).
func Flexed(weight float32, w layout.Widget) layout.FlexChild {
	return layout.Flexed(weight, w)
}

// Grow is shorthand for Flexed(1, w) (like Tailwind `flex-1`).
func Grow(w layout.Widget) layout.FlexChild {
	return layout.Flexed(1, w)
}

// ---------- Alignment constants (Tailwind naming) ----------

const (
	// Cross-axis alignment (items-*)
	ItemsStart    = layout.Start
	ItemsCenter   = layout.Middle
	ItemsEnd      = layout.End
	ItemsBaseline = layout.Baseline

	// Main-axis spacing (justify-*)
	JustifyStart   = layout.SpaceStart
	JustifyEnd     = layout.SpaceEnd
	JustifyCenter  = layout.SpaceSides
	JustifyBetween = layout.SpaceEvenly
	JustifyAround  = layout.SpaceAround
)
