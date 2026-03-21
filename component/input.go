package component

import (
	"image"
	"image/color"

	"gioui.org/font"
	"gioui.org/layout"
	"gioui.org/op"
	"gioui.org/op/clip"
	"gioui.org/op/paint"
	"gioui.org/text"
	"gioui.org/unit"
	"gioui.org/widget"

	"github.com/hongshengjie/gioui-kit/theme"
)

type InputVariant int

const (
	InputDefault InputVariant = iota
	InputBordered
	InputGhost
	InputPrimary
	InputSecondary
	InputAccent
	InputInfo
	InputSuccess
	InputWarning
	InputError
)

type InputSize int

const (
	InputMd InputSize = iota
	InputXs
	InputSm
	InputLg
)

// Input is a DaisyUI-style text input.
type Input struct {
	Editor      *widget.Editor
	Placeholder string
	Label       string
	Variant     InputVariant
	Size        InputSize
	th          *theme.Theme
}

func NewInput(th *theme.Theme, editor *widget.Editor, placeholder string) *Input {
	return &Input{
		Editor:      editor,
		Placeholder: placeholder,
		Variant:     InputBordered,
		Size:        InputMd,
		th:          th,
	}
}

func (inp *Input) WithLabel(label string) *Input {
	inp.Label = label
	return inp
}

func (inp *Input) WithVariant(v InputVariant) *Input {
	inp.Variant = v
	return inp
}

func (inp *Input) borderColor() color.NRGBA {
	th := inp.th
	switch inp.Variant {
	case InputPrimary:
		return th.Primary
	case InputSecondary:
		return th.Secondary
	case InputAccent:
		return th.Accent
	case InputInfo:
		return th.Info
	case InputSuccess:
		return th.Success
	case InputWarning:
		return th.Warning
	case InputError:
		return th.Error
	case InputGhost:
		return theme.Transparent
	default:
		return th.Base300
	}
}

func (inp *Input) height() unit.Dp {
	switch inp.Size {
	case InputXs:
		return 24
	case InputSm:
		return 32
	case InputLg:
		return 48
	default:
		return 40
	}
}

func (inp *Input) Layout(gtx layout.Context) layout.Dimensions {
	th := inp.th

	return layout.Flex{Axis: layout.Vertical}.Layout(gtx,
		// Label
		layout.Rigid(func(gtx layout.Context) layout.Dimensions {
			if inp.Label == "" {
				return layout.Dimensions{}
			}
			return layout.Inset{Bottom: th.Space1}.Layout(gtx, func(gtx layout.Context) layout.Dimensions {
				return drawText(gtx, th, inp.Label, th.BaseContent, th.SmSize, font.Medium)
			})
		}),
		// Input field
		layout.Rigid(func(gtx layout.Context) layout.Dimensions {
			h := gtx.Dp(inp.height())
			radius := gtx.Dp(th.RoundedLg)
			borderCol := inp.borderColor()

			return layout.Stack{}.Layout(gtx,
				layout.Expanded(func(gtx layout.Context) layout.Dimensions {
					sz := image.Pt(gtx.Constraints.Max.X, h)
					rrect := clip.UniformRRect(image.Rectangle{Max: sz}, radius)
					defer rrect.Push(gtx.Ops).Pop()

					// Background
					paint.ColorOp{Color: th.Base100}.Add(gtx.Ops)
					paint.PaintOp{}.Add(gtx.Ops)

					// Border
					if inp.Variant != InputGhost && sz.X > 0 && sz.Y > 0 {
						paint.FillShape(gtx.Ops, borderCol,
							clip.Stroke{
								Path:  clip.UniformRRect(image.Rectangle{Max: sz}, radius).Path(gtx.Ops),
								Width: float32(gtx.Dp(1)),
							}.Op(),
						)
					}
					return layout.Dimensions{Size: sz}
				}),
				layout.Stacked(func(gtx layout.Context) layout.Dimensions {
					gtx.Constraints.Min.X = gtx.Constraints.Max.X // fill full width so editor hit area covers the input
					gtx.Constraints.Min.Y = h
					return layout.Inset{
						Left: th.Space3, Right: th.Space3,
						Top: th.Space2, Bottom: th.Space2,
					}.Layout(gtx, func(gtx layout.Context) layout.Dimensions {
						if th.Shaper == nil {
							th.Shaper = text.NewShaper(text.NoSystemFonts(), text.WithCollection(defaultFonts()))
						}
						return inp.Editor.Layout(gtx, th.Shaper, font.Font{}, th.FontSize, op.CallOp{}, op.CallOp{})
					})
				}),
			)
		}),
	)
}
