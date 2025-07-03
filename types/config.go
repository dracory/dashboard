package types

// Config represents the configuration for the dashboard
type Config struct {
	User                 User
	MenuMainItems        []MenuItem
	MenuUserItems        []MenuItem
	MenuQuickAccessItems []MenuItem
}

// BootstrapConfig represents the specific configuration for the bootstrap template
type BootstrapConfig struct {
	NavbarBackgroundColorMode string
	NavbarBackgroundColor     string
	NavbarTextColor           string
}

// AdminLTEConfig represents the specific configuration for the adminlte template
type AdminLTEConfig struct {
}

// TablerConfig represents the specific configuration for the tabler template
type TablerConfig struct {
	Layout string
}
