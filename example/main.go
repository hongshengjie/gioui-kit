// Command demo shows a complete example application using gioui-kit.
package main

import (
	"fmt"
	"log"
	"net/http"
	_ "net/http/pprof" //nolint
	"os"

	"image/color"

	"gioui.org/app"
	"gioui.org/layout"
	"gioui.org/op"
	"gioui.org/unit"
	"gioui.org/widget"

	"github.com/hongshengjie/gioui-kit/component"
	kit "github.com/hongshengjie/gioui-kit/layout"
	"github.com/hongshengjie/gioui-kit/scaffold"
	"github.com/hongshengjie/gioui-kit/theme"
)

func main() {
	go func() {
		log.Println("pprof listening on :6060")
		log.Println(http.ListenAndServe("localhost:6060", nil))
	}()
	go func() {
		w := new(app.Window)
		w.Option(
			app.Title("GioUI Kit Demo"),
			app.Size(unit.Dp(1200), unit.Dp(800)),
		)
		if err := run(w); err != nil {
			log.Fatal(err)
		}
		os.Exit(0)
	}()
	app.Main()
}

// App holds all application state.
type App struct {
	th *theme.Theme

	// Shell
	shell   *scaffold.AppShell
	sidebar *scaffold.Sidebar
	navbar  *scaffold.Navbar

	// Overlays
	modal  *scaffold.Modal
	drawer *scaffold.Drawer
	toast  *scaffold.Toast

	// Scrolling
	scroll kit.ScrollY

	// Navigation
	pageIndex int // 0=Dashboard 1=Components 2=Layout 3=Forms 4=Settings

	// Components page: sub-tabs
	compTabs *component.Tabs

	// --- Button demo clickables ---
	btnPrimary   widget.Clickable
	btnSecondary widget.Clickable
	btnAccent    widget.Clickable
	btnOutline   widget.Clickable
	btnGhost     widget.Clickable
	btnLink      widget.Clickable
	btnError     widget.Clickable
	btnInfo      widget.Clickable
	btnSuccess   widget.Clickable
	btnWarning   widget.Clickable
	btnXs        widget.Clickable
	btnSm        widget.Clickable
	btnLg        widget.Clickable

	// --- Action buttons ---
	btnModal       widget.Clickable
	btnModalClose  widget.Clickable
	btnToast       widget.Clickable
	btnDrawer      widget.Clickable
	btnDrawerClose widget.Clickable

	// --- Mobile nav (hamburger + drawer items) ---
	btnHamburger      widget.Clickable
	btnDrawerDash     widget.Clickable
	btnDrawerComp     widget.Clickable
	btnDrawerLayout   widget.Clickable
	btnDrawerForms    widget.Clickable
	btnDrawerSettings widget.Clickable

	// --- Overview quick-start ---
	btnOverview1 widget.Clickable
	btnOverview2 widget.Clickable

	// --- Card embedded buttons ---
	btnCard1 widget.Clickable
	btnCard2 widget.Clickable
	btnCard3 widget.Clickable

	// --- Settings: theme picker ---
	btnLight   widget.Clickable
	btnDark    widget.Clickable
	btnCupcake widget.Clickable
	btnNord    widget.Clickable

	// --- Form editors ---
	editor1 widget.Editor
	editor2 widget.Editor
	editor3 widget.Editor
	editor4 widget.Editor

	// --- Toggles ---
	toggle1 widget.Bool
	toggle2 widget.Bool
	toggle3 widget.Bool

	// --- Progress ---
	progress float32
}

func NewApp() *App {
	th := theme.Light()

	sideItems := []scaffold.SidebarItem{
		{Label: "Dashboard", Icon: "◉", Active: true},
		{Label: "Components", Icon: "◫"},
		{Label: "Layout", Icon: "⊞"},
		{Label: "Forms", Icon: "✎"},
		{Label: "Settings", Icon: "⚙"},
	}

	a := &App{
		th:       th,
		navbar:   scaffold.NewNavbar(th),
		sidebar:  scaffold.NewSidebar(th, sideItems),
		modal:    scaffold.NewModal(th),
		drawer:   scaffold.NewDrawer(th),
		toast:    scaffold.NewToast(th),
		compTabs: component.NewTabs(th, []string{"Buttons", "Badges & Chips", "Alerts", "Avatars & Progress"}),
		progress: 0.65,
	}

	a.editor1.SingleLine = true
	a.editor2.SingleLine = true
	a.editor3.SingleLine = true
	a.editor4.SingleLine = true
	a.toggle1.Value = true

	a.sidebar.OnSelect = func(i int) { a.selectPage(i) }

	a.shell = scaffold.NewAppShell(th)
	a.scroll.List.Axis = layout.Vertical // zero-value is Horizontal; must set explicitly
	return a
}

// selectPage switches the active page and syncs the sidebar.
func (a *App) selectPage(i int) {
	a.pageIndex = i
	for j := range a.sidebar.Items {
		a.sidebar.Items[j].Active = j == i
	}
}

func run(w *app.Window) error {
	a := NewApp()
	var ops op.Ops

	for {
		switch e := w.Event().(type) {
		case app.DestroyEvent:
			return e.Err
		case app.FrameEvent:
			gtx := app.NewContext(&ops, e)

			// Hamburger → open drawer
			if a.btnHamburger.Clicked(gtx) {
				a.drawer.Open()
			}
			// Drawer nav items
			if a.btnDrawerDash.Clicked(gtx) {
				a.selectPage(0)
				a.drawer.Close()
			}
			if a.btnDrawerComp.Clicked(gtx) {
				a.selectPage(1)
				a.drawer.Close()
			}
			if a.btnDrawerLayout.Clicked(gtx) {
				a.selectPage(2)
				a.drawer.Close()
			}
			if a.btnDrawerForms.Clicked(gtx) {
				a.selectPage(3)
				a.drawer.Close()
			}
			if a.btnDrawerSettings.Clicked(gtx) {
				a.selectPage(4)
				a.drawer.Close()
			}

			// Modal
			if a.btnModal.Clicked(gtx) {
				a.modal.Show()
			}
			if a.btnModalClose.Clicked(gtx) {
				a.modal.Hide()
			}
			// Toast
			if a.btnToast.Clicked(gtx) {
				a.toast.Show("Operation completed successfully!")
			}
			// Drawer
			if a.btnDrawer.Clicked(gtx) {
				a.drawer.Open()
			}
			if a.btnDrawerClose.Clicked(gtx) {
				a.drawer.Close()
			}
			// Theme switcher
			if a.btnLight.Clicked(gtx) {
				*a.th = *theme.Light()
			}
			if a.btnDark.Clicked(gtx) {
				*a.th = *theme.Dark()
			}
			if a.btnCupcake.Clicked(gtx) {
				*a.th = *theme.Cupcake()
			}
			if a.btnNord.Clicked(gtx) {
				*a.th = *theme.Nord()
			}

			a.layout(gtx)
			e.Frame(gtx.Ops)
		}
	}
}

func (a *App) layout(gtx layout.Context) layout.Dimensions {
	th := a.th

	a.shell.Navbar = func(gtx layout.Context) layout.Dimensions {
		mobile := kit.ScreenBreakpoint(gtx) < kit.BreakpointLg
		return a.navbar.Layout(gtx,
			// start: hamburger on mobile, brand on desktop
			func(gtx layout.Context) layout.Dimensions {
				if mobile {
					return component.NewButton(th, &a.btnHamburger, "≡").WithVariant(component.BtnGhost).Layout(gtx)
				}
				return component.NewText(th, "GioUI Kit").H3().Bold().WithColor(th.Primary).Layout(gtx)
			},
			// center: brand on mobile, empty on desktop
			func(gtx layout.Context) layout.Dimensions {
				if !mobile {
					return layout.Dimensions{}
				}
				return component.NewText(th, "GioUI Kit").H3().Bold().WithColor(th.Primary).Layout(gtx)
			},
			// end: badge+avatar on desktop, avatar only on mobile
			func(gtx layout.Context) layout.Dimensions {
				return kit.FlexRow{Gap: 8, Alignment: kit.ItemsCenter}.Layout(gtx,
					kit.Rigid(func(gtx layout.Context) layout.Dimensions {
						if mobile {
							return layout.Dimensions{}
						}
						return component.NewBadge(th, "v0.1.0").WithVariant(component.BadgePrimary).Layout(gtx)
					}),
					kit.Rigid(func(gtx layout.Context) layout.Dimensions {
						return component.NewAvatar(th, "GK").Layout(gtx)
					}),
				)
			},
		)
	}

	a.shell.Sidebar = func(gtx layout.Context) layout.Dimensions {
		return a.sidebar.Layout(gtx)
	}
	a.shell.SidebarWidth = 220

	a.shell.Content = func(gtx layout.Context) layout.Dimensions {
		return a.layoutContent(gtx)
	}

	dims := a.shell.Layout(gtx)

	// Overlay: Drawer
	a.drawer.Layout(gtx, func(gtx layout.Context) layout.Dimensions {
		th := a.th
		return layout.Inset{
			Top: th.Space4, Left: th.Space4, Right: th.Space4,
		}.Layout(gtx, func(gtx layout.Context) layout.Dimensions {
			return kit.FlexCol{Gap: 16}.Layout(gtx,
				kit.Rigid(func(gtx layout.Context) layout.Dimensions {
					return kit.FlexRow{Gap: 8, Alignment: kit.ItemsCenter}.Layout(gtx,
						kit.Grow(func(gtx layout.Context) layout.Dimensions {
							return component.NewText(th, "Navigation").H3().Bold().Layout(gtx)
						}),
						kit.Rigid(func(gtx layout.Context) layout.Dimensions {
							return component.NewButton(th, &a.btnDrawerClose, "✕").WithVariant(component.BtnGhost).WithSize(component.BtnSm).Layout(gtx)
						}),
					)
				}),
				kit.Rigid(func(gtx layout.Context) layout.Dimensions {
					return kit.DividerH{Color: th.Base300}.Layout(gtx)
				}),
				kit.Rigid(func(gtx layout.Context) layout.Dimensions {
					return component.NewButton(th, &a.btnDrawerDash, "◉  Dashboard").WithVariant(component.BtnGhost).Layout(gtx)
				}),
				kit.Rigid(func(gtx layout.Context) layout.Dimensions {
					return component.NewButton(th, &a.btnDrawerComp, "◫  Components").WithVariant(component.BtnGhost).Layout(gtx)
				}),
				kit.Rigid(func(gtx layout.Context) layout.Dimensions {
					return component.NewButton(th, &a.btnDrawerLayout, "⊞  Layout").WithVariant(component.BtnGhost).Layout(gtx)
				}),
				kit.Rigid(func(gtx layout.Context) layout.Dimensions {
					return component.NewButton(th, &a.btnDrawerForms, "✎  Forms").WithVariant(component.BtnGhost).Layout(gtx)
				}),
				kit.Rigid(func(gtx layout.Context) layout.Dimensions {
					return component.NewButton(th, &a.btnDrawerSettings, "⚙  Settings").WithVariant(component.BtnGhost).Layout(gtx)
				}),
			)
		})
	})

	// Overlay: Modal
	a.modal.Layout(gtx, func(gtx layout.Context) layout.Dimensions {
		th := a.th
		return kit.FlexCol{Gap: 16}.Layout(gtx,
			kit.Rigid(func(gtx layout.Context) layout.Dimensions {
				return component.NewText(th, "Modal Dialog").H3().Bold().Layout(gtx)
			}),
			kit.Rigid(func(gtx layout.Context) layout.Dimensions {
				return component.NewText(th, "This is a DaisyUI-style modal rendered with Gio. Click the backdrop or Close button to dismiss.").Sm().WithColor(theme.Gray500).Layout(gtx)
			}),
			kit.Rigid(func(gtx layout.Context) layout.Dimensions {
				return kit.FlexRow{Gap: 8}.Layout(gtx,
					kit.Rigid(func(gtx layout.Context) layout.Dimensions {
						return component.NewButton(th, &a.btnModalClose, "Close").WithVariant(component.BtnPrimary).Layout(gtx)
					}),
					kit.Rigid(func(gtx layout.Context) layout.Dimensions {
						return component.NewButton(th, &a.btnModalClose, "Cancel").WithVariant(component.BtnGhost).Layout(gtx)
					}),
				)
			}),
		)
	})

	// Overlay: Toast
	a.toast.Layout(gtx)

	return dims
}

// layoutContent routes to the selected page.
func (a *App) layoutContent(gtx layout.Context) layout.Dimensions {
	return a.scroll.List.Layout(gtx, 1, func(gtx layout.Context, _ int) layout.Dimensions {
		th := a.th
		hPad := th.Space8
		if kit.ScreenBreakpoint(gtx) < kit.BreakpointMd {
			hPad = th.Space4
		}
		return layout.Inset{
			Top: th.Space6, Bottom: th.Space8,
			Left: hPad, Right: hPad,
		}.Layout(gtx, func(gtx layout.Context) layout.Dimensions {
			switch a.pageIndex {
			case 1:
				return a.pageComponents(gtx)
			case 2:
				return a.pageLayout(gtx)
			case 3:
				return a.pageForms(gtx)
			case 4:
				return a.pageSettings(gtx)
			default:
				return a.pageDashboard(gtx)
			}
		})
	})
}

// ─── Page: Dashboard ────────────────────────────────────────────────────────

func (a *App) pageDashboard(gtx layout.Context) layout.Dimensions {
	th := a.th
	return kit.FlexCol{Gap: 32}.Layout(gtx,
		// Hero
		kit.Rigid(func(gtx layout.Context) layout.Dimensions {
			return kit.FlexCol{Gap: 8}.Layout(gtx,
				kit.Rigid(func(gtx layout.Context) layout.Dimensions {
					return scaffold.NewBreadcrumb(th, "Home", "Dashboard").Layout(gtx)
				}),
				kit.Rigid(func(gtx layout.Context) layout.Dimensions {
					return component.NewText(th, "Welcome to GioUI Kit").H1().Layout(gtx)
				}),
				kit.Rigid(func(gtx layout.Context) layout.Dimensions {
					return component.NewText(th,
						"A comprehensive component library for Gio, inspired by TailwindCSS and DaisyUI.",
					).Sm().WithColor(theme.Gray500).Layout(gtx)
				}),
				kit.Rigid(func(gtx layout.Context) layout.Dimensions {
					return layout.Inset{Top: th.Space2}.Layout(gtx, func(gtx layout.Context) layout.Dimensions {
						return kit.FlexRow{Gap: 8, Alignment: kit.ItemsCenter}.Layout(gtx,
							kit.Rigid(func(gtx layout.Context) layout.Dimensions {
								return component.NewButton(th, &a.btnOverview1, "Browse Components").WithVariant(component.BtnPrimary).Layout(gtx)
							}),
							kit.Rigid(func(gtx layout.Context) layout.Dimensions {
								return component.NewButton(th, &a.btnOverview2, "View Source").WithVariant(component.BtnOutline).Layout(gtx)
							}),
						)
					})
				}),
			)
		}),

		// Stat cards — 1 col mobile → 2 sm → 4 md+
		kit.Rigid(func(gtx layout.Context) layout.Dimensions {
			return kit.Grid{Cols: 1, SmCols: 2, MdCols: 4, Gap: 16}.Layout(gtx,
				statCard(th, "12", "Components", theme.Blue500),
				statCard(th, "4", "Themes", theme.Purple500),
				statCard(th, "8", "Layout Types", theme.Emerald500),
				statCard(th, "6", "Form Controls", theme.Amber500),
			)
		}),

		// Recent activity
		kit.Rigid(func(gtx layout.Context) layout.Dimensions {
			return sectionCard(th, "Recent Activity", "Latest component updates", func(gtx layout.Context) layout.Dimensions {
				return kit.FlexCol{Gap: 12}.Layout(gtx,
					kit.Rigid(func(gtx layout.Context) layout.Dimensions {
						return component.NewAlert(th, "Button layout fixed — text no longer clipped by rounded corners.", component.AlertSuccess).Layout(gtx)
					}),
					kit.Rigid(func(gtx layout.Context) layout.Dimensions {
						return component.NewAlert(th, "Modal backdrop click now closes the dialog.", component.AlertInfo).Layout(gtx)
					}),
					kit.Rigid(func(gtx layout.Context) layout.Dimensions {
						return component.NewAlert(th, "Toast auto-dismisses after 3 seconds.", component.AlertInfo).Layout(gtx)
					}),
					kit.Rigid(func(gtx layout.Context) layout.Dimensions {
						return component.NewAlert(th, "Memory explosion bug fixed — clip.Stroke replaced with paint-over.", component.AlertWarning).Layout(gtx)
					}),
				)
			})(gtx)
		}),

		// Quick badges
		kit.Rigid(func(gtx layout.Context) layout.Dimensions {
			return sectionCard(th, "Tech Stack", "Built with", func(gtx layout.Context) layout.Dimensions {
				return kit.FlexRow{Gap: 8, Alignment: kit.ItemsCenter}.Layout(gtx,
					kit.Rigid(func(gtx layout.Context) layout.Dimensions {
						return component.NewBadge(th, "Go 1.22+").WithVariant(component.BadgePrimary).Layout(gtx)
					}),
					kit.Rigid(func(gtx layout.Context) layout.Dimensions {
						return component.NewBadge(th, "Gio v0.9").WithVariant(component.BadgeAccent).Layout(gtx)
					}),
					kit.Rigid(func(gtx layout.Context) layout.Dimensions {
						return component.NewBadge(th, "TailwindCSS").WithVariant(component.BadgeInfo).Layout(gtx)
					}),
					kit.Rigid(func(gtx layout.Context) layout.Dimensions {
						return component.NewBadge(th, "DaisyUI").WithVariant(component.BadgeSecondary).Layout(gtx)
					}),
					kit.Rigid(func(gtx layout.Context) layout.Dimensions {
						return component.NewBadge(th, "MIT License").WithVariant(component.BadgeSuccess).Layout(gtx)
					}),
				)
			})(gtx)
		}),
	)
}

func statCard(th *theme.Theme, value, label string, accent color.NRGBA) layout.Widget {
	return func(gtx layout.Context) layout.Dimensions {
		card := component.NewCard(th).WithBorder()
		return card.Layout(gtx, func(gtx layout.Context) layout.Dimensions {
			return kit.FlexCol{Gap: 4}.Layout(gtx,
				kit.Rigid(func(gtx layout.Context) layout.Dimensions {
					return component.NewText(th, value).H1().WithColor(accent).Layout(gtx)
				}),
				kit.Rigid(func(gtx layout.Context) layout.Dimensions {
					return component.NewText(th, label).Sm().WithColor(theme.Gray500).Layout(gtx)
				}),
			)
		})
	}
}

// ─── Page: Components ───────────────────────────────────────────────────────

func (a *App) pageComponents(gtx layout.Context) layout.Dimensions {
	th := a.th
	return kit.FlexCol{Gap: 32}.Layout(gtx,
		kit.Rigid(func(gtx layout.Context) layout.Dimensions {
			return pageHeader(th, gtx, "Components", "Dashboard / Components",
				"A complete showcase of all DaisyUI-inspired components.")
		}),

		// Sub-tabs
		kit.Rigid(func(gtx layout.Context) layout.Dimensions {
			return a.compTabs.Layout(gtx)
		}),

		// Tab content
		kit.Rigid(func(gtx layout.Context) layout.Dimensions {
			switch a.compTabs.Selected {
			case 1:
				return a.sectionBadgesChips(gtx)
			case 2:
				return a.sectionAlerts(gtx)
			case 3:
				return a.sectionAvatarsProgress(gtx)
			default:
				return a.sectionButtons(gtx)
			}
		}),
	)
}

func (a *App) sectionButtons(gtx layout.Context) layout.Dimensions {
	th := a.th
	return sectionCard(th, "Buttons", "All variants, sizes, and states", func(gtx layout.Context) layout.Dimensions {
		return kit.FlexCol{Gap: 24}.Layout(gtx,
			// Variants
			kit.Rigid(func(gtx layout.Context) layout.Dimensions {
				return subSection(th, gtx, "Variants", func(gtx layout.Context) layout.Dimensions {
					return kit.FlexRow{Gap: 8, Alignment: kit.ItemsCenter}.Layout(gtx,
						kit.Rigid(func(gtx layout.Context) layout.Dimensions {
							return component.NewButton(th, &a.btnPrimary, "Primary").WithVariant(component.BtnPrimary).Layout(gtx)
						}),
						kit.Rigid(func(gtx layout.Context) layout.Dimensions {
							return component.NewButton(th, &a.btnSecondary, "Secondary").WithVariant(component.BtnSecondary).Layout(gtx)
						}),
						kit.Rigid(func(gtx layout.Context) layout.Dimensions {
							return component.NewButton(th, &a.btnAccent, "Accent").WithVariant(component.BtnAccent).Layout(gtx)
						}),
						kit.Rigid(func(gtx layout.Context) layout.Dimensions {
							return component.NewButton(th, &a.btnInfo, "Info").WithVariant(component.BtnInfo).Layout(gtx)
						}),
						kit.Rigid(func(gtx layout.Context) layout.Dimensions {
							return component.NewButton(th, &a.btnSuccess, "Success").WithVariant(component.BtnSuccess).Layout(gtx)
						}),
						kit.Rigid(func(gtx layout.Context) layout.Dimensions {
							return component.NewButton(th, &a.btnWarning, "Warning").WithVariant(component.BtnWarning).Layout(gtx)
						}),
						kit.Rigid(func(gtx layout.Context) layout.Dimensions {
							return component.NewButton(th, &a.btnError, "Error").WithVariant(component.BtnError).Layout(gtx)
						}),
					)
				})
			}),
			// Ghost, Link, Outline
			kit.Rigid(func(gtx layout.Context) layout.Dimensions {
				return subSection(th, gtx, "Soft variants", func(gtx layout.Context) layout.Dimensions {
					return kit.FlexRow{Gap: 8, Alignment: kit.ItemsCenter}.Layout(gtx,
						kit.Rigid(func(gtx layout.Context) layout.Dimensions {
							return component.NewButton(th, &a.btnOutline, "Outline").WithVariant(component.BtnOutline).Layout(gtx)
						}),
						kit.Rigid(func(gtx layout.Context) layout.Dimensions {
							return component.NewButton(th, &a.btnGhost, "Ghost").WithVariant(component.BtnGhost).Layout(gtx)
						}),
						kit.Rigid(func(gtx layout.Context) layout.Dimensions {
							return component.NewButton(th, &a.btnLink, "Link").WithVariant(component.BtnLink).Layout(gtx)
						}),
					)
				})
			}),
			// Sizes
			kit.Rigid(func(gtx layout.Context) layout.Dimensions {
				return subSection(th, gtx, "Sizes", func(gtx layout.Context) layout.Dimensions {
					return kit.FlexRow{Gap: 8, Alignment: kit.ItemsCenter}.Layout(gtx,
						kit.Rigid(func(gtx layout.Context) layout.Dimensions {
							return component.NewButton(th, &a.btnXs, "XSmall").WithVariant(component.BtnPrimary).WithSize(component.BtnXs).Layout(gtx)
						}),
						kit.Rigid(func(gtx layout.Context) layout.Dimensions {
							return component.NewButton(th, &a.btnSm, "Small").WithVariant(component.BtnPrimary).WithSize(component.BtnSm).Layout(gtx)
						}),
						kit.Rigid(func(gtx layout.Context) layout.Dimensions {
							return component.NewButton(th, &a.btnPrimary, "Medium").WithVariant(component.BtnPrimary).WithSize(component.BtnMd).Layout(gtx)
						}),
						kit.Rigid(func(gtx layout.Context) layout.Dimensions {
							return component.NewButton(th, &a.btnLg, "Large").WithVariant(component.BtnPrimary).WithSize(component.BtnLg).Layout(gtx)
						}),
					)
				})
			}),
			// States
			kit.Rigid(func(gtx layout.Context) layout.Dimensions {
				return subSection(th, gtx, "States", func(gtx layout.Context) layout.Dimensions {
					disBtn := component.NewButton(th, &a.btnCard1, "Disabled")
					disBtn.WithVariant(component.BtnPrimary)
					disBtn.Disabled = true
					loadBtn := component.NewButton(th, &a.btnCard2, "Loading...")
					loadBtn.WithVariant(component.BtnSecondary)
					loadBtn.Loading = true
					return kit.FlexRow{Gap: 8, Alignment: kit.ItemsCenter}.Layout(gtx,
						kit.Rigid(func(gtx layout.Context) layout.Dimensions {
							return disBtn.Layout(gtx)
						}),
						kit.Rigid(func(gtx layout.Context) layout.Dimensions {
							return loadBtn.Layout(gtx)
						}),
						kit.Rigid(func(gtx layout.Context) layout.Dimensions {
							return component.NewButton(th, &a.btnModal, "Open Modal").WithVariant(component.BtnPrimary).Layout(gtx)
						}),
						kit.Rigid(func(gtx layout.Context) layout.Dimensions {
							return component.NewButton(th, &a.btnToast, "Show Toast").WithVariant(component.BtnSuccess).Layout(gtx)
						}),
					)
				})
			}),
		)
	})(gtx)
}

func (a *App) sectionBadgesChips(gtx layout.Context) layout.Dimensions {
	th := a.th
	return kit.FlexCol{Gap: 24}.Layout(gtx,
		// Badges
		kit.Rigid(func(gtx layout.Context) layout.Dimensions {
			return sectionCard(th, "Badges", "Status indicators and tags", func(gtx layout.Context) layout.Dimensions {
				return kit.FlexCol{Gap: 16}.Layout(gtx,
					kit.Rigid(func(gtx layout.Context) layout.Dimensions {
						return subSection(th, gtx, "Variants", func(gtx layout.Context) layout.Dimensions {
							return kit.FlexRow{Gap: 8, Alignment: kit.ItemsCenter}.Layout(gtx,
								kit.Rigid(func(gtx layout.Context) layout.Dimensions {
									return component.NewBadge(th, "Default").Layout(gtx)
								}),
								kit.Rigid(func(gtx layout.Context) layout.Dimensions {
									return component.NewBadge(th, "Primary").WithVariant(component.BadgePrimary).Layout(gtx)
								}),
								kit.Rigid(func(gtx layout.Context) layout.Dimensions {
									return component.NewBadge(th, "Secondary").WithVariant(component.BadgeSecondary).Layout(gtx)
								}),
								kit.Rigid(func(gtx layout.Context) layout.Dimensions {
									return component.NewBadge(th, "Accent").WithVariant(component.BadgeAccent).Layout(gtx)
								}),
								kit.Rigid(func(gtx layout.Context) layout.Dimensions {
									return component.NewBadge(th, "Info").WithVariant(component.BadgeInfo).Layout(gtx)
								}),
								kit.Rigid(func(gtx layout.Context) layout.Dimensions {
									return component.NewBadge(th, "Success").WithVariant(component.BadgeSuccess).Layout(gtx)
								}),
								kit.Rigid(func(gtx layout.Context) layout.Dimensions {
									return component.NewBadge(th, "Warning").WithVariant(component.BadgeWarning).Layout(gtx)
								}),
								kit.Rigid(func(gtx layout.Context) layout.Dimensions {
									return component.NewBadge(th, "Error").WithVariant(component.BadgeError).Layout(gtx)
								}),
								kit.Rigid(func(gtx layout.Context) layout.Dimensions {
									return component.NewBadge(th, "Ghost").WithVariant(component.BadgeGhost).Layout(gtx)
								}),
							)
						})
					}),
				)
			})(gtx)
		}),
		// Chips
		kit.Rigid(func(gtx layout.Context) layout.Dimensions {
			return sectionCard(th, "Chips", "Removable tag components", func(gtx layout.Context) layout.Dimensions {
				return kit.FlexRow{Gap: 8, Alignment: kit.ItemsCenter}.Layout(gtx,
					kit.Rigid(func(gtx layout.Context) layout.Dimensions {
						return component.NewChip(th, "Design").Layout(gtx)
					}),
					kit.Rigid(func(gtx layout.Context) layout.Dimensions {
						return component.NewChip(th, "Engineering").Layout(gtx)
					}),
					kit.Rigid(func(gtx layout.Context) layout.Dimensions {
						return component.NewChip(th, "Product").Layout(gtx)
					}),
					kit.Rigid(func(gtx layout.Context) layout.Dimensions {
						return component.NewChip(th, "Go").Layout(gtx)
					}),
					kit.Rigid(func(gtx layout.Context) layout.Dimensions {
						return component.NewChip(th, "Gio UI").Layout(gtx)
					}),
					kit.Rigid(func(gtx layout.Context) layout.Dimensions {
						return component.NewChip(th, "TailwindCSS").Layout(gtx)
					}),
				)
			})(gtx)
		}),
		// Skeleton
		kit.Rigid(func(gtx layout.Context) layout.Dimensions {
			return sectionCard(th, "Skeleton", "Loading placeholder", func(gtx layout.Context) layout.Dimensions {
				return kit.FlexCol{Gap: 8}.Layout(gtx,
					kit.Rigid(func(gtx layout.Context) layout.Dimensions {
						return component.NewSkeleton(th).Layout(gtx)
					}),
					kit.Rigid(func(gtx layout.Context) layout.Dimensions {
						s := component.NewSkeleton(th)
						s.Width = 300
						s.Height = 14
						return s.Layout(gtx)
					}),
					kit.Rigid(func(gtx layout.Context) layout.Dimensions {
						s := component.NewSkeleton(th)
						s.Width = 160
						s.Height = 14
						return s.Layout(gtx)
					}),
				)
			})(gtx)
		}),
	)
}

func (a *App) sectionAlerts(gtx layout.Context) layout.Dimensions {
	th := a.th
	return sectionCard(th, "Alerts", "Notification banners for user feedback", func(gtx layout.Context) layout.Dimensions {
		return kit.FlexCol{Gap: 12}.Layout(gtx,
			kit.Rigid(func(gtx layout.Context) layout.Dimensions {
				return component.NewAlert(th, "This is an informational message.", component.AlertInfo).Layout(gtx)
			}),
			kit.Rigid(func(gtx layout.Context) layout.Dimensions {
				return component.NewAlert(th, "Operation completed successfully! Your changes have been saved.", component.AlertSuccess).Layout(gtx)
			}),
			kit.Rigid(func(gtx layout.Context) layout.Dimensions {
				return component.NewAlert(th, "Please review your input before submitting.", component.AlertWarning).Layout(gtx)
			}),
			kit.Rigid(func(gtx layout.Context) layout.Dimensions {
				return component.NewAlert(th, "An error occurred. Please try again or contact support.", component.AlertError).Layout(gtx)
			}),
		)
	})(gtx)
}

func (a *App) sectionAvatarsProgress(gtx layout.Context) layout.Dimensions {
	th := a.th
	return kit.FlexCol{Gap: 24}.Layout(gtx,
		kit.Rigid(func(gtx layout.Context) layout.Dimensions {
			return sectionCard(th, "Avatars", "User profile circles", func(gtx layout.Context) layout.Dimensions {
				return kit.FlexRow{Gap: 16, Alignment: kit.ItemsCenter}.Layout(gtx,
					kit.Rigid(func(gtx layout.Context) layout.Dimensions {
						a := component.NewAvatar(th, "XS")
						a.Size = component.AvatarXs
						return a.Layout(gtx)
					}),
					kit.Rigid(func(gtx layout.Context) layout.Dimensions {
						a := component.NewAvatar(th, "SM")
						a.Size = component.AvatarSm
						return a.Layout(gtx)
					}),
					kit.Rigid(func(gtx layout.Context) layout.Dimensions {
						return component.NewAvatar(th, "MD").Layout(gtx)
					}),
					kit.Rigid(func(gtx layout.Context) layout.Dimensions {
						a := component.NewAvatar(th, "LG")
						a.Size = component.AvatarLg
						return a.Layout(gtx)
					}),
					kit.Rigid(func(gtx layout.Context) layout.Dimensions {
						return component.NewText(th, "XS  SM  MD  LG").Sm().WithColor(theme.Gray400).Layout(gtx)
					}),
				)
			})(gtx)
		}),
		kit.Rigid(func(gtx layout.Context) layout.Dimensions {
			return sectionCard(th, "Progress", "Progress indicators", func(gtx layout.Context) layout.Dimensions {
				return kit.FlexCol{Gap: 12}.Layout(gtx,
					kit.Rigid(func(gtx layout.Context) layout.Dimensions {
						return kit.FlexRow{Gap: 8, Alignment: kit.ItemsCenter}.Layout(gtx,
							kit.Rigid(func(gtx layout.Context) layout.Dimensions {
								return component.NewText(th, fmt.Sprintf("%.0f%%", a.progress*100)).Sm().Bold().Layout(gtx)
							}),
							kit.Grow(func(gtx layout.Context) layout.Dimensions {
								return component.NewProgress(th, a.progress).Layout(gtx)
							}),
						)
					}),
					kit.Rigid(func(gtx layout.Context) layout.Dimensions {
						p := component.NewProgress(th, 0.4)
						p.Variant = component.ProgressSuccess
						return kit.FlexRow{Gap: 8, Alignment: kit.ItemsCenter}.Layout(gtx,
							kit.Rigid(func(gtx layout.Context) layout.Dimensions {
								return component.NewText(th, "40%").Sm().WithColor(th.Success).Layout(gtx)
							}),
							kit.Grow(func(gtx layout.Context) layout.Dimensions { return p.Layout(gtx) }),
						)
					}),
					kit.Rigid(func(gtx layout.Context) layout.Dimensions {
						p := component.NewProgress(th, 0.75)
						p.Variant = component.ProgressWarning
						return kit.FlexRow{Gap: 8, Alignment: kit.ItemsCenter}.Layout(gtx,
							kit.Rigid(func(gtx layout.Context) layout.Dimensions {
								return component.NewText(th, "75%").Sm().WithColor(th.Warning).Layout(gtx)
							}),
							kit.Grow(func(gtx layout.Context) layout.Dimensions { return p.Layout(gtx) }),
						)
					}),
					kit.Rigid(func(gtx layout.Context) layout.Dimensions {
						p := component.NewProgress(th, 0.25)
						p.Variant = component.ProgressError
						return kit.FlexRow{Gap: 8, Alignment: kit.ItemsCenter}.Layout(gtx,
							kit.Rigid(func(gtx layout.Context) layout.Dimensions {
								return component.NewText(th, "25%").Sm().WithColor(th.Error).Layout(gtx)
							}),
							kit.Grow(func(gtx layout.Context) layout.Dimensions { return p.Layout(gtx) }),
						)
					}),
				)
			})(gtx)
		}),
	)
}

// ─── Page: Layout ───────────────────────────────────────────────────────────

func (a *App) pageLayout(gtx layout.Context) layout.Dimensions {
	th := a.th
	return kit.FlexCol{Gap: 32}.Layout(gtx,
		kit.Rigid(func(gtx layout.Context) layout.Dimensions {
			return pageHeader(th, gtx, "Layout", "Dashboard / Layout",
				"Grid and Flex layout primitives inspired by TailwindCSS.")
		}),

		// Cards section
		kit.Rigid(func(gtx layout.Context) layout.Dimensions {
			return sectionCard(th, "Cards", "Content containers", func(gtx layout.Context) layout.Dimensions {
				return kit.Grid{Cols: 1, MdCols: 2, LgCols: 3, Gap: 16}.Layout(gtx,
					func(gtx layout.Context) layout.Dimensions {
						return component.NewCard(th).WithBorder().CardWithHeader(gtx, "Bordered Card",
							func(gtx layout.Context) layout.Dimensions {
								return component.NewText(th, "This card has a visible border. Suitable for lighter backgrounds.").Sm().WithColor(theme.Gray500).Layout(gtx)
							})
					},
					func(gtx layout.Context) layout.Dimensions {
						return component.NewCard(th).CardWithHeader(gtx, "Default Card",
							func(gtx layout.Context) layout.Dimensions {
								return kit.FlexCol{Gap: 12}.Layout(gtx,
									kit.Rigid(func(gtx layout.Context) layout.Dimensions {
										return component.NewText(th, "Cards can contain any content including nested components.").Sm().WithColor(theme.Gray500).Layout(gtx)
									}),
									kit.Rigid(func(gtx layout.Context) layout.Dimensions {
										return component.NewButton(th, &a.btnCard1, "Action").WithVariant(component.BtnPrimary).WithSize(component.BtnSm).Layout(gtx)
									}),
								)
							})
					},
					func(gtx layout.Context) layout.Dimensions {
						return component.NewCard(th).WithCompact().CardWithHeader(gtx, "Compact Card",
							func(gtx layout.Context) layout.Dimensions {
								return kit.FlexCol{Gap: 8}.Layout(gtx,
									kit.Rigid(func(gtx layout.Context) layout.Dimensions {
										return component.NewText(th, "Compact variant with less padding.").Sm().WithColor(theme.Gray500).Layout(gtx)
									}),
									kit.Rigid(func(gtx layout.Context) layout.Dimensions {
										return kit.FlexRow{Gap: 4}.Layout(gtx,
											kit.Rigid(func(gtx layout.Context) layout.Dimensions {
												return component.NewBadge(th, "Go").WithVariant(component.BadgePrimary).Layout(gtx)
											}),
											kit.Rigid(func(gtx layout.Context) layout.Dimensions {
												return component.NewBadge(th, "Gio").WithVariant(component.BadgeAccent).Layout(gtx)
											}),
										)
									}),
								)
							})
					},
				)
			})(gtx)
		}),

		// Grid examples
		kit.Rigid(func(gtx layout.Context) layout.Dimensions {
			return sectionCard(th, "Grid Layout", "grid-cols-1 / grid-cols-2 / grid-cols-3", func(gtx layout.Context) layout.Dimensions {
				return kit.FlexCol{Gap: 20}.Layout(gtx,
					kit.Rigid(func(gtx layout.Context) layout.Dimensions {
						return subSection(th, gtx, "1 Column", func(gtx layout.Context) layout.Dimensions {
							return kit.Grid{Cols: 1, Gap: 8}.Layout(gtx,
								gridBox(th, "Col 1 / 1", theme.Blue100),
							)
						})
					}),
					kit.Rigid(func(gtx layout.Context) layout.Dimensions {
						return subSection(th, gtx, "2 Columns", func(gtx layout.Context) layout.Dimensions {
							return kit.Grid{Cols: 2, Gap: 8}.Layout(gtx,
								gridBox(th, "Col 1 / 2", theme.Blue100),
								gridBox(th, "Col 2 / 2", theme.Blue200),
							)
						})
					}),
					kit.Rigid(func(gtx layout.Context) layout.Dimensions {
						return subSection(th, gtx, "3 Columns", func(gtx layout.Context) layout.Dimensions {
							return kit.Grid{Cols: 3, Gap: 8}.Layout(gtx,
								gridBox(th, "Col 1 / 3", theme.Indigo100),
								gridBox(th, "Col 2 / 3", theme.Indigo200),
								gridBox(th, "Col 3 / 3", theme.Indigo300),
							)
						})
					}),
					kit.Rigid(func(gtx layout.Context) layout.Dimensions {
						return subSection(th, gtx, "4 Columns", func(gtx layout.Context) layout.Dimensions {
							return kit.Grid{Cols: 4, Gap: 8}.Layout(gtx,
								gridBox(th, "1/4", theme.Purple100),
								gridBox(th, "2/4", theme.Purple200),
								gridBox(th, "3/4", theme.Purple300),
								gridBox(th, "4/4", theme.Purple400),
							)
						})
					}),
				)
			})(gtx)
		}),

		// Flex layout examples
		kit.Rigid(func(gtx layout.Context) layout.Dimensions {
			return sectionCard(th, "Flex Layout", "FlexRow and FlexCol with gap and alignment", func(gtx layout.Context) layout.Dimensions {
				return kit.FlexCol{Gap: 20}.Layout(gtx,
					kit.Rigid(func(gtx layout.Context) layout.Dimensions {
						return subSection(th, gtx, "FlexRow — gap-8, items-center", func(gtx layout.Context) layout.Dimensions {
							return kit.FlexRow{Gap: 8, Alignment: kit.ItemsCenter}.Layout(gtx,
								kit.Rigid(func(gtx layout.Context) layout.Dimensions {
									return component.NewBadge(th, "Item A").WithVariant(component.BadgePrimary).Layout(gtx)
								}),
								kit.Rigid(func(gtx layout.Context) layout.Dimensions {
									return component.NewBadge(th, "Item B").WithVariant(component.BadgeSecondary).Layout(gtx)
								}),
								kit.Rigid(func(gtx layout.Context) layout.Dimensions {
									return component.NewBadge(th, "Item C").WithVariant(component.BadgeAccent).Layout(gtx)
								}),
								kit.Grow(func(gtx layout.Context) layout.Dimensions {
									return component.NewText(th, "← flex-1 spacer →").Sm().WithColor(theme.Gray400).Layout(gtx)
								}),
								kit.Rigid(func(gtx layout.Context) layout.Dimensions {
									return component.NewButton(th, &a.btnCard3, "End").WithVariant(component.BtnOutline).WithSize(component.BtnSm).Layout(gtx)
								}),
							)
						})
					}),
					kit.Rigid(func(gtx layout.Context) layout.Dimensions {
						return subSection(th, gtx, "FlexCol — gap-8", func(gtx layout.Context) layout.Dimensions {
							return kit.FlexCol{Gap: 8}.Layout(gtx,
								kit.Rigid(func(gtx layout.Context) layout.Dimensions {
									return component.NewAlert(th, "Row 1 — stretched full width", component.AlertInfo).Layout(gtx)
								}),
								kit.Rigid(func(gtx layout.Context) layout.Dimensions {
									return component.NewAlert(th, "Row 2 — stretched full width", component.AlertSuccess).Layout(gtx)
								}),
								kit.Rigid(func(gtx layout.Context) layout.Dimensions {
									return component.NewAlert(th, "Row 3 — stretched full width", component.AlertWarning).Layout(gtx)
								}),
							)
						})
					}),
					kit.Rigid(func(gtx layout.Context) layout.Dimensions {
						return subSection(th, gtx, "Dividers & Spacers", func(gtx layout.Context) layout.Dimensions {
							return kit.FlexCol{Gap: 12}.Layout(gtx,
								kit.Rigid(func(gtx layout.Context) layout.Dimensions {
									return component.NewText(th, "Section A").Sm().WithColor(theme.Gray500).Layout(gtx)
								}),
								kit.Rigid(func(gtx layout.Context) layout.Dimensions {
									return kit.DividerH{Color: th.Base300}.Layout(gtx)
								}),
								kit.Rigid(func(gtx layout.Context) layout.Dimensions {
									return component.NewText(th, "Section B").Sm().WithColor(theme.Gray500).Layout(gtx)
								}),
								kit.Rigid(func(gtx layout.Context) layout.Dimensions {
									return kit.DividerH{Color: th.Base300}.Layout(gtx)
								}),
								kit.Rigid(func(gtx layout.Context) layout.Dimensions {
									return component.NewText(th, "Section C").Sm().WithColor(theme.Gray500).Layout(gtx)
								}),
							)
						})
					}),
				)
			})(gtx)
		}),
	)
}

func gridBox(th *theme.Theme, label string, bg color.NRGBA) layout.Widget {
	return func(gtx layout.Context) layout.Dimensions {
		return kit.Box{
			Background: bg,
			Radius:     th.RoundedMd,
			Padding:    layout.UniformInset(th.Space3),
		}.Layout(gtx, func(gtx layout.Context) layout.Dimensions {
			return component.NewText(th, label).Sm().Layout(gtx)
		})
	}
}

// ─── Page: Forms ────────────────────────────────────────────────────────────

func (a *App) pageForms(gtx layout.Context) layout.Dimensions {
	th := a.th
	return kit.FlexCol{Gap: 32}.Layout(gtx,
		kit.Rigid(func(gtx layout.Context) layout.Dimensions {
			return pageHeader(th, gtx, "Forms", "Dashboard / Forms",
				"Input fields, toggles, and interactive controls.")
		}),

		// Inputs
		kit.Rigid(func(gtx layout.Context) layout.Dimensions {
			return sectionCard(th, "Text Inputs", "All input variants", func(gtx layout.Context) layout.Dimensions {
				return kit.Grid{Cols: 1, MdCols: 2, Gap: 20}.Layout(gtx,
					func(gtx layout.Context) layout.Dimensions {
						return kit.FlexCol{Gap: 16}.Layout(gtx,
							kit.Rigid(func(gtx layout.Context) layout.Dimensions {
								return component.NewInput(th, &a.editor1, "Default placeholder...").WithLabel("Default Input").Layout(gtx)
							}),
							kit.Rigid(func(gtx layout.Context) layout.Dimensions {
								return component.NewInput(th, &a.editor2, "Primary placeholder...").WithLabel("Primary Input").WithVariant(component.InputPrimary).Layout(gtx)
							}),
						)
					},
					func(gtx layout.Context) layout.Dimensions {
						return kit.FlexCol{Gap: 16}.Layout(gtx,
							kit.Rigid(func(gtx layout.Context) layout.Dimensions {
								return component.NewInput(th, &a.editor3, "Error state...").WithLabel("Error Input").WithVariant(component.InputError).Layout(gtx)
							}),
							kit.Rigid(func(gtx layout.Context) layout.Dimensions {
								return component.NewInput(th, &a.editor4, "Ghost style (no border)...").WithLabel("Ghost Input").WithVariant(component.InputGhost).Layout(gtx)
							}),
						)
					},
				)
			})(gtx)
		}),

		// Toggles
		kit.Rigid(func(gtx layout.Context) layout.Dimensions {
			return sectionCard(th, "Toggle Switches", "Boolean controls", func(gtx layout.Context) layout.Dimensions {
				return kit.FlexCol{Gap: 12}.Layout(gtx,
					kit.Rigid(func(gtx layout.Context) layout.Dimensions {
						return component.NewToggle(th, &a.toggle1, "Enable notifications (default on)").Layout(gtx)
					}),
					kit.Rigid(func(gtx layout.Context) layout.Dimensions {
						return component.NewToggle(th, &a.toggle2, "Dark mode").Layout(gtx)
					}),
					kit.Rigid(func(gtx layout.Context) layout.Dimensions {
						return component.NewToggle(th, &a.toggle3, "Auto-save drafts").Layout(gtx)
					}),
					kit.Rigid(func(gtx layout.Context) layout.Dimensions {
						state := "off"
						if a.toggle1.Value {
							state = "on"
						}
						return component.NewText(th, fmt.Sprintf("Notifications: %s", state)).Sm().WithColor(theme.Gray500).Layout(gtx)
					}),
				)
			})(gtx)
		}),

		// Overlays demo
		kit.Rigid(func(gtx layout.Context) layout.Dimensions {
			return sectionCard(th, "Overlays", "Modal, Drawer, and Toast", func(gtx layout.Context) layout.Dimensions {
				return kit.FlexCol{Gap: 16}.Layout(gtx,
					kit.Rigid(func(gtx layout.Context) layout.Dimensions {
						return component.NewText(th, "Modal").H4().Layout(gtx)
					}),
					kit.Rigid(func(gtx layout.Context) layout.Dimensions {
						return component.NewText(th, "Click backdrop or the Close button to dismiss.").Sm().WithColor(theme.Gray500).Layout(gtx)
					}),
					kit.Rigid(func(gtx layout.Context) layout.Dimensions {
						return component.NewButton(th, &a.btnModal, "Open Modal").WithVariant(component.BtnPrimary).Layout(gtx)
					}),
					kit.Rigid(func(gtx layout.Context) layout.Dimensions {
						return kit.DividerH{Color: th.Base300}.Layout(gtx)
					}),
					kit.Rigid(func(gtx layout.Context) layout.Dimensions {
						return component.NewText(th, "Drawer").H4().Layout(gtx)
					}),
					kit.Rigid(func(gtx layout.Context) layout.Dimensions {
						return component.NewText(th, "A slide-in panel from the left edge.").Sm().WithColor(theme.Gray500).Layout(gtx)
					}),
					kit.Rigid(func(gtx layout.Context) layout.Dimensions {
						return component.NewButton(th, &a.btnDrawer, "Open Drawer").WithVariant(component.BtnSecondary).Layout(gtx)
					}),
					kit.Rigid(func(gtx layout.Context) layout.Dimensions {
						return kit.DividerH{Color: th.Base300}.Layout(gtx)
					}),
					kit.Rigid(func(gtx layout.Context) layout.Dimensions {
						return component.NewText(th, "Toast").H4().Layout(gtx)
					}),
					kit.Rigid(func(gtx layout.Context) layout.Dimensions {
						return component.NewText(th, "Auto-dismisses after 3 seconds.").Sm().WithColor(theme.Gray500).Layout(gtx)
					}),
					kit.Rigid(func(gtx layout.Context) layout.Dimensions {
						return component.NewButton(th, &a.btnToast, "Show Toast").WithVariant(component.BtnSuccess).Layout(gtx)
					}),
				)
			})(gtx)
		}),
	)
}

// ─── Page: Settings ─────────────────────────────────────────────────────────

func (a *App) pageSettings(gtx layout.Context) layout.Dimensions {
	th := a.th
	return kit.FlexCol{Gap: 32}.Layout(gtx,
		kit.Rigid(func(gtx layout.Context) layout.Dimensions {
			return pageHeader(th, gtx, "Settings", "Dashboard / Settings",
				"Customize the application theme and appearance.")
		}),

		// Theme picker
		kit.Rigid(func(gtx layout.Context) layout.Dimensions {
			return sectionCard(th, "Theme", "Switch between built-in themes", func(gtx layout.Context) layout.Dimensions {
				return kit.FlexCol{Gap: 16}.Layout(gtx,
					kit.Rigid(func(gtx layout.Context) layout.Dimensions {
						return component.NewText(th, "Select a theme to apply it globally.").Sm().WithColor(theme.Gray500).Layout(gtx)
					}),
					kit.Rigid(func(gtx layout.Context) layout.Dimensions {
						return kit.FlexRow{Gap: 8, Alignment: kit.ItemsCenter}.Layout(gtx,
							kit.Rigid(func(gtx layout.Context) layout.Dimensions {
								return component.NewButton(th, &a.btnLight, "Light").WithVariant(component.BtnOutline).Layout(gtx)
							}),
							kit.Rigid(func(gtx layout.Context) layout.Dimensions {
								return component.NewButton(th, &a.btnDark, "Dark").WithVariant(component.BtnOutline).Layout(gtx)
							}),
							kit.Rigid(func(gtx layout.Context) layout.Dimensions {
								return component.NewButton(th, &a.btnCupcake, "Cupcake").WithVariant(component.BtnOutline).Layout(gtx)
							}),
							kit.Rigid(func(gtx layout.Context) layout.Dimensions {
								return component.NewButton(th, &a.btnNord, "Nord").WithVariant(component.BtnOutline).Layout(gtx)
							}),
						)
					}),
				)
			})(gtx)
		}),

		// Typography preview
		kit.Rigid(func(gtx layout.Context) layout.Dimensions {
			return sectionCard(th, "Typography", "Font size scale", func(gtx layout.Context) layout.Dimensions {
				return kit.FlexCol{Gap: 8}.Layout(gtx,
					kit.Rigid(func(gtx layout.Context) layout.Dimensions {
						return component.NewText(th, "Heading 1 — H1").H1().Layout(gtx)
					}),
					kit.Rigid(func(gtx layout.Context) layout.Dimensions {
						return component.NewText(th, "Heading 2 — H2").H2().Layout(gtx)
					}),
					kit.Rigid(func(gtx layout.Context) layout.Dimensions {
						return component.NewText(th, "Heading 3 — H3").H3().Layout(gtx)
					}),
					kit.Rigid(func(gtx layout.Context) layout.Dimensions {
						return component.NewText(th, "Heading 4 — H4").H4().Layout(gtx)
					}),
					kit.Rigid(func(gtx layout.Context) layout.Dimensions {
						return component.NewText(th, "Body — default font size").Layout(gtx)
					}),
					kit.Rigid(func(gtx layout.Context) layout.Dimensions {
						return component.NewText(th, "Small — sm size, muted").Sm().WithColor(theme.Gray500).Layout(gtx)
					}),
					kit.Rigid(func(gtx layout.Context) layout.Dimensions {
						return component.NewText(th, "XSmall — xs size, muted").Xs().WithColor(theme.Gray400).Layout(gtx)
					}),
				)
			})(gtx)
		}),

		// Color palette preview
		kit.Rigid(func(gtx layout.Context) layout.Dimensions {
			return sectionCard(th, "Color Palette", "Current theme semantic colors", func(gtx layout.Context) layout.Dimensions {
				return kit.Grid{Cols: 2, MdCols: 4, Gap: 12}.Layout(gtx,
					colorSwatch(th, "Primary", th.Primary, th.PrimaryContent),
					colorSwatch(th, "Secondary", th.Secondary, th.SecondaryContent),
					colorSwatch(th, "Accent", th.Accent, th.AccentContent),
					colorSwatch(th, "Neutral", th.Neutral, th.NeutralContent),
					colorSwatch(th, "Info", th.Info, th.InfoContent),
					colorSwatch(th, "Success", th.Success, th.SuccessContent),
					colorSwatch(th, "Warning", th.Warning, th.WarningContent),
					colorSwatch(th, "Error", th.Error, th.ErrorContent),
				)
			})(gtx)
		}),
	)
}

// ─── Helpers ────────────────────────────────────────────────────────────────

func pageHeader(th *theme.Theme, gtx layout.Context, title, breadcrumb, subtitle string) layout.Dimensions {
	return kit.FlexCol{Gap: 8}.Layout(gtx,
		kit.Rigid(func(gtx layout.Context) layout.Dimensions {
			parts := []string{"Home"}
			if breadcrumb != "" {
				parts = append(parts, title)
			}
			return scaffold.NewBreadcrumb(th, parts...).Layout(gtx)
		}),
		kit.Rigid(func(gtx layout.Context) layout.Dimensions {
			return component.NewText(th, title).H1().Layout(gtx)
		}),
		kit.Rigid(func(gtx layout.Context) layout.Dimensions {
			return component.NewText(th, subtitle).Sm().WithColor(theme.Gray500).Layout(gtx)
		}),
	)
}

func subSection(th *theme.Theme, gtx layout.Context, label string, body layout.Widget) layout.Dimensions {
	return kit.FlexCol{Gap: 8}.Layout(gtx,
		kit.Rigid(func(gtx layout.Context) layout.Dimensions {
			return component.NewText(th, label).H4().Layout(gtx)
		}),
		kit.Rigid(body),
	)
}

func sectionCard(th *theme.Theme, title, subtitle string, body layout.Widget) layout.Widget {
	return func(gtx layout.Context) layout.Dimensions {
		return kit.FlexCol{Gap: 12}.Layout(gtx,
			kit.Rigid(func(gtx layout.Context) layout.Dimensions {
				return kit.FlexCol{Gap: 4}.Layout(gtx,
					kit.Rigid(func(gtx layout.Context) layout.Dimensions {
						return component.NewText(th, title).H2().Layout(gtx)
					}),
					kit.Rigid(func(gtx layout.Context) layout.Dimensions {
						return component.NewText(th, subtitle).Sm().WithColor(theme.Gray400).Layout(gtx)
					}),
				)
			}),
			kit.Rigid(func(gtx layout.Context) layout.Dimensions {
				return component.NewCard(th).WithBorder().Layout(gtx, body)
			}),
		)
	}
}

func colorSwatch(th *theme.Theme, name string, bg, fg color.NRGBA) layout.Widget {
	return func(gtx layout.Context) layout.Dimensions {
		card := component.NewCard(th)
		return card.Layout(gtx, func(gtx layout.Context) layout.Dimensions {
			return kit.Box{
				Background: bg,
				Radius:     th.RoundedMd,
				Padding:    layout.Inset{Top: th.Space3, Bottom: th.Space3, Left: th.Space3, Right: th.Space3},
			}.Layout(gtx, func(gtx layout.Context) layout.Dimensions {
				return component.NewText(th, name).Sm().WithColor(fg).Layout(gtx)
			})
		})
	}
}
