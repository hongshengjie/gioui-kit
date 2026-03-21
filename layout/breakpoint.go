package layout

import "gioui.org/layout"

// Breakpoint represents a Tailwind-style screen-width breakpoint.
//
// Breakpoints mirror Tailwind's default thresholds:
//
//	BreakpointXs  < 640 dp   (phones)
//	BreakpointSm  ≥ 640 dp
//	BreakpointMd  ≥ 768 dp   (tablets)
//	BreakpointLg  ≥ 1024 dp  (laptops)
//	BreakpointXl  ≥ 1280 dp
//	Breakpoint2xl ≥ 1536 dp
type Breakpoint int

const (
	BreakpointXs  Breakpoint = iota // < 640 dp
	BreakpointSm                    // ≥ 640 dp
	BreakpointMd                    // ≥ 768 dp
	BreakpointLg                    // ≥ 1024 dp
	BreakpointXl                    // ≥ 1280 dp
	Breakpoint2xl                   // ≥ 1536 dp
)

// ScreenBreakpoint returns the active Tailwind breakpoint for gtx.
// It is based on gtx.Constraints.Max.X (the available width).
func ScreenBreakpoint(gtx layout.Context) Breakpoint {
	dp := ScreenWidthDp(gtx)
	switch {
	case dp >= 1536:
		return Breakpoint2xl
	case dp >= 1280:
		return BreakpointXl
	case dp >= 1024:
		return BreakpointLg
	case dp >= 768:
		return BreakpointMd
	case dp >= 640:
		return BreakpointSm
	default:
		return BreakpointXs
	}
}

// ScreenWidthDp returns the current available width in device-independent pixels.
func ScreenWidthDp(gtx layout.Context) float32 {
	return float32(gtx.Constraints.Max.X) / gtx.Metric.PxPerDp
}
