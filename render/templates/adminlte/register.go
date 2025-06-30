// Package adminlte provides AdminLTE template implementation for the dashboard
package adminlte

import (
	"github.com/dracory/dashboard/model/templateregistry"
)

// init registers the AdminLTE template with the template registry
func init() {
	templateregistry.Register(NewAdminLTETemplate())
}
