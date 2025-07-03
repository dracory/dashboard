package dashboard

import (
	"go/types"

	"github.com/dracory/dashboard/shared"
)

// ============================================================================
// Template constants
// - These are the same as the constants in shared package, but are redeclared
//   here for the ease of user
// ============================================================================

const TEMPLATE_DEFAULT = shared.TEMPLATE_DEFAULT
const TEMPLATE_ADMINLTE = shared.TEMPLATE_ADMINLTE
const TEMPLATE_TABLER = shared.TEMPLATE_TABLER
const TEMPLATE_BOOTSTRAP = shared.TEMPLATE_BOOTSTRAP

type Config struct {
	types.Config
}

// type BootstrapConfig struct {
// 	types.BootstrapConfig
// }

// type AdminlteConfig struct {
// 	types.AdminlteConfig
// }

// type TablerConfig struct {
// 	types.TablerConfig
// }
