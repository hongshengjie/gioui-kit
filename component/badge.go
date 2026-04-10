package component

import (
	"image"
	"image/color"

	"gioui.org/font"
	"gioui.org/layout"
	"gioui.org/op/clip"
	"gioui.org/op/paint"

	"github.com/ossprovider/gioui-kit/theme"
)

type BadgeVariant int

const (
	BadgeDefault BadgeVariant = iota
	BadgePrimary
	BadgeSecondary
	BadgeAccent
	BadgeInfo
	BadgeSuccess
	BadgeWarning
	BadgeError
	BadgeOutline
	BadgeGhost
)

// Badge is a DaisyUI-style badge/tag component.
type Badge struct {
	Text    string
	Variant BadgeVariant
	th      *theme.Theme
}

func NewBadge(th *theme.Theme, text string) *Badge {
	return &Badge{Text: text, th: th}
}

func (b *Badge) WithVariant(v BadgeVariant) *Badge {
	b.Variant = v
	return b
}

func (b *Badge) colors() (bg, fg color.NRGBA) {
	th := b.th
	switch b.Variant {
	case BadgePrimary:
		return th.Primary, th.PrimaryContent
	case BadgeSecondary:
		return th.Secondary, th.SecondaryContent
	case BadgeAccent:
		return th.Accent, th.AccentContent
	case BadgeInfo:
		return th.Info, th.InfoContent
	case BadgeSuccess:
		return th.Success, th.SuccessContent
	case BadgeWarning:
		return th.Warning, th.WarningContent
	case BadgeError:
		return th.Error, th.ErrorContent
	case BadgeGhost:
		return th.Base200, th.BaseContent
	default:
		return th.Neutral, th.NeutralContent
	}
}

func (b *Badge) Layout(gtx layout.Context) layout.Dimensions {
	bg, fg := b.colors()

	// Padding is inside Stacked so Expanded.Min reflects the full pill size.
	return layout.Stack{Alignment: layout.Center}.Layout(gtx,
		layout.Expanded(func(gtx layout.Context) layout.Dimensions {
			sz := gtx.Constraints.Min
			// Cap radius to half the shortest side to avoid degenerate Bezier curves.
			radius := min(sz.X, sz.Y) / 2
			defer clip.UniformRRect(image.Rectangle{Max: sz}, radius).Push(gtx.Ops).Pop()
			paint.ColorOp{Color: bg}.Add(gtx.Ops)
			paint.PaintOp{}.Add(gtx.Ops)
			return layout.Dimensions{Size: sz}
		}),
		layout.Stacked(func(gtx layout.Context) layout.Dimensions {
			return layout.Inset{Top: 2, Bottom: 2, Left: 8, Right: 8}.Layout(gtx, func(gtx layout.Context) layout.Dimensions {
				return drawText(gtx, b.th, b.Text, fg, b.th.XsSize, font.SemiBold)
			})
		}),
	)
}
