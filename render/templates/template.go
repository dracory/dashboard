package templates

import (
	"github.com/dracory/omni"
	"github.com/gouniverse/hb"
)

// Template defines the interface for dashboard templates
type Template interface {
	// RenderPage renders a complete page with the given content and dashboard renderer
	RenderPage(content string, d DashboardRenderer) (*hb.Tag, error)

	// GetName returns the template's name
	GetName() string

	// GetCSSLinks returns the CSS link tags for the template
	GetCSSLinks(isDarkMode bool) []*hb.Tag

	// GetJSScripts returns the JavaScript script tags for the template
	GetJSScripts() []*hb.Tag

	// GetCustomCSS returns any custom CSS for the template
	GetCustomCSS() string

	// GetCustomJS returns any custom JavaScript for the template
	GetCustomJS() string

	// RenderHeader renders the template-specific header
	RenderHeader(d DashboardRenderer) *hb.Tag

	// RenderFooter renders the template-specific footer
	RenderFooter(d DashboardRenderer) *hb.Tag

	// RenderAtom renders an Omni atom using the template's styling
	RenderAtom(a *omni.Atom) (*hb.Tag, error)

	// RenderDashboard renders a dashboard using the template's layout
	RenderDashboard(d DashboardRenderer) (*hb.Tag, error)
}
