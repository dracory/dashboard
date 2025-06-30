package templatemanager

import (
	"sync"

	"github.com/dracory/dashboard/model/interfaces"
)

// Manager implements the templateregistry.TemplateRegistry interface
type Manager struct {
	templates   map[string]interfaces.Template
	defaultTmpl interfaces.Template
	mu         sync.RWMutex
}

var instance *Manager
var once sync.Once

// NewManager creates a new template manager
func NewManager() *Manager {
	return &Manager{
		templates: make(map[string]interfaces.Template),
	}
}

// Initialize initializes the template manager with the default templates
// This is a no-op now as templates register themselves
func (m *Manager) Initialize() {}

// Register registers a template with the manager
func (m *Manager) Register(template interfaces.Template) {
	if template == nil {
		return
	}

	m.mu.Lock()
	defer m.mu.Unlock()

	m.templates[template.GetName()] = template

	// Set as default if this is the first template
	if len(m.templates) == 1 {
		m.defaultTmpl = template
	}
}

// Get returns a template by name, or the default template if not found
func (m *Manager) Get(name string) interfaces.Template {
	m.mu.RLock()
	defer m.mu.RUnlock()

	if name == "" {
		return m.defaultTmpl
	}

	tmpl, exists := m.templates[name]
	if !exists {
		return m.defaultTmpl
	}

	return tmpl
}

// GetDefault returns the default template
func (m *Manager) GetDefault() interfaces.Template {
	m.mu.RLock()
	defer m.mu.RUnlock()
	return m.defaultTmpl
}

// GetTemplates returns a copy of all registered templates
func (m *Manager) GetTemplates() map[string]interfaces.Template {
	m.mu.RLock()
	defer m.mu.RUnlock()

	templates := make(map[string]interfaces.Template, len(m.templates))
	for k, v := range m.templates {
		templates[k] = v
	}
	return templates
}
