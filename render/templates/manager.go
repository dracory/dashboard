package templates

import (
	"strings"
	"sync"

	"github.com/dracory/dashboard/config"
	"github.com/dracory/dashboard/render/templates/adminlte"
	"github.com/dracory/dashboard/render/templates/bootstrap"
	"github.com/dracory/dashboard/render/templates/shared"
	"github.com/dracory/dashboard/render/templates/tabler"
)

// manager handles template management
type manager struct {
	templates   map[string]shared.Template
	defaultTmpl shared.Template
}

var instance *manager
var once sync.Once

// Manager returns the singleton instance of the template manager
func Manager() *manager {
	once.Do(func() {
		instance = &manager{
			templates: make(map[string]shared.Template),
		}
		// Initialize all known templates at startup
		instance.initializeTemplates()
	})
	return instance
}

// initializeTemplates initializes all available templates
func (m *manager) initializeTemplates() {
	// Initialize all templates
	m.templates[config.TEMPLATE_ADMINLTE] = adminlte.NewAdminLTETemplate()
	m.templates[config.TEMPLATE_BOOTSTRAP] = bootstrap.NewBootstrapTemplate()
	m.templates[config.TEMPLATE_TABLER] = tabler.NewTablerTemplate()

	// Set default template
	m.defaultTmpl = m.templates[config.TEMPLATE_DEFAULT]
}

// Get returns a template by name, or the default template if not found
func (m *manager) Get(name string) shared.Template {
	if name == "" || name == "light" || name == "dark" {
		return m.defaultTmpl
	}

	// Try exact match first
	if tmpl, exists := m.templates[name]; exists {
		return tmpl
	}

	// Fall back to case-insensitive match
	name = strings.ToLower(name)
	for key, tmpl := range m.templates {
		if strings.ToLower(key) == name {
			return tmpl
		}
	}

	return m.defaultTmpl
}

// GetDefault returns the default template
func (m *manager) GetDefault() shared.Template {
	return m.defaultTmpl
}

// GetTemplates returns all registered templates
func (m *manager) GetTemplates() map[string]shared.Template {
	return m.templates
}

// Initialize is kept for backward compatibility
// Use InitializeTemplates() instead to avoid import cycles
func Initialize() {
	// This function is now a no-op
	// Templates should be initialized using InitializeTemplates()
}
