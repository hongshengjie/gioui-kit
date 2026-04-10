package component

import (
	"image"
	"image/color"

	"gioui.org/font"
	"gioui.org/io/pointer"
	"gioui.org/layout"
	"gioui.org/op/clip"
	"gioui.org/op/paint"
	"gioui.org/unit"
	"gioui.org/widget"

	"github.com/ossprovider/gioui-kit/theme"
)

// FabPosition defines which screen corner the FAB is anchored to.
type FabPosition int

const (
	FabBottomRight FabPosition = iota
	FabBottomLeft
	FabTopRight
	FabTopLeft
)

// Fab is a DaisyUI-style floating action button rendered as a corner overlay.
// Call Layout after all main content with full-window constraints.
type Fab struct {
	Icon      *widget.Icon
	Label     string     // non-empty → extended pill-shaped FAB
	Tooltip   string     // label shown beside the button in FabGroup secondary actions
	Variant   BtnVariant
	Size      BtnSize
	Position  FabPosition
	Clickable *widget.Clickable
	th        *theme.Theme
}

// NewFab creates a FAB with primary variant, large size, at the bottom-right corner.
func NewFab(th *theme.Theme, click *widget.Clickable, icon *widget.Icon) *Fab {
	return &Fab{
		Icon:      icon,
		Variant:   BtnPrimary,
		Size:      BtnLg,
		Position:  FabBottomRight,
		Clickable: click,
		th:        th,
	}
}

func (f *Fab) WithVariant(v BtnVariant) *Fab  { f.Variant = v; return f }
func (f *Fab) WithSize(s BtnSize) *Fab         { f.Size = s; return f }
func (f *Fab) WithPosition(p FabPosition) *Fab { f.Position = p; return f }

// WithLabel makes the FAB extended (pill-shaped) with an icon and text label.
func (f *Fab) WithLabel(label string) *Fab { f.Label = label; return f }

// WithTooltip sets the label shown beside this action in a FabGroup.
func (f *Fab) WithTooltip(tip string) *Fab { f.Tooltip = tip; return f }

func (f *Fab) sizeDp() (btnDp, iconDp unit.Dp) {
	switch f.Size {
	case BtnXs:
		return 32, 14
	case BtnSm:
		return 40, 18
	case BtnMd:
		return 52, 22
	default: // BtnLg
		return 64, 28
	}
}

func (f *Fab) fabColors() (bg, fg color.NRGBA) {
	th := f.th
	switch f.Variant {
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
	default:
		return th.Neutral, th.NeutralContent
	}
}

func fabDirection(p FabPosition) layout.Direction {
	switch p {
	case FabBottomLeft:
		return layout.SW
	case FabTopRight:
		return layout.NE
	case FabTopLeft:
		return layout.NW
	default: // FabBottomRight
		return layout.SE
	}
}

// Layout renders the FAB as a corner overlay.
// Must be called with full-window gtx (after all main content).
func (f *Fab) Layout(gtx layout.Context) layout.Dimensions {
	th := f.th
	return fabDirection(f.Position).Layout(gtx, func(gtx layout.Context) layout.Dimensions {
		return layout.UniformInset(th.Space6).Layout(gtx, func(gtx layout.Context) layout.Dimensions {
			return f.draw(gtx)
		})
	})
}

// draw renders the button only (no corner positioning). Used by FabGroup.
func (f *Fab) draw(gtx layout.Context) layout.Dimensions {
	if f.Label != "" {
		return f.drawExtended(gtx)
	}
	return f.drawCircle(gtx)
}

func (f *Fab) drawCircle(gtx layout.Context) layout.Dimensions {
	btnDp, iconDp := f.sizeDp()
	bg, fg := f.fabColors()

	if f.Clickable.Hovered() {
		bg = theme.Lerp(bg, theme.Black, 0.1)
	}
	if f.Clickable.Pressed() {
		bg = theme.Lerp(bg, theme.Black, 0.2)
	}

	btnPx := gtx.Dp(btnDp)
	sz := image.Pt(btnPx, btnPx)
	r := btnPx / 2
	gtx.Constraints = layout.Exact(sz)

	return f.Clickable.Layout(gtx, func(gtx layout.Context) layout.Dimensions {
		return layout.Stack{Alignment: layout.Center}.Layout(gtx,
			layout.Expanded(func(gtx layout.Context) layout.Dimensions {
				drawFabShadow(gtx, sz, r)
				defer clip.UniformRRect(image.Rectangle{Max: sz}, r).Push(gtx.Ops).Pop()
				paint.ColorOp{Color: bg}.Add(gtx.Ops)
				paint.PaintOp{}.Add(gtx.Ops)
				pointer.CursorPointer.Add(gtx.Ops)
				return layout.Dimensions{Size: sz}
			}),
			layout.Stacked(func(gtx layout.Context) layout.Dimensions {
				iconPx := gtx.Dp(iconDp)
				gtx.Constraints = layout.Exact(image.Pt(iconPx, iconPx))
				if f.Icon != nil {
					return f.Icon.Layout(gtx, fg)
				}
				return layout.Dimensions{Size: image.Pt(iconPx, iconPx)}
			}),
		)
	})
}

func (f *Fab) drawExtended(gtx layout.Context) layout.Dimensions {
	th := f.th
	btnDp, iconDp := f.sizeDp()
	bg, fg := f.fabColors()

	if f.Clickable.Hovered() {
		bg = theme.Lerp(bg, theme.Black, 0.1)
	}
	if f.Clickable.Pressed() {
		bg = theme.Lerp(bg, theme.Black, 0.2)
	}

	btnPx := gtx.Dp(btnDp)
	r := btnPx / 2
	iconPx := gtx.Dp(iconDp)

	return f.Clickable.Layout(gtx, func(gtx layout.Context) layout.Dimensions {
		return layout.Stack{Alignment: layout.Center}.Layout(gtx,
			layout.Expanded(func(gtx layout.Context) layout.Dimensions {
				sz := gtx.Constraints.Min
				sz.Y = btnPx
				drawFabShadow(gtx, sz, r)
				defer clip.UniformRRect(image.Rectangle{Max: sz}, r).Push(gtx.Ops).Pop()
				paint.ColorOp{Color: bg}.Add(gtx.Ops)
				paint.PaintOp{}.Add(gtx.Ops)
				pointer.CursorPointer.Add(gtx.Ops)
				return layout.Dimensions{Size: sz}
			}),
			layout.Stacked(func(gtx layout.Context) layout.Dimensions {
				return layout.Inset{Left: th.Space5, Right: th.Space5}.Layout(gtx,
					func(gtx layout.Context) layout.Dimensions {
						return layout.Flex{Axis: layout.Horizontal, Alignment: layout.Middle}.Layout(gtx,
							layout.Rigid(func(gtx layout.Context) layout.Dimensions {
								gtx.Constraints = layout.Exact(image.Pt(iconPx, iconPx))
								if f.Icon != nil {
									return f.Icon.Layout(gtx, fg)
								}
								return layout.Dimensions{Size: image.Pt(iconPx, iconPx)}
							}),
							layout.Rigid(layout.Spacer{Width: th.Space2}.Layout),
							layout.Rigid(func(gtx layout.Context) layout.Dimensions {
								return drawText(gtx, th, f.Label, fg, th.SmSize, font.SemiBold)
							}),
						)
					},
				)
			}),
		)
	})
}

// drawFabShadow paints a soft circular drop-shadow beneath the FAB.
func drawFabShadow(gtx layout.Context, sz image.Point, radius int) {
	base := color.NRGBA{A: 80}
	for i := 3; i >= 1; i-- {
		expand := i * gtx.Dp(3)
		sr := image.Rectangle{
			Min: image.Pt(-expand/2, gtx.Dp(2)),
			Max: image.Pt(sz.X+expand/2, sz.Y+expand/2+gtx.Dp(4)),
		}
		c := base
		c.A = uint8(int(base.A) / (i + 1))
		rr := radius + expand/2
		stack := clip.UniformRRect(sr, rr).Push(gtx.Ops)
		paint.ColorOp{Color: c}.Add(gtx.Ops)
		paint.PaintOp{}.Add(gtx.Ops)
		stack.Pop()
	}
}

// ─── FabGroup ────────────────────────────────────────────────────────────────

// FabGroup is an expandable FAB that reveals secondary action buttons when clicked.
// The Main FAB toggles the expanded state; secondary Actions are shown above/below it.
type FabGroup struct {
	Main     *Fab
	Actions  []*Fab // secondary actions; Tooltip is shown as label alongside button
	Expanded bool
	th       *theme.Theme
}

// NewFabGroup creates a FabGroup. Pass the main FAB and any secondary action FABs.
func NewFabGroup(th *theme.Theme, main *Fab, actions ...*Fab) *FabGroup {
	return &FabGroup{
		Main:    main,
		Actions: actions,
		th:      th,
	}
}

// Layout renders the FAB group as a corner overlay.
// Must be called with full-window gtx (after all main content).
func (g *FabGroup) Layout(gtx layout.Context) layout.Dimensions {
	th := g.th
	if g.Main.Clickable.Clicked(gtx) {
		g.Expanded = !g.Expanded
	}

	isBottom := g.Main.Position == FabBottomRight || g.Main.Position == FabBottomLeft
	isRight := g.Main.Position == FabBottomRight || g.Main.Position == FabTopRight

	return fabDirection(g.Main.Position).Layout(gtx, func(gtx layout.Context) layout.Dimensions {
		return layout.UniformInset(th.Space6).Layout(gtx, func(gtx layout.Context) layout.Dimensions {
			// Build items: for bottom positions actions appear above main, else below.
			type item struct {
				action  *Fab
				isMain  bool
			}
			var items []item
			if isBottom {
				// Actions first (top), main last (bottom)
				for i := len(g.Actions) - 1; i >= 0; i-- {
					items = append(items, item{action: g.Actions[i]})
				}
				items = append(items, item{isMain: true})
			} else {
				items = append(items, item{isMain: true})
				for _, a := range g.Actions {
					items = append(items, item{action: a})
				}
			}

			children := make([]layout.FlexChild, 0, len(items)*2)
			for idx, it := range items {
				if idx > 0 {
					children = append(children, layout.Rigid(layout.Spacer{Height: th.Space3}.Layout))
				}
				if it.isMain {
					children = append(children, layout.Rigid(func(gtx layout.Context) layout.Dimensions {
						return g.Main.draw(gtx)
					}))
				} else if g.Expanded {
					children = append(children, layout.Rigid(func(gtx layout.Context) layout.Dimensions {
						return g.drawAction(gtx, it.action, isRight)
					}))
				}
			}
			return layout.Flex{Axis: layout.Vertical, Alignment: layout.Middle}.Layout(gtx, children...)
		})
	})
}

// drawAction renders a secondary action: [label] [button] or [button] [label].
func (g *FabGroup) drawAction(gtx layout.Context, action *Fab, labelOnLeft bool) layout.Dimensions {
	th := g.th

	if action.Tooltip == "" {
		return action.draw(gtx)
	}

	labelWidget := func(gtx layout.Context) layout.Dimensions {
		return drawFabActionLabel(gtx, th, action.Tooltip)
	}
	btnWidget := func(gtx layout.Context) layout.Dimensions {
		return action.draw(gtx)
	}

	return layout.Flex{Axis: layout.Horizontal, Alignment: layout.Middle}.Layout(gtx,
		layout.Rigid(func(gtx layout.Context) layout.Dimensions {
			if labelOnLeft {
				return labelWidget(gtx)
			}
			return btnWidget(gtx)
		}),
		layout.Rigid(layout.Spacer{Width: th.Space2}.Layout),
		layout.Rigid(func(gtx layout.Context) layout.Dimensions {
			if labelOnLeft {
				return btnWidget(gtx)
			}
			return labelWidget(gtx)
		}),
	)
}

// drawFabActionLabel renders a small pill-shaped label badge.
func drawFabActionLabel(gtx layout.Context, th *theme.Theme, label string) layout.Dimensions {
	radius := gtx.Dp(th.RoundedMd)
	bg := th.Neutral
	fg := th.NeutralContent

	return layout.Stack{}.Layout(gtx,
		layout.Expanded(func(gtx layout.Context) layout.Dimensions {
			sz := gtx.Constraints.Min
			if sz.X > 0 && sz.Y > 0 {
				defer clip.UniformRRect(image.Rectangle{Max: sz}, radius).Push(gtx.Ops).Pop()
				paint.ColorOp{Color: bg}.Add(gtx.Ops)
				paint.PaintOp{}.Add(gtx.Ops)
			}
			return layout.Dimensions{Size: sz}
		}),
		layout.Stacked(func(gtx layout.Context) layout.Dimensions {
			return layout.Inset{
				Top: th.Space1, Bottom: th.Space1,
				Left: th.Space3, Right: th.Space3,
			}.Layout(gtx, func(gtx layout.Context) layout.Dimensions {
				return drawText(gtx, th, label, fg, th.SmSize, font.Medium)
			})
		}),
	)
}
