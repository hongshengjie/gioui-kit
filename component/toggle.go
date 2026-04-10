package component

import (
	"image"

	"gioui.org/font"
	"gioui.org/io/pointer"
	"gioui.org/layout"
	"gioui.org/op/clip"
	"gioui.org/op/paint"
	"gioui.org/widget"

	"github.com/ossprovider/gioui-kit/theme"
)

// Toggle is a DaisyUI-style toggle switch.
type Toggle struct {
	Bool    *widget.Bool
	Label   string
	Variant BtnVariant
	th      *theme.Theme
}

func NewToggle(th *theme.Theme, b *widget.Bool, label string) *Toggle {
	return &Toggle{
		Bool:    b,
		Label:   label,
		Variant: BtnPrimary,
		th:      th,
	}
}

func (t *Toggle) Layout(gtx layout.Context) layout.Dimensions {
	th := t.th

	return layout.Flex{
		Alignment: layout.Middle,
	}.Layout(gtx,
		// Toggle track
		layout.Rigid(func(gtx layout.Context) layout.Dimensions {
			trackW := gtx.Dp(44)
			trackH := gtx.Dp(24)
			thumbSize := gtx.Dp(20)
			radius := trackH / 2

			return t.Bool.Layout(gtx, func(gtx layout.Context) layout.Dimensions {
				// Track
				trackColor := th.Base300
				if t.Bool.Value {
					trackColor = th.Primary
				}
				trackRect := image.Rect(0, 0, trackW, trackH)
				defer clip.UniformRRect(trackRect, radius).Push(gtx.Ops).Pop()
				paint.ColorOp{Color: trackColor}.Add(gtx.Ops)
				paint.PaintOp{}.Add(gtx.Ops)

				// Thumb
				thumbX := gtx.Dp(2)
				if t.Bool.Value {
					thumbX = trackW - thumbSize - gtx.Dp(2)
				}
				thumbY := (trackH - thumbSize) / 2
				thumbRect := image.Rect(thumbX, thumbY, thumbX+thumbSize, thumbY+thumbSize)
				defer clip.UniformRRect(thumbRect, thumbSize/2).Push(gtx.Ops).Pop()
				paint.ColorOp{Color: theme.White}.Add(gtx.Ops)
				paint.PaintOp{}.Add(gtx.Ops)

				pointer.CursorPointer.Add(gtx.Ops)
				return layout.Dimensions{Size: image.Pt(trackW, trackH)}
			})
		}),
		// Label
		layout.Rigid(func(gtx layout.Context) layout.Dimensions {
			if t.Label == "" {
				return layout.Dimensions{}
			}
			return layout.Inset{Left: th.Space2}.Layout(gtx, func(gtx layout.Context) layout.Dimensions {
				return drawText(gtx, th, t.Label, th.BaseContent, th.FontSize, font.Normal)
			})
		}),
	)
}
