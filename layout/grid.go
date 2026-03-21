package layout

import (
	"image"

	"gioui.org/layout"
	"gioui.org/op"
	"gioui.org/unit"
)

// Grid provides a responsive grid layout (like Tailwind `grid grid-cols-N gap-N`).
//
// Responsive column counts mirror Tailwind's breakpoint prefixes:
//
//	Grid{Cols: 1, MdCols: 2, LgCols: 3, Gap: 16}
//	// → 1 col on mobile, 2 on tablet, 3 on desktop
//
// Zero values for breakpoint-specific fields mean "inherit from the next
// smaller breakpoint". Cols is always the fallback.
type Grid struct {
	Cols   int     // default / mobile-first column count
	SmCols int     // ≥ 640 dp; 0 = use Cols
	MdCols int     // ≥ 768 dp; 0 = use SmCols (or Cols)
	LgCols int     // ≥ 1024 dp; 0 = use MdCols (or …)
	XlCols int     // ≥ 1280 dp; 0 = use LgCols (or …)
	Gap    unit.Dp
}

// activeCols returns the column count for the current screen width.
func (g Grid) activeCols(gtx layout.Context) int {
	bp := ScreenBreakpoint(gtx)

	// Walk from the active breakpoint downward; take the first non-zero value.
	candidates := []int{g.Cols, g.SmCols, g.MdCols, g.LgCols, g.XlCols}
	// indices align with BreakpointXs=0 … BreakpointXl=4; cap at XlCols
	idx := int(bp)
	if idx >= len(candidates) {
		idx = len(candidates) - 1
	}
	for i := idx; i >= 0; i-- {
		if candidates[i] > 0 {
			return candidates[i]
		}
	}
	return 1
}

// Layout renders widgets in a grid.
func (g Grid) Layout(gtx layout.Context, widgets ...layout.Widget) layout.Dimensions {
	g.Cols = g.activeCols(gtx)
	if g.Cols <= 0 {
		g.Cols = 1
	}
	gapPx := gtx.Dp(g.Gap)
	totalGap := gapPx * (g.Cols - 1)
	colWidth := (gtx.Constraints.Max.X - totalGap) / g.Cols

	rows := (len(widgets) + g.Cols - 1) / g.Cols
	var totalHeight int

	for row := 0; row < rows; row++ {
		if row > 0 {
			totalHeight += gapPx
		}
		var rowHeight int
		for col := 0; col < g.Cols; col++ {
			idx := row*g.Cols + col
			if idx >= len(widgets) {
				break
			}
			// Position and size each cell
			offX := col * (colWidth + gapPx)
			macro := op.Record(gtx.Ops)
			cgtx := gtx
			cgtx.Constraints.Min.X = colWidth
			cgtx.Constraints.Max.X = colWidth
			dims := widgets[idx](cgtx)
			call := macro.Stop()

			offset := op.Offset(image.Pt(offX, totalHeight)).Push(gtx.Ops)
			call.Add(gtx.Ops)
			offset.Pop()

			if dims.Size.Y > rowHeight {
				rowHeight = dims.Size.Y
			}
		}
		totalHeight += rowHeight
	}

	return layout.Dimensions{
		Size: image.Pt(gtx.Constraints.Max.X, totalHeight),
	}
}
