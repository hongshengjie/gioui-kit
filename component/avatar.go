package component

import (
	"image"

	"gioui.org/font"
	"gioui.org/layout"
	"gioui.org/op"
	"gioui.org/op/clip"
	"gioui.org/op/paint"
	"gioui.org/unit"

	"github.com/hongshengjie/gioui-kit/theme"
)

type AvatarSize int

const (
	AvatarMd AvatarSize = iota
	AvatarXs
	AvatarSm
	AvatarLg
)

// Avatar renders a circular avatar placeholder.
type Avatar struct {
	Initials string
	Size     AvatarSize
	Online   bool
	th       *theme.Theme
}

func NewAvatar(th *theme.Theme, initials string) *Avatar {
	return &Avatar{Initials: initials, Size: AvatarMd, th: th}
}

func (a *Avatar) sizeDp() unit.Dp {
	switch a.Size {
	case AvatarXs:
		return 24
	case AvatarSm:
		return 32
	case AvatarLg:
		return 64
	default:
		return 48
	}
}

func (a *Avatar) Layout(gtx layout.Context) layout.Dimensions {
	th := a.th
	sz := gtx.Dp(a.sizeDp())
	rect := image.Rect(0, 0, sz, sz)

	// Circle background
	defer clip.UniformRRect(rect, sz/2).Push(gtx.Ops).Pop()
	paint.ColorOp{Color: th.Primary}.Add(gtx.Ops)
	paint.PaintOp{}.Add(gtx.Ops)

	// Initials
	textSize := unit.Sp(float32(sz) * 0.4)
	macro := op.Record(gtx.Ops)
	dims := drawText(gtx, th, a.Initials, th.PrimaryContent, textSize, font.Bold)
	call := macro.Stop()

	offX := (sz - dims.Size.X) / 2
	offY := (sz - dims.Size.Y) / 2
	defer op.Offset(image.Pt(offX, offY)).Push(gtx.Ops).Pop()
	call.Add(gtx.Ops)

	return layout.Dimensions{Size: image.Pt(sz, sz)}
}
