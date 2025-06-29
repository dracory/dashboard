package dashboard

import (
	"github.com/dracory/dashboard/model"
	"github.com/dracory/dashboard/render"
)

// RenderDashboardToHTML renders the dashboard to HTML using the render package
// This function is separate from the Dashboard.ToHTML method to avoid import cycles
func RenderDashboardToHTML(d model.DashboardRenderer) string {
	// Use the component-based render logic from the render package
	page := render.RenderPage(d)
	
	// Convert the rendered page to HTML string
	return page.ToHTML()
}
