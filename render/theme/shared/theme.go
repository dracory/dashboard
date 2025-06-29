package shared

import (
	"github.com/dracory/dashboard/model"
	"github.com/gouniverse/hb"
)

// DashboardRenderer defines the interface for dashboard renderers
type DashboardRenderer = model.DashboardRenderer

// Theme defines the interface for dashboard themes
type Theme interface {
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

	// RenderHeader renders the theme-specific header
	RenderHeader(d DashboardRenderer) *hb.Tag

	// RenderFooter renders the theme-specific footer
	RenderFooter(d DashboardRenderer) *hb.Tag
}

// DefaultTheme is a basic theme implementation that can be used as a fallback
type DefaultTheme struct{}

// GetName returns the name of the default theme
func (t *DefaultTheme) GetName() string {
	return "default"
}

// GetCSSLinks returns no CSS links for the default theme
func (t *DefaultTheme) GetCSSLinks(isDarkMode bool) []*hb.Tag {
	// Return basic CSS reset as a fallback
	return []*hb.Tag{
		hb.Style(`
			* { box-sizing: border-box; margin: 0; padding: 0; }
			body { font-family: -apple-system, BlinkMacSystemFont, "Segoe UI", Roboto, "Helvetica Neue", Arial, sans-serif; line-height: 1.6; }
		`),
	}
}

// GetJSScripts returns no JavaScript for the default theme
func (t *DefaultTheme) GetJSScripts() []*hb.Tag {
	return nil
}

// GetCustomCSS returns empty string as there's no custom CSS for the default theme
func (t *DefaultTheme) GetCustomCSS() string {
	return ""
}

// GetCustomJS returns empty string as there's no custom JavaScript for the default theme
func (t *DefaultTheme) GetCustomJS() string {
	return ""
}

// RenderHeader renders a basic header for the default theme
func (t *DefaultTheme) RenderHeader(d DashboardRenderer) *hb.Tag {
	header := hb.NewHeader().Class("p-3 bg-light border-bottom")
	container := hb.NewDiv().Class("container-fluid")
	headerInner := hb.NewDiv().Class("d-flex flex-wrap align-items-center justify-content-center justify-content-lg-start")
	logoLink := hb.NewA().Href("/").Class("d-flex align-items-center mb-2 mb-lg-0 text-dark text-decoration-none")
	logoLink.AddChild(hb.NewSpan().Text("Dashboard").Class("fs-4"))
	headerInner.AddChild(logoLink)
	container.AddChild(headerInner)
	header.AddChild(container)
	return header
}

// RenderFooter renders a basic footer for the default theme
func (t *DefaultTheme) RenderFooter(d DashboardRenderer) *hb.Tag {
	leftCol := hb.NewDiv().
		Class("col-12 col-md-6").
		AddChild(hb.NewTag("small").Text("© 2025 Dashboard").Class("text-muted"))

	rightCol := hb.NewDiv().
		Class("col-12 col-md-6 text-end").
		AddChild(hb.NewTag("small").Text("Powered by Dashboard").Class("text-muted"))

	row := hb.NewDiv().
		Class("row").
		AddChild(leftCol).
		AddChild(rightCol)

	container := hb.NewDiv().
		Class("container-fluid").
		AddChild(row)

	footer := hb.NewFooter().
		Class("d-flex flex-wrap justify-content-between align-items-center py-3 my-4 border-top").
		AddChild(container)

	return footer
}
