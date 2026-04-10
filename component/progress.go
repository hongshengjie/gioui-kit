package component

import (
	"image"
	"image/color"

	"gioui.org/layout"
	"gioui.org/op/clip"
	"gioui.org/op/paint"

	"github.com/ossprovider/gioui-kit/theme"
)

type ProgressVariant int

const (
	ProgressPrimary ProgressVariant = iota
	ProgressSecondary
	ProgressAccent
	ProgressInfo
	ProgressSuccess
	ProgressWarning
	ProgressError
)

// Progress renders a progress bar.
type Progress struct {
	Value   float32 // 0.0 to 1.0
	Variant ProgressVariant
	th      *theme.Theme
}

func NewProgress(th *theme.Theme, value float32) *Progress {
	return &Progress{Value: value, th: th}
}

func (p *Progress) color() color.NRGBA {
	th := p.th
	switch p.Variant {
	case ProgressSecondary:
		return th.Secondary
	case ProgressAccent:
		return th.Accent
	case ProgressInfo:
		return th.Info
	case ProgressSuccess:
		return th.Success
	case ProgressWarning:
		return th.Warning
	case ProgressError:
		return th.Error
	default:
		return th.Primary
	}
}

func (p *Progress) Layout(gtx layout.Context) layout.Dimensions {
	th := p.th
	h := gtx.Dp(8)
	w := gtx.Constraints.Max.X
	radius := h / 2

	// Track
	trackRect := image.Rect(0, 0, w, h)
	defer clip.UniformRRect(trackRect, radius).Push(gtx.Ops).Pop()
	paint.ColorOp{Color: th.Base300}.Add(gtx.Ops)
	paint.PaintOp{}.Add(gtx.Ops)

	// Fill
	fillW := int(float32(w) * p.Value)
	if fillW > 0 {
		fillRect := image.Rect(0, 0, fillW, h)
		defer clip.UniformRRect(fillRect, radius).Push(gtx.Ops).Pop()
		paint.ColorOp{Color: p.color()}.Add(gtx.Ops)
		paint.PaintOp{}.Add(gtx.Ops)
	}

	return layout.Dimensions{Size: image.Pt(w, h)}
}
