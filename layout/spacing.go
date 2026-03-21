package layout

import (
	"gioui.org/layout"
	"gioui.org/unit"
)

// P returns a uniform Inset (like Tailwind `p-4`).
func P(dp unit.Dp) layout.Inset {
	return layout.UniformInset(dp)
}

// Px returns horizontal padding (like Tailwind `px-4`).
func Px(dp unit.Dp) layout.Inset {
	return layout.Inset{Left: dp, Right: dp}
}

// Py returns vertical padding (like Tailwind `py-4`).
func Py(dp unit.Dp) layout.Inset {
	return layout.Inset{Top: dp, Bottom: dp}
}

// Pt returns top padding.
func Pt(dp unit.Dp) layout.Inset {
	return layout.Inset{Top: dp}
}

// Pb returns bottom padding.
func Pb(dp unit.Dp) layout.Inset {
	return layout.Inset{Bottom: dp}
}

// Pl returns left padding.
func Pl(dp unit.Dp) layout.Inset {
	return layout.Inset{Left: dp}
}

// Pr returns right padding.
func Pr(dp unit.Dp) layout.Inset {
	return layout.Inset{Right: dp}
}

// Inset4 creates an inset with all 4 sides specified.
func Inset4(top, right, bottom, left unit.Dp) layout.Inset {
	return layout.Inset{Top: top, Right: right, Bottom: bottom, Left: left}
}
