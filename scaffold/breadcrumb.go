package scaffold

import (
	"gioui.org/font"
	"gioui.org/layout"

	"github.com/hongshengjie/gioui-kit/theme"
)

type Breadcrumb struct {
	Items    []string
	children []layout.FlexChild // reused across frames
	th       *theme.Theme
}

func NewBreadcrumb(th *theme.Theme, items ...string) *Breadcrumb {
	return &Breadcrumb{Items: items, th: th}
}

func (b *Breadcrumb) Layout(gtx layout.Context) layout.Dimensions {
	th := b.th
	b.children = b.children[:0]

	for i, item := range b.Items {
		i, item := i, item
		if i > 0 {
			b.children = append(b.children, layout.Rigid(func(gtx layout.Context) layout.Dimensions {
				return layout.Inset{Left: th.Space2, Right: th.Space2}.Layout(gtx, func(gtx layout.Context) layout.Dimensions {
					return drawLabel(gtx, th, "/", theme.Opacity(th.BaseContent, 0.4), th.SmSize, font.Normal)
				})
			}))
		}
		b.children = append(b.children, layout.Rigid(func(gtx layout.Context) layout.Dimensions {
			col := th.BaseContent
			weight := font.Normal
			if i == len(b.Items)-1 {
				weight = font.Medium
			} else {
				col = theme.Opacity(col, 0.6)
			}
			return drawLabel(gtx, th, item, col, th.SmSize, weight)
		}))
	}

	return layout.Flex{Alignment: layout.Middle}.Layout(gtx, b.children...)
}
