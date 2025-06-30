package templates

import "github.com/gouniverse/hb"

// Template defines the interface for dashboard themes
type Template interface {
	// GetName returns the theme's name
	GetName() string

	// GetCSSLinks returns the CSS link tags for the theme
	GetCSSLinks(isDarkMode bool) []*hb.Tag

	// GetJSScripts returns the JavaScript script tags for the theme
	GetJSScripts() []*hb.Tag

	// GetCustomCSS returns any custom CSS for the theme
	GetCustomCSS() string

	// GetCustomJS returns any custom JavaScript for the theme
	GetCustomJS() string
}
