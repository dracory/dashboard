package bootstrap

import (
	"github.com/dracory/dashboard/types"
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
	layout.AddTop(hb.Raw(dashboardMenuNavbar(dashboard)), hb.BORDER_LAYOUT_ALIGN_LEFT, hb.BORDER_LAYOUT_ALIGN_MIDDLE)
	layout.AddCenter(hb.Raw(center(content)), hb.BORDER_LAYOUT_ALIGN_LEFT, hb.BORDER_LAYOUT_ALIGN_TOP)
	return layout.ToHTML()
}

// center is a helper function to center content
func (t *Template) center(content string) string {
	return content
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
	webpage.AddStyleURL("https://cdn.jsdelivr.net/npm/bootstrap@5.1.3/dist/css/bootstrap.min.css")

	// Add Bootstrap Icons
	webpage.AddStyleURL("https://cdn.jsdelivr.net/npm/bootstrap-icons@1.8.1/font/bootstrap-icons.css")

	// Add Bootstrap JS Bundle with Popper
	webpage.AddScriptURL("https://cdn.jsdelivr.net/npm/bootstrap@5.1.3/dist/js/bootstrap.bundle.min.js")

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

// center is a helper function to center content
func center(content string) string {
	return `<div class="d-flex justify-content-center align-items-center" style="height: 100vh;">` + content + `</div>`
}
