package dashboard

// ToHTML renders the dashboard as HTML
func (d *Dashboard) ToHTML() string {
	// We'll implement a simple wrapper function that calls the render package
	// This avoids the import cycle by using a separate function
	return RenderDashboardToHTML(d)
}
