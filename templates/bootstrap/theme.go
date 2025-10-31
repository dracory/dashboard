package bootstrap

import (
	"github.com/samber/lo"
	"strconv"
	"strings"
)

// isThemeDark checks if the given theme is a dark theme
func isThemeDark(theme string) bool {
	_, isDark := themesDark[theme]
	return isDark
}

// navbarHasBackgroundThemeClass determines if the navbar should use a background theme class
func navbarHasBackgroundThemeClass(navbarBackgroundColor, navbarBackgroundColorMode string) bool {
	hasNavbarBackgroundColor := navbarBackgroundColor != ""
	return !hasNavbarBackgroundColor && navbarBackgroundColorMode != ""
}

// navbarBackgroundThemeClass returns the appropriate background theme class for the navbar
func navbarBackgroundThemeClass(navbarBackgroundColor, navbarBackgroundColorMode string) string {
	return lo.Ternary(
		navbarHasBackgroundThemeClass(navbarBackgroundColor, navbarBackgroundColorMode),
		"bg-"+navbarBackgroundColorMode,
		"",
	)
}

// navbarButtonThemeClass returns the appropriate button theme class for navbar buttons
func navbarButtonThemeClass(navbarBackgroundColor, navbarBackgroundColorMode string) string {
	return lo.Ternary(
		navbarHasBackgroundThemeClass(navbarBackgroundColor, navbarBackgroundColorMode),
		"btn-outline-"+navbarBackgroundColorMode,
		"",
	)
}

func defaultNavbarTextColor(navbarBackgroundColor, navbarBackgroundColorMode string) string {
	if navbarBackgroundColorMode != "" {
		switch navbarBackgroundColorMode {
		case "light", "warning", "info", "white":
			return "#212529"
		default:
			return "#ffffff"
		}
	}

	if navbarBackgroundColor != "" {
		if r, g, b, ok := hexToRGB(navbarBackgroundColor); ok {
			brightness := 0.299*float64(r) + 0.587*float64(g) + 0.114*float64(b)
			if brightness > 186 {
				return "#212529"
			}
			return "#ffffff"
		}
	}

	return ""
}

func hexToRGB(hex string) (int, int, int, bool) {
	cleaned := strings.TrimPrefix(hex, "#")
	if len(cleaned) != 6 {
		return 0, 0, 0, false
	}

	value, err := strconv.ParseUint(cleaned, 16, 32)
	if err != nil {
		return 0, 0, 0, false
	}

	r := int((value >> 16) & 0xFF)
	g := int((value >> 8) & 0xFF)
	b := int(value & 0xFF)

	return r, g, b, true
}
