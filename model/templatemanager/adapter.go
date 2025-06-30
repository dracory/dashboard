package templatemanager

import (
	"github.com/dracory/dashboard/model/interfaces"
	"github.com/dracory/dashboard/model/templateregistry"
)

// registryAdapter adapts Manager to implement templateregistry.TemplateRegistry
type registryAdapter struct {
	*Manager
}

// NewRegistryAdapter creates a new adapter that makes Manager implement templateregistry.TemplateRegistry
func NewRegistryAdapter(manager *Manager) templateregistry.TemplateRegistry {
	return &registryAdapter{Manager: manager}
}

// Register registers a template with the registry
func (a *registryAdapter) Register(template interfaces.Template) {
	a.Manager.Register(template)
}

// Get returns a template by name, or the default template if not found
func (a *registryAdapter) Get(name string) interfaces.Template {
	return a.Manager.Get(name)
}

// GetDefault returns the default template
func (a *registryAdapter) GetDefault() interfaces.Template {
	return a.Manager.GetDefault()
}

// GetTemplates returns all registered templates
func (a *registryAdapter) GetTemplates() map[string]interfaces.Template {
	return a.Manager.GetTemplates()
}
