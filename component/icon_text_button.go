package component

import (
	"image"
	"image/color"

	"gioui.org/font"
	"gioui.org/layout"
	"gioui.org/op/clip"
	"gioui.org/op/paint"
	"gioui.org/unit"
	"gioui.org/widget"

	kitlayout "github.com/hongshengjie/gioui-kit/layout"
	"github.com/hongshengjie/gioui-kit/theme"
)

// IconTextButton is a DaisyUI-style button with a leading icon and a text label.
type IconTextButton struct {
	Icon      *widget.Icon
	Label     string
	Variant   BtnVariant
	Size      BtnSize
	Disabled  bool
	Clickable *widget.Clickable
	th        *theme.Theme
}

// NewIconTextButton creates a new icon+text button.
func NewIconTextButton(th *theme.Theme, click *widget.Clickable, icon *widget.Icon, label string) *IconTextButton {
	return &IconTextButton{
		Icon:      icon,
		Label:     label,
		Variant:   BtnDefault,
		Size:      BtnMd,
		Clickable: click,
		th:        th,
	}
}

// WithVariant sets the button style variant.
func (b *IconTextButton) WithVariant(v BtnVariant) *IconTextButton {
	b.Variant = v
	return b
}

// WithSize sets the button size.
func (b *IconTextButton) WithSize(s BtnSize) *IconTextButton {
	b.Size = s
	return b
}

func (b *IconTextButton) colors() (bg, fg color.NRGBA) {
	th := b.th
	switch b.Variant {
	case BtnPrimary:
		return th.Primary, th.PrimaryContent
	case BtnSecondary:
		return th.Secondary, th.SecondaryContent
	case BtnAccent:
		return th.Accent, th.AccentContent
	case BtnInfo:
		return th.Info, th.InfoContent
	case BtnSuccess:
		return th.Success, th.SuccessContent
	case BtnWarning:
		return th.Warning, th.WarningContent
	case BtnError:
		return th.Error, th.ErrorContent
	case BtnGhost:
		return theme.Transparent, th.BaseContent
	case BtnLink:
		return theme.Transparent, th.Primary
	case BtnOutline:
		return theme.Transparent, th.BaseContent
	default:
		return th.Neutral, th.NeutralContent
	}
}

func (b *IconTextButton) sizeParams() (pad layout.Inset, iconDp unit.Dp) {
	th := b.th
	switch b.Size {
	case BtnXs:
		return layout.Inset{Top: th.Space1, Bottom: th.Space1, Left: th.Space2, Right: th.Space2}, unit.Dp(12)
	case BtnSm:
		return layout.Inset{Top: th.Space1, Bottom: th.Space1, Left: th.Space3, Right: th.Space3}, unit.Dp(14)
	case BtnLg:
		return layout.Inset{Top: th.Space3, Bottom: th.Space3, Left: th.Space6, Right: th.Space6}, unit.Dp(20)
	default: // BtnMd
		return layout.Inset{Top: th.Space2, Bottom: th.Space2, Left: th.Space4, Right: th.Space4}, unit.Dp(16)
	}
}

// Layout renders the icon+text button.
func (b *IconTextButton) Layout(gtx layout.Context) layout.Dimensions {
	th := b.th
	bg, fg := b.colors()
	radius := gtx.Dp(th.RoundedLg)
	pad, iconDp := b.sizeParams()

	if b.Clickable.Hovered() && !b.Disabled {
		bg = theme.Lerp(bg, theme.Black, 0.1)
	}
	if b.Clickable.Pressed() && !b.Disabled {
		bg = theme.Lerp(bg, theme.Black, 0.2)
	}
	if b.Disabled {
		bg = theme.Opacity(bg, 0.5)
		fg = theme.Opacity(fg, 0.5)
	}

	return b.Clickable.Layout(gtx, func(gtx layout.Context) layout.Dimensions {
		return layout.Stack{Alignment: layout.Center}.Layout(gtx,
			layout.Expanded(func(gtx layout.Context) layout.Dimensions {
				sz := gtx.Constraints.Min
				if bg.A > 0 {
					defer clip.UniformRRect(image.Rectangle{Max: sz}, radius).Push(gtx.Ops).Pop()
					paint.ColorOp{Color: bg}.Add(gtx.Ops)
					paint.PaintOp{}.Add(gtx.Ops)
				}
				return layout.Dimensions{Size: sz}
			}),
			layout.Stacked(func(gtx layout.Context) layout.Dimensions {
				return pad.Layout(gtx, func(gtx layout.Context) layout.Dimensions {
					return kitlayout.FlexRow{Gap: th.Space2, Alignment: layout.Middle}.Layout(gtx,
						layout.Rigid(func(gtx layout.Context) layout.Dimensions {
							sz := gtx.Dp(iconDp)
							gtx.Constraints = layout.Exact(image.Pt(sz, sz))
							return b.Icon.Layout(gtx, fg)
						}),
						layout.Rigid(func(gtx layout.Context) layout.Dimensions {
							return drawText(gtx, th, b.Label, fg, th.SmSize, font.Normal)
						}),
					)
				})
			}),
		)
	})
}
