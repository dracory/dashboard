// Package templates provides template management for the dashboard
package templates

import (
	"fmt"
	"github.com/dracory/dashboard/render/templates/shared"
)

// TemplateInitializer is a function that initializes a template
type TemplateInitializer func() shared.Template

var templateInitializers []TemplateInitializer

// RegisterTemplateInitializer registers a template initializer
func RegisterTemplateInitializer(initializer TemplateInitializer) {
	templateInitializers = append(templateInitializers, initializer)
}

// InitializeRegisteredTemplates initializes all registered templates
func InitializeRegisteredTemplates() {
	m := Manager()
	for _, initializer := range templateInitializers {
		tmpl := initializer()
		m.Register(tmpl)
		fmt.Printf("[DEBUG] Initialized template: %s\n", tmpl.GetName())
	}
}

// InitializeTemplates is an alias for InitializeRegisteredTemplates for backward compatibility
// and to maintain a clean API
func InitializeTemplates() {
	InitializeRegisteredTemplates()
}
