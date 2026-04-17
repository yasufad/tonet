package ui

import (
	"image/color"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/theme"
)

// customTheme defines the visual theme for tonet
type customTheme struct{}

var _ fyne.Theme = (*customTheme)(nil)

func (m customTheme) Color(name fyne.ThemeColorName, variant fyne.ThemeVariant) color.Color {
	if name == theme.ColorNameBackground {
		return color.NRGBA{R: 15, G: 17, B: 23, A: 255} // #0f1117
	}
	if name == theme.ColorNamePrimary {
		return color.NRGBA{R: 0, G: 201, B: 167, A: 255} // #00c9a7
	}
	if name == theme.ColorNameForeground {
		return color.NRGBA{R: 232, G: 234, B: 240, A: 255} // #e8eaf0
	}
	
	// Fallback to default theme colors for everything else
	return theme.DefaultTheme().Color(name, theme.VariantDark)
}

func (m customTheme) Font(style fyne.TextStyle) fyne.Resource {
	return theme.DefaultTheme().Font(style)
}

func (m customTheme) Icon(name fyne.ThemeIconName) fyne.Resource {
	return theme.DefaultTheme().Icon(name)
}

func (m customTheme) Size(name fyne.ThemeSizeName) float32 {
	// Slightly larger padding for breathing room
	if name == theme.SizeNamePadding {
		return 8
	}
	return theme.DefaultTheme().Size(name)
}
