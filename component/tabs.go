package component

import (
	"image"

	"gioui.org/font"
	"gioui.org/layout"
	"gioui.org/op/clip"
	"gioui.org/op/paint"
	"gioui.org/widget"

	"github.com/hongshengjie/gioui-kit/theme"
)

type TabVariant int

const (
	TabBoxed TabVariant = iota
	TabBordered
	TabLifted
)

// Tabs manages a tabbed interface.
type Tabs struct {
	Items    []string
	Selected int
	Variant  TabVariant
	clicks   []widget.Clickable
	children []layout.FlexChild // reused across frames to avoid per-frame alloc
	th       *theme.Theme
}

func NewTabs(th *theme.Theme, items []string) *Tabs {
	return &Tabs{
		Items:  items,
		clicks: make([]widget.Clickable, len(items)),
		th:     th,
	}
}

func (t *Tabs) Layout(gtx layout.Context) layout.Dimensions {
	th := t.th

	// Check clicks
	for i := range t.clicks {
		if t.clicks[i].Clicked(gtx) {
			t.Selected = i
		}
	}

	if cap(t.children) < len(t.Items) {
		t.children = make([]layout.FlexChild, len(t.Items))
	}
	t.children = t.children[:len(t.Items)]
	for i, item := range t.Items {
		i, item := i, item
		t.children[i] = layout.Rigid(func(gtx layout.Context) layout.Dimensions {
			isActive := i == t.Selected
			return t.clicks[i].Layout(gtx, func(gtx layout.Context) layout.Dimensions {
				padding := layout.Inset{
					Top: th.Space2, Bottom: th.Space2,
					Left: th.Space4, Right: th.Space4,
				}
				return padding.Layout(gtx, func(gtx layout.Context) layout.Dimensions {
					return layout.Stack{Alignment: layout.S}.Layout(gtx,
						layout.Stacked(func(gtx layout.Context) layout.Dimensions {
							fg := th.BaseContent
							weight := font.Normal
							if isActive {
								fg = th.Primary
								weight = font.SemiBold
							}
							return drawText(gtx, th, item, fg, th.SmSize, weight)
						}),
						layout.Expanded(func(gtx layout.Context) layout.Dimensions {
							if !isActive {
								return layout.Dimensions{}
							}
							sz := gtx.Constraints.Min
							// Active indicator
							indicatorH := gtx.Dp(2)
							indicatorRect := image.Rect(0, sz.Y-indicatorH, sz.X, sz.Y)
							paint.FillShape(gtx.Ops, th.Primary,
								clip.Rect(indicatorRect).Op(),
							)
							return layout.Dimensions{Size: sz}
						}),
					)
				})
			})
		})
	}

	return layout.Flex{}.Layout(gtx, t.children...)
}
