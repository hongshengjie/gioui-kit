# CLAUDE.md

This file provides guidance to Claude Code (claude.ai/code) when working with code in this repository.

## Commands

```bash
# Run the example/demo app
cd example && go run .

# Build all packages
go build ./...

# Vet all packages
go vet ./...
```

## Architecture

GioUI Kit is a **TailwindCSS + DaisyUI inspired** component library for [Gio UI](https://gioui.org). The packages form a dependency hierarchy — lower layers have no knowledge of higher ones:

```
theme → layout → modifier → component → scaffold
```

### Package responsibilities

- **`theme/`** — Foundation layer. Defines `Theme` struct with DaisyUI semantic color tokens (`Primary`, `Base100`, etc.), Tailwind color palette constants (`Slate50`…`Slate950`, `Blue500`, etc.), typography scale (`H1Size`…`XsSize`), spacing scale (`Space0`…`Space16`), and border-radius constants. Every component accepts `*theme.Theme`.

- **`layout/`** — Flexbox/grid/spacing utilities mirroring Tailwind class names. `FlexRow`/`FlexCol` wrap Gio's `layout.Flex` with gap support. `Grid` provides n-column layouts. `Box`/`Container` provide padded wrappers. Size helpers (`W`, `H`, `MinW`, `MaxW`, `WFull`) return `func(gtx, widget)` closures. Spacing helpers (`P`, `Px`, `Py`, `Pt`, etc.) return `layout.Inset`. `ScrollY`/`ScrollX` wrap `widget.List`.

- **`modifier/`** — Visual decorator layer. Each type (`Shadow`, `Bg`, `Rounded`, `Ring`, `LinearGradient`, `Opacity`) implements `.Layout(gtx, widget)` by painting to the clip/paint ops stack before delegating to the child widget.

- **`component/`** — DaisyUI-style stateful components: `Button`, `Badge`, `Card`, `Alert`, `Input`, `Toggle`, `Avatar`, `Progress`, `Tabs`, `Chip`, `Skeleton`, `Text`. All follow the builder pattern: `NewXxx(th, ...).WithVariant(...).WithSize(...).Layout(gtx)`.

- **`scaffold/`** — App-level layout primitives: `AppShell` (navbar + sidebar + scrollable content), `Navbar`, `Sidebar` (with `OnSelect` callback), `Modal`, `Drawer` (left/right), `BottomNav`, `Toast`, `Breadcrumb`.

### Key patterns

- **Theme is passed explicitly** — every component constructor takes `*theme.Theme` as first argument; there is no global theme.
- **Immediate-mode rendering** — all `.Layout(gtx, ...)` calls must happen every frame; components hold widget state (e.g. `widget.Clickable`) separately from rendering logic.
- **Builder pattern for components** — use `With*` methods before calling `.Layout(gtx)`.
- **Gio layout children** — `kit.Rigid(fn)`, `kit.Grow(fn)`, `kit.Flexed(weight, fn)` create `layout.FlexChild` values for use with `FlexRow`/`FlexCol`.
- **Responsive / breakpoints** — `kit.ScreenBreakpoint(gtx)` returns `BreakpointXs/Sm/Md/Lg/Xl/2xl` based on available width (Tailwind thresholds: 640/768/1024/1280/1536 dp). `Grid` accepts `SmCols`, `MdCols`, `LgCols`, `XlCols` for per-breakpoint column counts. `AppShell.HideSidebarBelow` (default `BreakpointLg`) hides the inline sidebar on narrow screens.

### Module

`github.com/hongshengjie/gioui-kit`, requires `gioui.org v0.9.0`.
