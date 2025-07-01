package bootstrap

import (
	"github.com/samber/lo"
)

// isThemeDark checks if the given theme is a dark theme
func isThemeDark(theme string) bool {
	_, isDark := ThemesDark[theme]
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
