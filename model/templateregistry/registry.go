package templateregistry

import "github.com/dracory/dashboard/model/interfaces"

// TemplateRegistry defines the interface for template registration
type TemplateRegistry interface {
	// Register registers a template with the registry
	Register(template interfaces.Template)
	// Get returns a template by name, or the default template if not found
	Get(name string) interfaces.Template
	// GetDefault returns the default template
	GetDefault() interfaces.Template
	// GetTemplates returns all registered templates
	GetTemplates() map[string]interfaces.Template
}

// registry is the global template registry
var registry TemplateRegistry

// SetRegistry sets the global template registry
func SetRegistry(r TemplateRegistry) {
	registry = r
}

// Register registers a template with the global registry
func Register(template interfaces.Template) {
	if registry != nil {
		registry.Register(template)
	}
}

// Get returns a template by name from the global registry
func Get(name string) interfaces.Template {
	if registry != nil {
		return registry.Get(name)
	}
	return nil
}

// GetDefault returns the default template from the global registry
func GetDefault() interfaces.Template {
	if registry != nil {
		return registry.GetDefault()
	}
	return nil
}

// GetTemplates returns all registered templates from the global registry
func GetTemplates() map[string]interfaces.Template {
	if registry != nil {
		return registry.GetTemplates()
	}
	return nil
}
