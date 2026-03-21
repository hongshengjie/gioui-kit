// Package scaffold provides app-level layout scaffolds for Gio UI.
//
// These are high-level layout patterns commonly used in applications:
//
//	scaffold.AppShell{} - Full app shell with navbar + sidebar + content
//	scaffold.Navbar{}   - Top navigation bar
//	scaffold.Sidebar{}  - Side navigation panel
//	scaffold.Drawer{}   - Slide-in drawer overlay
//	scaffold.Modal{}    - Modal dialog overlay
//	scaffold.BottomNav{} - Mobile bottom navigation
package scaffold

import (
	"image"

	"gioui.org/layout"
	"gioui.org/op/clip"
	"gioui.org/op/paint"
	"gioui.org/unit"

	kit "github.com/hongshengjie/gioui-kit/layout"
	"github.com/hongshengjie/gioui-kit/theme"
)

// AppShell provides a complete application layout with optional
// navbar, sidebar, and main content area.
//
// Layout structure (wide screen):
//
//	┌─────────────── Navbar ───────────────┐
//	│                                      │
//	├──────────┬───────────────────────────┤
//	│          │                           │
//	│ Sidebar  │       Content Area        │
//	│          │                           │
//	│          │                           │
//	└──────────┴───────────────────────────┘
//
// On narrow screens (below HideSidebarBelow), the sidebar is omitted and
// the content area fills the full width.  Pair with a Drawer for mobile
// sidebar access.
type AppShell struct {
	Navbar             layout.Widget
	Sidebar            layout.Widget
	SidebarWidth       unit.Dp
	Content            layout.Widget
	// HideSidebarBelow hides the inline sidebar when the screen is narrower
	// than this breakpoint.  Defaults to BreakpointLg (< 1024 dp).
	HideSidebarBelow   kit.Breakpoint
	th                 *theme.Theme
}

func NewAppShell(th *theme.Theme) *AppShell {
	return &AppShell{
		SidebarWidth:     256,
		HideSidebarBelow: kit.BreakpointLg,
		th:               th,
	}
}

func (a *AppShell) WithNavbar(w layout.Widget) *AppShell {
	a.Navbar = w
	return a
}

func (a *AppShell) WithSidebar(w layout.Widget, width unit.Dp) *AppShell {
	a.Sidebar = w
	a.SidebarWidth = width
	return a
}

func (a *AppShell) WithContent(w layout.Widget) *AppShell {
	a.Content = w
	return a
}

// Layout renders the full app shell.
func (a *AppShell) Layout(gtx layout.Context) layout.Dimensions {
	th := a.th
	showSidebar := a.Sidebar != nil && kit.ScreenBreakpoint(gtx) >= a.HideSidebarBelow

	return layout.Flex{Axis: layout.Vertical}.Layout(gtx,
		// Navbar
		layout.Rigid(func(gtx layout.Context) layout.Dimensions {
			if a.Navbar == nil {
				return layout.Dimensions{}
			}
			return a.Navbar(gtx)
		}),
		// Body: Sidebar + Content
		layout.Flexed(1, func(gtx layout.Context) layout.Dimensions {
			return layout.Flex{Axis: layout.Horizontal}.Layout(gtx,
				// Sidebar (hidden on narrow screens)
				layout.Rigid(func(gtx layout.Context) layout.Dimensions {
					if !showSidebar {
						return layout.Dimensions{}
					}
					sideW := gtx.Dp(a.SidebarWidth)
					gtx.Constraints.Min.X = sideW
					gtx.Constraints.Max.X = sideW
					gtx.Constraints.Min.Y = gtx.Constraints.Max.Y

					sz := image.Pt(sideW, gtx.Constraints.Max.Y)
					paint.FillShape(gtx.Ops, th.Base200,
						clip.Rect{Max: sz}.Op(),
					)
					return a.Sidebar(gtx)
				}),
				// Divider (hidden with sidebar)
				layout.Rigid(func(gtx layout.Context) layout.Dimensions {
					if !showSidebar {
						return layout.Dimensions{}
					}
					sz := image.Pt(1, gtx.Constraints.Max.Y)
					paint.FillShape(gtx.Ops, th.Base300,
						clip.Rect{Max: sz}.Op(),
					)
					return layout.Dimensions{Size: sz}
				}),
				// Content
				layout.Flexed(1, func(gtx layout.Context) layout.Dimensions {
					if a.Content == nil {
						return layout.Dimensions{Size: gtx.Constraints.Max}
					}
					sz := gtx.Constraints.Max
					paint.FillShape(gtx.Ops, th.Base100,
						clip.Rect{Max: sz}.Op(),
					)
					return a.Content(gtx)
				}),
			)
		}),
	)
}
