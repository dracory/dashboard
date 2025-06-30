// Package tabler provides Tabler template implementation for the dashboard
package tabler

import (
	"github.com/dracory/dashboard/model/templateregistry"
)

// init registers the Tabler template with the template registry
func init() {
	templateregistry.Register(New())
}
