// Package theme provides theme management for the dashboard
package theme

import "github.com/dracory/dashboard/render/theme/shared"

// ThemeInitializer is a function that initializes a theme
type ThemeInitializer func() shared.Theme

var themeInitializers []ThemeInitializer

// RegisterThemeInitializer registers a theme initializer
func RegisterThemeInitializer(initializer ThemeInitializer) {
	themeInitializers = append(themeInitializers, initializer)
}

// InitializeRegisteredThemes initializes all registered themes
func InitializeRegisteredThemes() {
	m := Manager()
	for _, initializer := range themeInitializers {
		m.Register(initializer())
	}
}
