package theme

import (
	"fmt"
	"sync"

	"github.com/dracory/dashboard/render/theme/shared"
)

var (
	themes      = make(map[string]shared.Theme)
	defaultTheme shared.Theme
	once        sync.Once
	instance   *manager
)

// manager handles theme registration and retrieval
type manager struct {
	themes       map[string]shared.Theme
	defaultTheme shared.Theme
}

// Manager returns the singleton instance of the theme manager
func Manager() *manager {
	once.Do(func() {
		instance = &manager{
			themes:       make(map[string]shared.Theme),
			defaultTheme: nil,
		}
	})
	return instance
}

// Register adds a theme to the manager
func (m *manager) Register(theme shared.Theme) {
	name := theme.GetName()
	fmt.Printf("[DEBUG] Registering theme: %s\n", name)
	m.themes[name] = theme
	if m.defaultTheme == nil {
		m.defaultTheme = theme
		fmt.Printf("[DEBUG] Set as default theme: %s\n", name)
	}
}

// Get returns a theme by name, or the default theme if not found
func (m *manager) Get(name string) shared.Theme {
	fmt.Printf("[DEBUG] Requested theme: %s\n", name)
	if theme, exists := m.themes[name]; exists {
		fmt.Printf("[DEBUG] Found theme: %s\n", name)
		return theme
	}
	fmt.Printf("[WARN] Theme not found: %s, using default: %s\n", name, m.defaultTheme.GetName())
	return m.defaultTheme
}

// SetDefault sets the default theme
func (m *manager) SetDefault(theme shared.Theme) {
	m.defaultTheme = theme
}

// GetDefault returns the default theme
func (m *manager) GetDefault() shared.Theme {
	return m.defaultTheme
}

// GetThemes returns all registered themes
func (m *manager) GetThemes() map[string]shared.Theme {
	return m.themes
}

// Initialize is kept for backward compatibility
// Use InitializeThemes() instead to avoid import cycles
func Initialize() {
	// This function is now a no-op
	// Themes should be initialized using InitializeThemes()
}
