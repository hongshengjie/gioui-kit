// Package theme provides a TailwindCSS/DaisyUI-inspired theming system for Gio UI.
//
// Color naming follows TailwindCSS conventions (e.g., Slate50, Blue500)
// with DaisyUI semantic tokens (Primary, Secondary, Accent, etc.).
package theme

import (
	"image/color"

	"gioui.org/text"
	"gioui.org/unit"
)

// Theme holds DaisyUI-style semantic color tokens and typography settings.
type Theme struct {
	// DaisyUI semantic colors
	Primary        color.NRGBA
	PrimaryContent color.NRGBA
	Secondary        color.NRGBA
	SecondaryContent color.NRGBA
	Accent        color.NRGBA
	AccentContent color.NRGBA
	Neutral        color.NRGBA
	NeutralContent color.NRGBA

	// State colors
	Info        color.NRGBA
	InfoContent color.NRGBA
	Success        color.NRGBA
	SuccessContent color.NRGBA
	Warning        color.NRGBA
	WarningContent color.NRGBA
	Error        color.NRGBA
	ErrorContent color.NRGBA

	// Surface colors
	Base100     color.NRGBA // bg base
	Base200     color.NRGBA // bg slightly darker
	Base300     color.NRGBA // bg more darker
	BaseContent color.NRGBA // text on base

	// Typography
	FontSize unit.Sp
	H1Size   unit.Sp
	H2Size   unit.Sp
	H3Size   unit.Sp
	H4Size   unit.Sp
	SmSize   unit.Sp
	XsSize   unit.Sp

	// Font collection for Gio shaper
	Shaper *text.Shaper

	// Spacing scale (Tailwind-like)
	Space0  unit.Dp // 0
	Space1  unit.Dp // 4
	Space2  unit.Dp // 8
	Space3  unit.Dp // 12
	Space4  unit.Dp // 16
	Space5  unit.Dp // 20
	Space6  unit.Dp // 24
	Space8  unit.Dp // 32
	Space10 unit.Dp // 40
	Space12 unit.Dp // 48
	Space16 unit.Dp // 64

	// Border radius
	RoundedNone unit.Dp
	RoundedSm   unit.Dp
	RoundedMd   unit.Dp
	RoundedLg   unit.Dp
	RoundedXl   unit.Dp
	Rounded2xl  unit.Dp
	RoundedFull unit.Dp
}

// ---------- Preset Themes (DaisyUI-inspired) ----------

// Light returns the default light theme (similar to DaisyUI "light").
func Light() *Theme {
	return &Theme{
		Primary:          Indigo500,
		PrimaryContent:   White,
		Secondary:        Purple500,
		SecondaryContent: White,
		Accent:           Cyan500,
		AccentContent:    White,
		Neutral:          Gray700,
		NeutralContent:   Gray100,

		Info:           Blue500,
		InfoContent:    White,
		Success:        Emerald500,
		SuccessContent: White,
		Warning:        Amber500,
		WarningContent: White,
		Error:          Red500,
		ErrorContent:   White,

		Base100:     White,
		Base200:     Gray100,
		Base300:     Gray300,
		BaseContent: Gray900,

		FontSize: 16,
		H1Size:   30,
		H2Size:   24,
		H3Size:   20,
		H4Size:   18,
		SmSize:   14,
		XsSize:   12,

		Space0:  0,
		Space1:  4,
		Space2:  8,
		Space3:  12,
		Space4:  16,
		Space5:  20,
		Space6:  24,
		Space8:  32,
		Space10: 40,
		Space12: 48,
		Space16: 64,

		RoundedNone: 0,
		RoundedSm:   2,
		RoundedMd:   6,
		RoundedLg:   8,
		RoundedXl:   12,
		Rounded2xl:  16,
		RoundedFull: 9999,
	}
}

// Dark returns a dark theme (similar to DaisyUI "dark").
func Dark() *Theme {
	return &Theme{
		Primary:          Indigo400,
		PrimaryContent:   White,
		Secondary:        Purple400,
		SecondaryContent: White,
		Accent:           Cyan400,
		AccentContent:    Slate900,
		Neutral:          Slate600,
		NeutralContent:   Slate200,

		Info:           Blue400,
		InfoContent:    Slate900,
		Success:        Emerald400,
		SuccessContent: Slate900,
		Warning:        Amber400,
		WarningContent: Slate900,
		Error:          Rose400,
		ErrorContent:   White,

		Base100:     Slate900,
		Base200:     Slate800,
		Base300:     Slate700,
		BaseContent: Slate100,

		FontSize: 16,
		H1Size:   30,
		H2Size:   24,
		H3Size:   20,
		H4Size:   18,
		SmSize:   14,
		XsSize:   12,

		Space0:  0,
		Space1:  4,
		Space2:  8,
		Space3:  12,
		Space4:  16,
		Space5:  20,
		Space6:  24,
		Space8:  32,
		Space10: 40,
		Space12: 48,
		Space16: 64,

		RoundedNone: 0,
		RoundedSm:   2,
		RoundedMd:   6,
		RoundedLg:   8,
		RoundedXl:   12,
		Rounded2xl:  16,
		RoundedFull: 9999,
	}
}

// Cupcake returns a pastel-friendly theme (DaisyUI "cupcake").
func Cupcake() *Theme {
	t := Light()
	t.Primary = rgb(0x65c3c8)
	t.PrimaryContent = rgb(0x223D40)
	t.Secondary = rgb(0xef9fbc)
	t.SecondaryContent = rgb(0x49242e)
	t.Accent = rgb(0xeeaf3a)
	t.AccentContent = rgb(0x452f10)
	t.Neutral = rgb(0x291334)
	t.NeutralContent = rgb(0xe8d8f0)
	t.Base100 = rgb(0xfaf7f5)
	t.Base200 = rgb(0xefeae6)
	t.Base300 = rgb(0xe7e2df)
	t.BaseContent = rgb(0x291334)
	return t
}

// Nord returns a Nord-palette theme.
func Nord() *Theme {
	t := Light()
	t.Primary = rgb(0x5E81AC)
	t.PrimaryContent = rgb(0xECEFF4)
	t.Secondary = rgb(0x81A1C1)
	t.SecondaryContent = rgb(0x2E3440)
	t.Accent = rgb(0x88C0D0)
	t.AccentContent = rgb(0x2E3440)
	t.Neutral = rgb(0x4C566A)
	t.NeutralContent = rgb(0xD8DEE9)
	t.Info = rgb(0x88C0D0)
	t.Success = rgb(0xA3BE8C)
	t.Warning = rgb(0xEBCB8B)
	t.Error = rgb(0xBF616A)
	t.Base100 = rgb(0xECEFF4)
	t.Base200 = rgb(0xE5E9F0)
	t.Base300 = rgb(0xD8DEE9)
	t.BaseContent = rgb(0x2E3440)
	return t
}
