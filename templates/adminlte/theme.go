package adminlte

// Theme constants
const (
	// ThemeDefault is the default theme
	ThemeDefault = "default"
	// ThemeDark is the dark theme
	ThemeDark = "dark"
	// ThemeLight is the light theme
	ThemeLight = "light"
)

// ThemeNames returns a map of theme names to their display names
func ThemeNames() map[string]string {
	return map[string]string{
		ThemeDefault: "Default",
		ThemeDark:    "Dark",
		ThemeLight:   "Light",
	}
}

// ThemeColors returns a map of theme names to their color codes
func ThemeColors() map[string]string {
	return map[string]string{
		ThemeDefault: "#3c8dbc",
		ThemeDark:    "#222d32",
		ThemeLight:   "#f4f6f9",
	}
}

// ThemeIcons returns a map of theme names to their icon classes
func ThemeIcons() map[string]string {
	return map[string]string{
		ThemeDefault: "fas fa-adjust",
		ThemeDark:    "fas fa-moon",
		ThemeLight:   "fas fa-sun",
	}
}

// ThemeHandler handles theme switching
func ThemeHandler(theme string) string {
	switch theme {
	case ThemeDark:
		return ThemeDark
	case ThemeLight:
		return ThemeLight
	default:
		return ThemeDefault
	}
}
