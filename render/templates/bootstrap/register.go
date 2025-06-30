// Package bootstrap provides Bootstrap template implementation for the dashboard
package bootstrap

import (
	"github.com/dracory/dashboard/model/templateregistry"
)

// init registers the Bootstrap template with the template registry
func init() {
	templateregistry.Register(New())
}
