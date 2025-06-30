package templates

import (
	"fmt"
	"sync"

	"github.com/dracory/dashboard/render/templates/shared"
)

var (
	templates    = make(map[string]shared.Template)
	defaultTmpl  shared.Template
	once         sync.Once
	instance     *manager
)

// manager handles template registration and retrieval
type manager struct {
	templates    map[string]shared.Template
	defaultTmpl  shared.Template
}

// Manager returns the singleton instance of the template manager
func Manager() *manager {
	once.Do(func() {
		instance = &manager{
			templates:    make(map[string]shared.Template),
			defaultTmpl:  nil,
		}
	})
	return instance
}

// Register adds a template to the manager
func (m *manager) Register(tmpl shared.Template) {
	name := tmpl.GetName()
	fmt.Printf("[DEBUG] Registering template: %s\n", name)
	m.templates[name] = tmpl
	if m.defaultTmpl == nil {
		m.defaultTmpl = tmpl
		fmt.Printf("[DEBUG] Set as default template: %s\n", name)
	}
}

// Get returns a template by name, or the default template if not found
func (m *manager) Get(name string) shared.Template {
	fmt.Printf("[DEBUG] Requested template: %s\n", name)
	if tmpl, exists := m.templates[name]; exists {
		fmt.Printf("[DEBUG] Found template: %s\n", name)
		return tmpl
	}
	fmt.Printf("[WARN] Template not found: %s, using default: %s\n", name, m.defaultTmpl.GetName())
	return m.defaultTmpl
}

// SetDefault sets the default template
func (m *manager) SetDefault(tmpl shared.Template) {
	m.defaultTmpl = tmpl
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
