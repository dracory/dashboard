package shared

type ThemeNameContextKey struct{}

// Template constants
const TEMPLATE_ADMINLTE = "adminlte"
const TEMPLATE_BOOTSTRAP = "bootstrap"
const TEMPLATE_TABLER = "tabler"
const TEMPLATE_DEFAULT = TEMPLATE_BOOTSTRAP

// Menu type constants
// const MENU_TYPE_MODAL = "modal"
// const MENU_TYPE_OFFCANVAS = "offcanvas"

// MenuType constants for dashboard menu types
const (
	TEMPLATE_BOOTSTRAP_MENU_TYPE_MODAL     = "modal"
	TEMPLATE_BOOTSTRAP_MENU_TYPE_OFFCANVAS = "offcanvas"
	THEME_COOKIE_KEY                       = "theme"
)
