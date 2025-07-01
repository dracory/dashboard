package bootstrap

import (
	"github.com/dracory/dashboard/types"
	"github.com/gouniverse/cdn"
	"github.com/gouniverse/hb"
)

// Template implements the types.TemplateInterface for Bootstrap-based templates
type Template struct{}

// Ensure Template implements the TemplateInterface
var _ types.TemplateInterface = (*Template)(nil)

// layout generates the main layout structure for the dashboard
func (t *Template) layout(dashboard types.DashboardInterface) string {
	content := dashboard.GetContent()
	layout := hb.NewBorderLayout()
	layout.AddTop(hb.Raw(topNavigation(dashboard)), hb.BORDER_LAYOUT_ALIGN_LEFT, hb.BORDER_LAYOUT_ALIGN_MIDDLE)
	layout.AddCenter(hb.Raw(content), hb.BORDER_LAYOUT_ALIGN_LEFT, hb.BORDER_LAYOUT_ALIGN_TOP)
	return layout.ToHTML()
}

// ToHTML generates the complete HTML for the dashboard page
func (t *Template) ToHTML(dashboard types.DashboardInterface) string {
	// Create a new webpage
	webpage := hb.Webpage()

	// Set the page title
	webpage.SetTitle(dashboard.GetTitle())

	// Add favicon
	if favicon := favicon(); favicon != "" {
		webpage.SetFavicon(favicon)
	}

	// Add CSS
	if style := templateStyle(); style != "" {
		webpage.AddStyle(style)
	}

	// Add JavaScript
	if script := templateScript(); script != "" {
		webpage.AddScript(script)
	}

	// Add Bootstrap CSS
	webpage.AddStyleURL(cdn.BootstrapCss_5_3_3())

	// Add Bootstrap Icons
	webpage.AddStyleURL(cdn.BootstrapIconsCss_1_11_3())

	// Add Bootstrap JS Bundle with Popper
	webpage.AddScriptURL(cdn.BootstrapJs_5_3_3())

	// Add custom CSS if any
	for _, style := range dashboard.GetStyles() {
		webpage.AddStyle(style)
	}

	// Add custom scripts if any
	for _, script := range dashboard.GetScripts() {
		webpage.AddScript(script)
	}

	// Generate the layout
	layoutHTML := t.layout(dashboard)

	// Add the layout to the webpage
	webpage.Body().Child(hb.Raw(layoutHTML))

	// Handle redirect if needed
	if redirectURL := dashboard.GetRedirectUrl(); redirectURL != "" {
		redirectTime := "0"
		if dashboard.GetRedirectTime() != "" {
			redirectTime = dashboard.GetRedirectTime()
		}
		webpage.Meta(hb.Meta().
			Attr("http-equiv", "refresh").
			Attr("content", redirectTime+"; url = "+redirectURL))
	}

	// Generate the final HTML
	return webpage.ToHTML()
}
