package layout

import (
	"gioui.org/layout"
)

// ScrollY is a convenience wrapper around layout.List for vertical scrolling.
type ScrollY struct {
	List layout.List
}

// NewScrollY creates a new vertical scrollable list.
func NewScrollY() *ScrollY {
	return &ScrollY{
		List: layout.List{
			Axis: layout.Vertical,
		},
	}
}

// Layout renders a scrollable list of widgets.
func (s *ScrollY) Layout(gtx layout.Context, count int, w layout.ListElement) layout.Dimensions {
	return s.List.Layout(gtx, count, w)
}

// ScrollX is a horizontal scrollable list.
type ScrollX struct {
	List layout.List
}

func NewScrollX() *ScrollX {
	return &ScrollX{
		List: layout.List{
			Axis: layout.Horizontal,
		},
	}
}

func (s *ScrollX) Layout(gtx layout.Context, count int, w layout.ListElement) layout.Dimensions {
	return s.List.Layout(gtx, count, w)
}
